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

package builder

import (
	"fmt"

	"android.googlesource.com/platform/tools/gpu/atom"
	"android.googlesource.com/platform/tools/gpu/database"
	"android.googlesource.com/platform/tools/gpu/gfxapi"
	"android.googlesource.com/platform/tools/gpu/log"
	"android.googlesource.com/platform/tools/gpu/service"
)

// BuildLazy writes to out the schema.Report resource resulting from the given ResolveReport request.
func (r *BuildReport) BuildLazy(c interface{}, d database.Database, l log.Logger) (interface{}, error) {
	atoms, err := ResolveAtoms(r.Capture.Atoms(), d, l)
	if err != nil {
		return nil, err
	}

	report := &service.Report{}

	s := gfxapi.NewState()

	mutate := func(i int, a atom.Atom) {
		defer func() {
			if err := recover(); err != nil {
				report.Items = append(report.Items, service.ReportItem{
					Severity: log.Critical,
					Message:  fmt.Sprintf("%s", err),
					Atom:     uint64(i),
				})
			}
		}()
		if err := a.Mutate(s, d, l); err != nil {
			report.Items = append(report.Items, service.ReportItem{
				Severity: log.Error,
				Message:  err.Error(),
				Atom:     uint64(i),
			})
		}
	}

	for i, a := range atoms {
		mutate(i, a)
	}

	return report, nil
}
