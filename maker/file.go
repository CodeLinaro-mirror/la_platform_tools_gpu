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

package maker

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"time"
)

// File returns an Entity that represents a file.
// The entities name will be the absolute path of the file.
// If path is an entity already, it will be tested to make sure it is a file and
// returned.
// If it is a string, the the entity map will be checked for a matching file
// entry and if one is not found, a new one will be added and returned.
func File(path interface{}) *file {
	switch path := path.(type) {
	case string:
		abs, err := filepath.Abs(path)
		if err != nil {
			log.Fatalf("%s", err)
		}
		e := FindEntity(abs)
		if e != nil {
			f, is := e.(*file)
			if !is {
				log.Fatalf("%s is not a file entity (%T)", abs, e)
			}
			return f
		}
		f := &file{abs: abs}
		f.stat, _ = os.Stat(f.abs)
		AddEntity(f)
		return f
	case *file:
		return path
	default:
		log.Fatalf("cannot convert from %T to file entity", path)
		return nil
	}
}

// Dir returns an Entity that represents a directory.
// The entities name will be the absolute path of the directory.
// If path is an entity already, it will be tested to make sure it is a
// directory and returned.
// If it is a string, the the entity map will be checked for a matching directory
// entry and if one is not found, a new one will be added and returned.
// It also adds the rules to create the directory if needed.
func Dir(path interface{}) *file {
	d := File(path)
	if Creator(d) == nil {
		NewStep(makeDir).Creates(d)
	}
	return d
}

// DirOf attempts to make a Dir for the parent of the specified File.
func DirOf(path interface{}) *file {
	switch path := path.(type) {
	case string:
		return dirOf(path)
	case *file:
		return dirOf(path.Name())
	default:
		log.Fatalf("cannot get parent dir from %T", path)
		return nil
	}
}

// FilesOf reads the list of files in path, filters them with the supplied
// filter and returns the set of file entities that matched.
func FilesOf(path interface{}, filter func(os.FileInfo) bool) []*file {
	dir := Dir(path)
	infos, _ := ioutil.ReadDir(dir.Name())
	files := []*file{}
	for _, i := range infos {
		if filter(i) {
			files = append(files, File(filepath.Join(dir.Name(), i.Name())))
		}
	}
	return files
}

// IsFile returns true if the supplied entity is of file type.
func IsFile(e Entity) bool {
	_, is := e.(*file)
	return is
}

type file struct {
	abs  string
	stat os.FileInfo
}

// Name returns the full absolute path to the file.
func (f *file) Name() string { return f.abs }

// String returns the Name of the file.
func (f *file) String() string { return f.abs }

// Exists returns true if the file could be statted.
func (f *file) Exists() bool { return f.stat != nil }

// Timestamp returns the last modified time reported by the file system.
func (f *file) Timestamp() time.Time {
	if f.stat == nil {
		return time.Time{}
	}
	return f.stat.ModTime()
}

// NeedsUpdate returns true if the file does not exist, or is older than the
// supplied timestamp.
func (f *file) NeedsUpdate(t time.Time) bool {
	if f.stat == nil {
		return true
	}
	return f.stat.ModTime().Before(t)
}

// Updated refreshes the exists and timestamp information for the file.
func (f *file) Updated() { f.stat, _ = os.Stat(f.abs) }

func dirOf(name string) *file {
	path := filepath.Dir(name)
	if len(path) == 0 {
		return nil
	}
	return Dir(path)
}

func makeDir(s *Step) error {
	for _, out := range s.outputs {
		if err := os.MkdirAll(out.Name(), os.ModePerm); err != nil {
			return err
		}
	}
	return nil
}
