// Copyright (C) 2014 The Android Open Source Project
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

package scan

import "golang.org/x/tools/go/gcimporter"

// Process resolves all type informtaion for the loaded sources.
func (s *Scanner) Process() error {
	dirs := make([]*Directory, 0, len(s.Directories))
	for _, dir := range s.Directories {
		if dir.Scan {
			dirs = append(dirs, dir)
		}
	}
	for _, dir := range dirs {
		if err := s.process(dir); err != nil {
			return err
		}
	}
	return nil
}

func (s *Scanner) process(dir *Directory) error {
	if !dir.Scan && !dir.Module.processed && !s.ForceSource {
		t, err := gcimporter.Import(s.config.Packages, dir.ImportPath)
		if err == nil {
			dir.Module.processed = true
			dir.Module.Types = t
		}
	}
	if !dir.loaded {
		s.load(dir)
	}
	if !dir.Module.processed {
		dir.Module.processed = true
		if err := s.parse(&dir.Module); err != nil {
			return err
		}
		if err := s.typeCheck(dir, &dir.Module); err != nil {
			return err
		}
		s.config.Packages[dir.ImportPath] = dir.Module.Types
	}
	if dir.Scan && !dir.Test.processed && len(dir.Test.Sources) > 0 {
		dir.Test.processed = true
		dir.Test.Files = dir.Module.Files
		if err := s.parse(&dir.Test); err != nil {
			return err
		}
		if err := s.typeCheck(dir, &dir.Test); err != nil {
			return err
		}
	}
	return nil
}

func (s *Scanner) typeCheck(dir *Directory, module *Module) error {
	t, err := s.config.Check(dir.ImportPath, s.FileSet, module.Files, nil)
	module.Types = t
	return err
}
