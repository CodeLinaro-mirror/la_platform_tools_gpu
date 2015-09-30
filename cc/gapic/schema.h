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

#ifndef GAPIC_SCHEMA_H
#define GAPIC_SCHEMA_H

#include <initializer_list>
#include <memory>
#include <string>

#include "encoder.h"
#include "log.h"

namespace gapic {
namespace schema {

class Type {
 public:
  enum TypeTag {
    PrimitiveTag,
    StructTag,
    PointerTag,
    InterfaceTag,
    VariantTag,
    AnyTag,
    SliceTag,
    ArrayTag,
    MapTag,
  };

  virtual void encode(Encoder& e) const = 0;
};

class Field {
 public:
  // Move
  Field(Field&& field) : mDeclared(field.mDeclared), mType(std::move(field.mType)) {}

  Field(const std::string& declared, Type* type)
      : mDeclared(declared), mType(type) {}

  void encode(Encoder& e) const {
    mType->encode(e);
  }
 private:
  std::string mDeclared;
  std::unique_ptr<Type> mType;
};

class Entity {
 public:
  Entity() = default;

  Entity(const gapic::Id& id,
         const std::string& package,
         const std::string& name,
         const std::string& identity,
         const std::string& version,
         std::initializer_list<Field> fields)
      :
      mTypeId(id),
      mPackage(package),
      mName(name),
      mIdentity(identity),
      mVersion(version),
      mExported(true),
      mFields(std::move(fields)) {}

  Entity(const gapic::Id& id,
         const std::string& package,
         const std::string& name,
         const std::string& identity,
         const std::string& version,
         std::initializer_list<Field> fields,
         std::initializer_list<std::unique_ptr<Encodable>> metadata)
      :
      mTypeId(id),
      mPackage(package),
      mName(name),
      mIdentity(identity),
      mVersion(version),
      mExported(true),
      mFields(std::move(fields)),
      mMetadata(std::move(metadata)) {}

  void encode(Encoder& e) const {
    e.String(mPackage);
    e.String(mIdentity);
    e.String(mVersion);
    e.Uint32(uint32_t(mFields.size()));
    for (const auto& f : mFields) {
      f.encode(e);
    }
  }

  const gapic::Id& TypeId() const {
    return mTypeId;
  }
 private:
  gapic::Id mTypeId;
  std::string mPackage;
  std::string mName;
  std::string mIdentity;
  std::string mVersion;
  bool mExported;
  std::initializer_list<Field> mFields;
  std::initializer_list<std::unique_ptr<Encodable>> mMetadata;
};

class Primitive : public Type {
 public:
  enum Method {
    ID,
    Bool,
    Int8,
    Uint8,
    Int16,
    Uint16,
    Int32,
    Uint32,
    Int64,
    Uint64,
    Float32,
    Float64,
    String,
  };

  Method method() const { return mMethod; }

  Primitive(const std::string& name, Method method) : mName(name), mMethod(method) {}

  void encode(Encoder& e) const {
    e.Uint8(uint8_t(Type::PrimitiveTag) | (uint8_t(mMethod) << 4));
  }
 private:
  std::string mName;
  Method mMethod;
};

class Struct : public Type {
 public:
  Struct(const std::string& relative, const Entity& entity)
      : mRelative(relative), mEntity(entity) {}

  void encode(Encoder& e) const {
    e.Uint8(uint8_t(StructTag));
    e.Entity(mEntity);
  }
 private:
  std::string mRelative;
  const Entity& mEntity;
};

class Pointer : public Type {
 public:
  explicit Pointer(Type* type) : mType(type) {}

  void encode(Encoder& e) const {
    e.Uint8(uint8_t(PointerTag));
    mType->encode(e);
  }
 private:
  std::unique_ptr<Type> mType;
};

class Interface : public Type {
 public:
  explicit Interface(const std::string& name) : mName(name) {}

  void encode(Encoder& e) const {
    e.Uint8(uint8_t(InterfaceTag));
  }
 private:
  std::string mName;
};

class Variant : public Type {
 public:
  explicit Variant(const std::string& name) : mName(name) {}

  void encode(Encoder& e) const {
    e.Uint8(uint8_t(VariantTag));
  }
 private:
  std::string mName;
};

class Any : public Type {
 public:
  Any() = default;

  void encode(Encoder& e) const {
    e.Uint8(uint8_t(AnyTag));
  }
};

class Slice : public Type {
 public:
  Slice(const std::string& alias, Type* valueType)
      : mAlias(alias), mValueType(valueType) {}
  void encode(Encoder& e) const {
    e.Uint8(uint8_t(SliceTag));
    mValueType->encode(e);
  }
 private:
  std::string mAlias;
  std::unique_ptr<Type> mValueType;
};

class Array : public Type {
 public:
  Array(const std::string& alias, Type* valueType, uint32_t size)
      : mAlias(alias), mValueType(valueType), mSize(size) {}
  void encode(Encoder& e) const {
    e.Uint8(uint8_t(ArrayTag));
    e.Uint32(mSize);
    mValueType->encode(e);
  }
 private:
  std::string mAlias;
  std::unique_ptr<Type> mValueType;
  uint32_t mSize;
};

class Map : public Type {
 public:
  Map(const std::string& alias, Type* keyType, Type* valueType)
      : mAlias(alias),
        mKeyType(keyType),
        mValueType(valueType) {}
  void encode(Encoder& e) const {
      e.Uint8(uint8_t(MapTag));
      mKeyType->encode(e);
      mValueType->encode(e);
  }
 private:
  std::string mAlias;
  std::unique_ptr<Type> mKeyType;
  std::unique_ptr<Type> mValueType;
};

class Constant {
 public:
  // Move
  Constant(Constant&& c) : mName(c.mName), mValue(c.mValue) {}

  // Only uint32_t is supported in C++ at the moment, but it would
  // be fairly easy to support all the primitive types.
  Constant(const std::string& name, uint32_t value) :
      mName(name), mValue(value) {}

  void encode(Encoder& e) const {
    e.String(mName);
    e.Uint32(mValue);
  }
 private:
  std::string mName;
  uint32_t mValue;
};

class ConstantSet {
 public:
  ConstantSet() = default;
  ConstantSet(Primitive* type, std::initializer_list<Constant> entries) :
      mType(type), mEntries(std::move(entries)) {
    if (mType->method() != Primitive::Uint32) {
      GAPID_FATAL("Constant and ConstantSet only support uint32");
    }
  }

  void encode(Encoder& e) const {
    mType->encode(e);
    e.Uint32(uint32_t(mEntries.size()));
    for (const auto& entry : mEntries) {
      entry.encode(e);
    }
  }
 private:
  std::unique_ptr<Primitive> mType;
  std::initializer_list<Constant> mEntries;
};

}  // namespace schema
}  // namespace gapic

#endif
