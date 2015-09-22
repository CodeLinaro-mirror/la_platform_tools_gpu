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

package do

import (
	"io"
	"log"
	"os"

	"android.googlesource.com/platform/tools/gpu/maker/config"
	"android.googlesource.com/platform/tools/gpu/maker/graph"
)

// CopyFile builds a new Step that copies the file from src to dst, and returns
// the dst entity.
// The step will depend on the src, and the existance of the directory dst is
// inside.
func CopyFile(dst, src *graph.Path) {
	graph.NewStep(func(s *graph.Step) error {
		if config.Verbose > 0 {
			log.Printf("-> cp %v to %v", src, dst)
		}
		in, err := os.Open(src.Name())
		if err != nil {
			return err
		}
		defer in.Close()
		out, err := os.Create(dst.Name())
		if err != nil {
			return err
		}
		defer out.Close()
		_, err = io.Copy(out, in)
		if err != nil {
			return err
		}
		return nil
	}).Creates(dst).DependsOn(src, dst.Parent())
}
