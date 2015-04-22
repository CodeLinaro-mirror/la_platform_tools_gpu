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

package build

import (
	"bytes"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

// File represents the path to a file or directory.
type File string

// Path returns a File formed from the joined path segments.
func Path(segments ...string) File {
	return File(filepath.Join(segments...))
}

// Join returns a File formed from joining this File with ext.
func (f File) Join(ext ...string) File {
	return File(filepath.Join(append([]string{string(f)}, ext...)...))
}

// RelativeTo returns the path of this File relative to base.
func (f File) RelativeTo(base File) string {
	rel, err := filepath.Rel(string(base), string(f))
	if err != nil {
		panic(err)
	}
	return rel
}

// Absolute returns the absolute path of this File.
func (f File) Absolute() string {
	abs, err := filepath.Abs(string(f))
	if err != nil {
		panic(err)
	}
	return abs
}

// Exists returns true if this File exists.
func (f File) Exists() bool {
	_, err := os.Stat(string(f.Absolute()))
	return err == nil
}

// Contains returns true if this directory contains the file f.
func (d File) Contains(f File) bool {
	a := d.Absolute()
	b := f.Absolute()
	if len(b) <= len(a)+1 {
		return false
	}
	return b[:len(a)] == a && b[len(a)] == filepath.Separator
}

// LastModified returns the time the file was last modified.
func (f File) LastModified() time.Time {
	s, err := os.Stat(string(f.Absolute()))
	if err != nil {
		panic(err)
	}
	return s.ModTime()
}

// Delete deletes the File.
func (f File) Delete() error {
	return os.Remove(string(f))
}

// Name returns the name part of the File (without directories).
func (f File) Name() string {
	return filepath.Base(string(f))
}

// Dir returns the directory part of the File (without filename).
func (f File) Dir() string {
	return filepath.Dir(string(f))
}

// Ext returns the extension of the file, including the '.', or an empty string
// if the file has no extension.
func (f File) Ext() string {
	return filepath.Ext(string(f))
}

// ChangeExt returns a new File with the extension changed to ext.
func (f File) ChangeExt(ext string) File {
	prev := filepath.Ext(string(f))
	noext := f[:len(f)-len(prev)]
	return noext + File(ext)
}

// Glob returns a FileSet of all files matching any of the specified patterns.
func (f File) Glob(patterns ...string) FileSet {
	fs := FileSet{}
	for _, patten := range patterns {
		glob, err := filepath.Glob(string(f.Join(patten)))
		if err != nil {
			panic(err)
		}
		files := make(FileSet, len(glob))
		for i := range glob {
			files[i] = File(glob[i])
		}
		fs = fs.Append(files...)
	}
	return fs
}

// Matches returns true if the File matches any of the patterns.
func (f File) Matches(patterns ...string) bool {
	abs := filepath.Base(f.Absolute())
	for _, pattern := range patterns {
		matched, err := filepath.Match(pattern, abs)
		if err != nil {
			panic(err)
		}
		if matched {
			return true
		}
	}
	return false
}

// MkdirAll creates the directory hierarchy to hold this File.
func (f File) MkdirAll() {
	abs := f.Absolute()
	dir, _ := filepath.Split(abs)
	if len(dir) > 0 {
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			panic(err)
		}
	}
}

// CopyTo copied this File to dst, replacing any existing file at dst.
func (f File) CopyTo(dst File) error {
	s, err := os.Open(f.Absolute())
	if err != nil {
		return err
	}
	defer s.Close()

	fi, err := s.Stat()
	if err != nil {
		return err
	}

	d, err := os.Create(dst.Absolute())
	if err != nil {
		return err
	}
	defer d.Close()

	if err := d.Chmod(fi.Mode()); err != nil {
		return err
	}

	_, err = io.Copy(d, s)
	return err
}

// Exec executes this File with the specified arguments.
func (f File) Exec(env Environment, args ...string) error {
	return f.ExecAt(env, "", args...)
}

// ExecAt executes this File with the specified arguments with the working
// directory set to wd.
func (f File) ExecAt(env Environment, wd File, args ...string) error {
	logger := env.Logger.Enter("Exec")

	var path string
	if f.Exists() {
		path = string(f)
	} else {
		var err error
		path, err = exec.LookPath(string(f))
		if err != nil {
			return err
		}
	}
	if env.Verbose {
		logger.Info("%s %v", path, args)
	}

	cmd := exec.Command(path, args...)
	cmd.Dir = string(wd)
	buffer := &bytes.Buffer{}
	cmd.Stdout = buffer
	cmd.Stderr = buffer
	err := cmd.Run()
	switch {
	case err != nil:
		logger.Error("\n\n%s\n--- %s failed: %v ---", string(buffer.Bytes()), f.Name(), err)

	case env.Verbose:
		if msg := string(buffer.Bytes()); msg != "" {
			logger.Info("\n%s\n--- %s succeeded ---", string(buffer.Bytes()), f.Name())
		} else {
			logger.Info("%s succeeded", f.Name())
		}
	}
	return err
}

// LookPath looks for the file f on the system PATH, returning the absolute
// path to the file if found, otherwise an empty File.
func (f File) LookPath() File {
	if path, err := exec.LookPath(string(f)); err == nil {
		return File(path)
	} else {
		return ""
	}
}
