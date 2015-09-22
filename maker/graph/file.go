// Copyright (C) 2015 The Android Open Source Project
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package graph

import (
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"time"

	"path/filepath"
)

// canonicalPath is an internal function that resolves paths with any kind of slash to the os specific fully resolved
// form.
func canonicalPath(path ...string) (string, error) {
	return filepath.Abs(filepath.FromSlash(filepath.Join(path...)))
}

// File returns an Entity that represents a file. The entities name will be the
// absolute path of the file. The entity map will be checked for a matching file
// entry and if one is not found, a new one will be added and returned.
// base can be either a string or *Path, all strings will have slash translation applied.
func File(base interface{}, join ...string) *Path {
	path := make([]string, 1+len(join))
	switch base := base.(type) {
	case *Path:
		path[0] = base.Name()
	case string:
		path[0] = filepath.FromSlash(base)
	}
	copy(path[1:], join)
	abs, err := canonicalPath(path...)
	if err != nil {
		log.Fatalf("%s", err)
	}
	abs = filepath.Clean(abs)
	e := FindEntity(abs)
	if e != nil {
		f, is := e.(*Path)
		if !is {
			log.Fatalf("%s is not a file entity (%T)", abs, e)
		}
		return f
	}
	f := &Path{abs: abs}
	f.stat, _ = os.Stat(f.abs)
	AddEntity(f)
	return f
}

// Dir returns an Entity that represents a directory. The entities name will be
// the absolute path of the directory. The entity map will be checked for a
// matching directory entry and if one is not found, a new one will be added and
// returned.
// It also adds the rules to create the directory if needed.
// base can be either a string or *Path, all strings will have slash translation applied.
func Dir(base interface{}, join ...string) *Path {
	d := File(base, join...)
	if Creator(d) == nil {
		NewStep(func(s *Step) error {
			if err := os.MkdirAll(d.abs, os.ModePerm); err != nil {
				return err
			}
			return nil
		}).Creates(d)
	}
	return d
}

// FindTool finds an executable on the host search path, and returns a File
// entity for the tool if found.
func FindTool(name string) *Path {
	path, err := exec.LookPath(name)
	if err != nil {
		return nil
	}
	return File(path)
}

// Path is the common type used for both file and directory entities in the build graph.
type Path struct {
	abs   string
	stat  os.FileInfo
	isDir bool
}

// Name returns the full absolute path to the file.
func (f *Path) Name() string { return f.abs }

// String returns the Name of the file.
func (f *Path) String() string { return f.abs }

// Timestamp returns the last modified time reported by the file system.
func (f *Path) Timestamp() time.Time {
	if f.stat == nil {
		return time.Time{}
	}
	if f.stat.IsDir() {
		// Do not cause dependency to rebuild just because of a directory change
		return time.Time{}
	}
	return f.stat.ModTime()
}

// NeedsUpdate returns true if the file does not exist, or is older than the
// supplied timestamp.
func (f *Path) NeedsUpdate(t time.Time) bool {
	if f.stat == nil {
		return true
	}
	return f.stat.ModTime().Before(t)
}

// Updated refreshes the exists and timestamp information for the file.
func (f *Path) Updated() { f.stat, _ = os.Stat(f.abs) }

// Files reads the list of files in path.
func (p *Path) Files() Set {
	infos, _ := ioutil.ReadDir(p.abs)
	list := Set{}
	for _, i := range infos {
		if !i.IsDir() {
			list = list.Append(File(p, i.Name()))
		}
	}
	return list
}

// Parent returns the parent directory of this path.
func (p *Path) Parent() *Path {
	path := filepath.Dir(p.abs)
	if len(path) == 0 {
		return nil
	}
	return Dir(path)
}

// Child returns a child directory of this path.
func (p *Path) Child(join ...string) *Path {
	return Dir(p, join...)
}

// File returns a file in this directory.
func (p *Path) File(join ...string) *Path {
	return File(p, join...)
}

// Basename returns the last path component of this filename.
func (p *Path) Basename() string {
	return filepath.Base(p.abs)
}
