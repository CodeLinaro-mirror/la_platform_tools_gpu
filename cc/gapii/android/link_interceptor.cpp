/*
 * Copyright (C) 2015 The Android Open Source Project
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

// This module is concerned with installing the interceptors in a
// running program. This is required to do captures of Java programs
// on non-rooted Android devices. The interceptors are installed
// at library initialization time (before main). This is done by
// modifying the global offset table of the ELF libraries aleady
// loaded into memory. Some code from the NDK version of the crazy
// linker is used.

#include "dlinterceptor.h"

#include <gapic/dl_loader.h>
#include <gapii/gles_exports.h>
#include <gapic/mutex.h>
#include <gapic/lock.h>
#include <gapic/log.h>

#include <crazy_linker_elf_view.h>
#include <crazy_linker_elf_symbols.h>
#include <crazy_linker_error.h>
#include <linker_phdr.h>

#include <set>
#include <unordered_map>

#include <dlfcn.h>
#include <link.h>
#include <sys/mman.h>
#include <unistd.h>

namespace {

inline ELF::Addr page_start(ELF::Addr addr) {
  return addr & PAGE_MASK;
}

inline ELF::Addr page_end(ELF::Addr addr) {
  return page_start(addr + PAGE_SIZE - 1);
}

typedef int ELFPF;
typedef int ProtFlag;

inline ProtFlag maybe_map_flag(ELFPF pf, ELFPF from, ProtFlag prot) {
  if ( (pf & from) != 0 ) {
    return prot;
  }
  return 0;
}

inline ProtFlag elfpf_to_prot(ELFPF pf) {
  return maybe_map_flag(pf, PF_X, PROT_EXEC)
      | maybe_map_flag(pf, PF_R, PROT_READ)
      | maybe_map_flag(pf, PF_W, PROT_WRITE);
}

// Used to set the protection bits of all loaded segments plus PROT_WRITE.
// Note this is similar to the routine _phdr_table_set_load_prot in the
// crazy linker, however, that routine does not change the permission on
// pages which were marked writable in the ELF header. However, on Android
// some of these pages are no longer writable at the time we want to relink
// the program, so we need to mark them writable.
bool phdr_table_make_writable(const ELF::Phdr* phdr_table,
                              int phdr_count,
                              ELF::Addr load_bias,
                              crazy::Error* error) {
  const ELF::Phdr* phdr = phdr_table;
  const ELF::Phdr* phdr_limit = phdr + phdr_count;

  for (; phdr < phdr_limit; phdr++) {
    if (phdr->p_type != PT_LOAD) {
      continue;
    }

    ELF::Addr seg_page_start = page_start(phdr->p_vaddr) + load_bias;
    ELF::Addr seg_page_end =
        page_end(phdr->p_vaddr + phdr->p_memsz) + load_bias;

    int ret = mprotect(reinterpret_cast<void*>(seg_page_start),
                       seg_page_end - seg_page_start,
                       elfpf_to_prot(phdr->p_flags) | PROT_WRITE);
    if (ret < 0) {
      error->Format("mprotect %x..%x %s", seg_page_start, seg_page_end,
                    strerror(errno));
      return false;
    }
  }
  return true;
}

// Used to protect the load segments. Note this just resets the segments
// to the protection given in the ELF header. This is less restrictive
// than when the relink started. The alternative would be to read
// /proc/self/maps before starting and try to reset to those protections.
bool phdr_table_reset_load_prot(const ELF::Phdr* phdr_table,
                                int phdr_count,
                                ELF::Addr load_bias,
                                crazy::Error* error) {
  if (phdr_table_protect_segments(phdr_table, phdr_count, load_bias) != 0) {
    error->Format("failed to reset memory protection: %s", strerror(errno));
    return false;
  }
  return true;
}

// SymbolResolver used to provide a mapping from symbol name to address
// for the interceptor symbols.
class SymbolResolver {
 public:
  // Add appends a symbol to the resolver table.
  template<typename T>
  void Add(const char* name, T address);

  // Lookup a symbol_name. If the symbol is a known interceptor symbol
  // the address of the interceptor is returned. Otherwise zero is returned.
  ELF::Addr Lookup(const char* symbol_name) const;
 private:
  std::unordered_map<std::string, ELF::Addr> mSymbols;
};

template<typename T>
void SymbolResolver::Add(const char* name, T address) {
  mSymbols.emplace(name, reinterpret_cast<ELF::Addr>(address));
}

ELF::Addr SymbolResolver::Lookup(const char* symbol_name) const {
  auto it = mSymbols.find(symbol_name);
  if (it == mSymbols.end()) {
    return 0;
  }

  return it->second;
}

// Look the name of a relocation symbol in the ELF symbol table.
template <class Rel>  // ELF::Rel or ELF::Rela, with or without addeds.
const char* SymbolName(const Rel& rel, const crazy::ElfSymbols& symbols) {
  const ELF::Word rel_symbol = ELF_R_SYM(rel.r_info);
  if (rel_symbol == 0) {
    return NULL;
  }

  // If this is a symbolic relocation, compute the symbol's address.
  return symbols.LookupNameById(rel_symbol);
}

// Apply a single interceptor relocation. Symbol 'sym_name' interceptor at
// 'address'. RELA Relocation 'rela' and 'load_bias' for this shared object.
bool ApplyResolvedReloc(const char* sym_name, const ELF::Rela& rela, ELF::Addr address, size_t load_bias) {
  const ELF::Word rela_type = ELF_R_TYPE(rela.r_info);
  const ELF::Sword addend = rela.r_addend;

  const ELF::Addr reloc = static_cast<ELF::Addr>(rela.r_offset + load_bias);

  // Apply the relocation.
  ELF::Addr* target = reinterpret_cast<ELF::Addr*>(reloc);
  switch (rela_type) {
#ifdef __aarch64__
    case R_AARCH64_JUMP_SLOT:
      *target = address + addend;
      return true;

    case R_AARCH64_GLOB_DAT:
      *target = address + addend;
      return true;

    case R_AARCH64_ABS64:
      // implement if needed.
      GAPID_WARNING("Can not intercept R_AARCH64_ABS64 relocation %s", sym_name);
      return false;

    case R_AARCH64_RELATIVE:
      // implement if needed.
      GAPID_WARNING("Can not intercept R_AARCH64_RELATIVE relocation %s", sym_name);
      return false;

    case R_AARCH64_COPY:
      GAPID_WARNING("Invalid R_ARM_COPY relocation in shared library %s", sym_name);
      return false;

    default:
      GAPID_WARNING("Unknown rela_type in %d in %s", rela_type, sym_name);
      return false;
#else
    default:
      GAPID_FATAL("Trying a RELA relocation but the architecture is not AARCH64");
      return false;
#endif  // __aarch64__
  }

  return true;
}

// Apply a single interceptor relocation. Symbol 'sym_name' interceptor at
// 'address'. REL Relocation 'rel' and 'load_bias' for this shared object.
bool ApplyResolvedReloc(const char* sym_name, const ELF::Rel& rel, ELF::Addr address, size_t load_bias) {
  const ELF::Word rel_type = ELF_R_TYPE(rel.r_info);
  ELF::Addr* target = reinterpret_cast<ELF::Addr*>(rel.r_offset + load_bias);
  switch (rel_type) {
#ifdef __arm__
    case R_ARM_JUMP_SLOT:
    case R_ARM_GLOB_DAT:
      *target = address;
      return true;

    case R_ARM_ABS32:
      // implement if needed.
      GAPID_WARNING("Can not intercept R_ARM_ABS32 relocation %s", sym_name);
      return false;

    case R_ARM_REL32:
      // implement if needed.
      GAPID_WARNING("Can not intercept R_ARM_REL32 relocation %s", sym_name);
      return false;

    case R_ARM_RELATIVE:
      GAPID_DEBUG("R_ARM_RELATIVE safe to ignore %s", sym_name);
      return true;

    case R_ARM_COPY:
      GAPID_WARNING("Invalid R_ARM_COPY relocation in shared library %s", sym_name);
      return false;

    default:
      GAPID_WARNING("Unknown rel_type in %d in %s", rel_type, sym_name);
      return false;
#else
    default:
      GAPID_FATAL("Trying a REL relocation but the architecture is not ARM");
      return false;
#endif
  }
  // Not reached.
  GAPID_FATAL("This code should not have been reached in ApplyResolvedReloc");
  return false;
}

// Apply a single interceptor relocation 'rel'.
template <class Rel>  // ELF::Rel or ELF::Rela, with or without addeds.
bool ApplyRelReloc(const Rel& rel,
                   const crazy::ElfSymbols& symbols,
                   const SymbolResolver& resolver, size_t load_bias) {
  const ELF::Word rel_type = ELF_R_TYPE(rel.r_info);
  if (rel_type == R_ARM_NONE) {
    return true;
  }

  // Lookup the symbol name in the ELF symbol table.
  const char* sym_name = SymbolName(rel, symbols);
  if (sym_name == NULL) {
    return true;
  }

  ELF::Addr address = resolver.Lookup(sym_name);
  if (!address) {
    // Don't need to do anything for this symbol (not an interceptor).
    return true;
  }
  return ApplyResolvedReloc(sym_name, rel, address, load_bias);
}

// Apply the interceptor relocations for the relocation table 'rel'
// of size 'rel_count'.
template <class Rel>   // ELF::Rel or ELF::Rela, with or without addeds.
bool ApplyRelRelocs(const Rel* rel,
                    size_t rel_count,
                    const crazy::ElfSymbols& symbols,
                    const SymbolResolver& resolver,
                    size_t load_bias) {
  if (!rel) {
    return true;
  }

  int errors = 0;
  for (size_t rel_n = 0; rel_n < rel_count; rel++, rel_n++) {
    if (!ApplyRelReloc(*rel, symbols, resolver, load_bias)) {
      errors++;
    }
  }

  return errors == 0;
}

class ElfReloc {
 public:
  ElfReloc() : relocations_type_(0),
               plt_relocations_(0),
               plt_relocations_size_(0),
               relocations_(0),
               relocations_size_(0) {}

  // Read the dynamic section in the ELF header to locate the relocation
  // tables. This is analogous to ElfRelocations::Init in the crazy linker.
  // We can not use ElfRelocations directly as it does additive relocations,
  // which are not safe to repeat.
  bool Init(const crazy::ElfView& view, crazy::Error* error);

  // Apply the interceptor relocations to a shared object.
  // name - is the name of the shared object (for logging)
  // symbols - ELF symbol table for this shared object.
  // reloc - ELF relocation table for this shared object.
  // resolver - provides symbol to address mapping for the interceptors.
  // load_bias - load bias of the shared object (how much it moved by relocation).
  //
  // Analogous to ElfRelocations::ApplyAll in the crazy linker.
  bool ApplyInterceptorRelocations(const char* name,
                                   const crazy::ElfSymbols& symbols,
                                   const SymbolResolver& resolver,
                                   size_t load_bias);
 private:
  // Type of relocation DT_REL or DT_RELA
  ELF::Addr relocations_type_;
  // Address and size of the PLT relocation table.
  ELF::Addr plt_relocations_;
  size_t plt_relocations_size_;
  // Address and size of the relocation table.
  ELF::Addr relocations_;
  size_t relocations_size_;
};

bool ElfReloc::Init(const crazy::ElfView& view, crazy::Error* error) {
  // We handle only Rel or Rela, but not both. If DT_RELA or DT_RELASZ
  // then we require DT_PLTREL to agree.
  bool has_rela_relocations = false;
  bool has_rel_relocations = false;

  // Parse the dynamic table.
  crazy::ElfView::DynamicIterator dyn(&view);
  for (; dyn.HasNext(); dyn.GetNext()) {
    ELF::Addr dyn_value = dyn.GetValue();
    uintptr_t dyn_addr = dyn.GetAddress(view.load_bias());

    const ELF::Addr tag = dyn.GetTag();
    switch (tag) {
      case DT_PLTREL:
        if (dyn_value != DT_REL && dyn_value != DT_RELA) {
          *error = "Invalid DT_PLTREL value in dynamic section";
          return false;
        }
        relocations_type_ = dyn_value;
        break;
      case DT_JMPREL:
        plt_relocations_ = dyn_addr;
        break;
      case DT_PLTRELSZ:
        plt_relocations_size_ = dyn_value;
        break;
      case DT_RELA:
      case DT_REL:
        if (relocations_) {
          *error = "Unsupported DT_RELA/DT_REL combination in dynamic section";
          return false;
        }
        relocations_ = dyn_addr;
        if (tag == DT_RELA)
          has_rela_relocations = true;
        else
          has_rel_relocations = true;
        break;
      case DT_RELASZ:
      case DT_RELSZ:
        if (relocations_size_) {
          *error = "Unsupported DT_RELASZ/DT_RELSZ combination in dyn section";
          return false;
        }
        relocations_size_ = dyn_value;
        if (tag == DT_RELASZ)
          has_rela_relocations = true;
        else
          has_rel_relocations = true;
        break;
      default:
        ;
    }
  }

  if (relocations_type_ != DT_REL && relocations_type_ != DT_RELA) {
    *error = "Unsupported or missing DT_PLTREL in dynamic section";
    return false;
  }

  if (relocations_type_ == DT_REL && has_rela_relocations) {
    *error = "Found DT_RELA in dyn section, but DT_PLTREL is DT_REL";
    return false;
  }
  if (relocations_type_ == DT_RELA && has_rel_relocations) {
    *error = "Found DT_REL in dyn section, but DT_PLTREL is DT_RELA";
    return false;
  }

  return true;
}

bool ElfReloc::ApplyInterceptorRelocations(const char* name,
                                           const crazy::ElfSymbols& symbols,
                                           const SymbolResolver& resolver,
                                           size_t load_bias) {
  // There are two types of relocation DT_REL and DT_RELA. The only difference
  // is that RELA relocations have an addend. This is an additional offset
  // from the location, it primarily exists to make the job of the static
  // linker easier. The other way of looking at it is that REL relocation is
  // a relocation with an addend of zero. Note a shared object has one or the
  // other, not both. The ARM32 linker use REL and the ARM64 linker uses RELA.
  if (relocations_type_ == DT_REL) {
    // There are two relocation tables in the header. This is because when
    // it is supported PLT relocations are done lazily. Lazy resolution is not
    // supported on Android ARM32 or ARM64, so we treat both tables the same.
    if (!ApplyRelRelocs(reinterpret_cast<ELF::Rel*>(relocations_),
                        relocations_size_ / sizeof(ELF::Rel),
                        symbols,
                        resolver, load_bias)) {
      GAPID_WARNING("Errors in the Rel relocations for .dyn.rel: %s", name);
      return false;
    }
    if (!ApplyRelRelocs(reinterpret_cast<ELF::Rel*>(plt_relocations_),
                        plt_relocations_size_ / sizeof(ELF::Rel),
                        symbols,
                        resolver, load_bias)) {
      GAPID_WARNING("Errors in the Rel relocations for .dyn.plt: %s", name);
      return false;
    }
  } else if (relocations_type_ == DT_RELA) {
    if (!ApplyRelRelocs(reinterpret_cast<ELF::Rela*>(relocations_),
                        relocations_size_ / sizeof(ELF::Rela),
                        symbols,
                        resolver, load_bias)) {
      GAPID_WARNING("Errors in the Rela relocations for .dyn.rel: %s", name);
      return false;
    }
    if (!ApplyRelRelocs(reinterpret_cast<ELF::Rela*>(plt_relocations_),
                        plt_relocations_size_ / sizeof(ELF::Rela),
                        symbols,
                        resolver, load_bias)) {
      GAPID_WARNING("Errors in the Rela relocations for .dyn.plt: %s", name);
      return false;
    }
  }
  return true;
}

// Call dladdr and get the error message, if error.
bool dladdr_check(const void *addr, Dl_info *info) {
  dlerror();   // clear error condition
  if (dladdr(addr, info) == 0) {  // zero means error.
    const char* error = dlerror();
    if (error != NULL) {
      GAPID_WARNING("dladdr error %s", error);
    } else {
      GAPID_WARNING("unknown dladdr error");
    }
    return false;
  }
  return true;
}

// This symbol is just used so that to identify whether the current program
// header is for spy itself.
static const void* me = NULL;

// This callback is called by dl_iterate_phdr. It is passed the size and
// address of the ELF program headers. A SymbolResolver object is passed as
// 'data'. It modifies the ELF relocations for this program, so that the
// spy interceptor functions in the SymbolResolver are used in preference
// to the 'real' ones. It does nothing when it is called for spy.so itself.
int LinkInterceptorsCb(struct dl_phdr_info *info, size_t size, void *data) {
  const int kContinue = 0;

  if (info->dlpi_addr == 0 || info->dlpi_phdr == 0) {
    // For some unknown reason dl_iterate_phdr calls us with zeros, so
    // just ignore it and move on.
    return kContinue;
  }

  if (data == NULL) {
    GAPID_FATAL("LinkInterceptorsCb not passed a symbol resolver");
    return kContinue;
  }

  const SymbolResolver& resolver = *static_cast<SymbolResolver*>(data);
  crazy::Error error;
  crazy::ElfView elfView;
  crazy::ElfSymbols elfSymbols;
  ElfReloc elfReloc;

  // Identify the shared object we are looking at.
  Dl_info dl_info;
  if (!dladdr_check(info->dlpi_phdr, &dl_info)) {
    return kContinue;
  }
  const char* dlname = dl_info.dli_fname;

  if (gapii::DlInterceptor::isDriver(dlname)) {
    GAPID_DEBUG("Not patching %s as it is a driver", dlname);
    return kContinue;
  }

  // Identify the shared object which contains this code (i.e. the spy).
  Dl_info dl_self;
  if (!dladdr_check(&me, &dl_self)) {
    return kContinue;
  } else if (dl_info.dli_fbase == dl_self.dli_fbase) {
    // Don't insert interceptors into ourself, that would cause a loop.
    return kContinue;
  }

  if (!elfView.InitUnmapped(info->dlpi_addr, info->dlpi_phdr, info->dlpi_phnum, &error)) {
    GAPID_WARNING("%s: elfView.InitUnmapped failed", dlname);
  } else if (elfView.load_bias() != info->dlpi_addr) {
    // The load_bias must match the load address, because of a bug in the NDK
    // version of the crazy linker. The only case it is not likely to match
    // is when the ELF library was embedded in a zipfile. Ironically that
    // support was written by me (anton@).
    GAPID_WARNING("%s: load bias not dlpi_addr for addr %p phdr %p bias=%x: %s",
        dlname, info->dlpi_addr, info->dlpi_phdr, elfView.load_bias(), dlname);
  } else if (!elfReloc.Init(elfView, &error)) {
    GAPID_WARNING("%s: elfReloc.Init failed: %s", dlname, error.c_str());
  } else if (!elfSymbols.Init(&elfView)) {
    GAPID_WARNING("%s: elfSymbols.Init failed", dlname);
  } else if (!phdr_table_make_writable(info->dlpi_phdr, info->dlpi_phnum, elfView.load_bias(), &error)) {
    GAPID_WARNING("%s: phdr_table_make_writable failed: %s", dlname, error.c_str());
  } else if (!elfReloc.ApplyInterceptorRelocations(dlname, elfSymbols, resolver, elfView.load_bias())) {
    GAPID_WARNING("%s: elfReloc.ApplyInterceptorRelocations failed", dlname);
  } else if (!phdr_table_reset_load_prot(info->dlpi_phdr, info->dlpi_phnum, elfView.load_bias(), &error)) {
    GAPID_WARNING("%s: phdr_table_reset_load_prot failed: %s", dlname, error.c_str());
  }
  return kContinue;
}

// LinkDlInterceptorCb calls through to LinkInterceptorsCb *once* per loaded .so during
// the lifetime of the application.
gapic::Mutex gDlopenMutex;
int LinkDlInterceptorCb(struct dl_phdr_info *info, size_t size, void *data) {
  {
    gapic::Lock<gapic::Mutex> lock(&gDlopenMutex);
    static std::set<const char*> set;
    if (!set.insert(info->dlpi_name).second) {
      return 0; // Already patched this .so
    }
  }
  GAPID_INFO("Patching dlopen for %s", info->dlpi_name);
  return LinkInterceptorsCb(info, size, data);
}


void LinkGfxInterceptors() {
  using namespace gapii;

  SymbolResolver resolver;

  // kGLESExports is a NULL terminated, name to function table containing
  // all the interceptors. It is code generated and lives in gles_exports.cpp.
  // Build the symbol resolver from it.
  for (int i = 0; kGLESExports[i].mName != NULL; ++i) {
    resolver.Add(kGLESExports[i].mName, kGLESExports[i].mFunc);
  }

  // Check to see if the interceptors are needed. If the library is being
  // preloaded we don't need to do anything. We just look to see if
  // eglInitialize is an interceptor.
  static const char* eglInitialize = "eglInitialize";

  void* inUse = dlsym(RTLD_DEFAULT, eglInitialize);
  if (inUse == NULL) {
    GAPID_WARNING("dlsym did not find %s", eglInitialize);
    return;
  }

  ELF::Addr interceptor = resolver.Lookup(eglInitialize);
  if (interceptor == 0) {
    GAPID_WARNING("Did not find interceptor for %s", eglInitialize);
    return;
  }

  if (reinterpret_cast<void*>(interceptor) == inUse) {
    GAPID_DEBUG("Found interceptor %s in use. Not relinking", eglInitialize);
    return;
  }

  dl_iterate_phdr(LinkInterceptorsCb, &resolver);
}

void onDlopen(void* handle, const char* name) {
  using namespace gapii;

  // Patch the newly opened library so that dlopen() and dlsym() are intercepted.
  SymbolResolver resolver;
  resolver.Add("dlopen", DlInterceptor::dlopen);
  resolver.Add("dlsym", DlInterceptor::dlsym);
  dl_iterate_phdr(LinkDlInterceptorCb, &resolver);
}

void LinkDlInterceptor() {
  using namespace gapii;

  DlInterceptor::init(dlopen, dlsym, onDlopen);

  // Intercept external calls to dlopen() and dlsym() with the DlInterceptor functions.
  SymbolResolver resolver;
  resolver.Add("dlopen", DlInterceptor::dlopen);
  resolver.Add("dlsym", DlInterceptor::dlsym);

  // Make sure that internal calls to dlopen() and dlsym() continue to use the real functions.
  gapic::DlLoader::setCustomLoader(DlInterceptor::load);
  gapic::DlLoader::setCustomResolver(DlInterceptor::resolve);

  dl_iterate_phdr(LinkDlInterceptorCb, &resolver);
}

void LinkInterceptors() {
  LinkGfxInterceptors();
  LinkDlInterceptor();
}

}  // anonymous namespace

//
// Run the link interceptor automatically when the library is loaded.
//
// This is done so the only modification needed to a Java app is a call to
// load library in the main activity:
//
//   static {
//     System.loadLibrary("spy");
//   }
//
// As this means that the code runs before main, care needs to be taken to
// avoid using any other load time initialized globals, since they may not
// have been initialized yet.
//

class LinkInterceptorsHook {
 public:
  LinkInterceptorsHook() {
    LinkInterceptors();
  }
} gLinkInterceptorsHook;
