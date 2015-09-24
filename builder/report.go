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
	"reflect"
	"sort"

	"android.googlesource.com/platform/tools/gpu/atom"
	"android.googlesource.com/platform/tools/gpu/binary"
	"android.googlesource.com/platform/tools/gpu/database"
	"android.googlesource.com/platform/tools/gpu/gfxapi"
	"android.googlesource.com/platform/tools/gpu/log"
	"android.googlesource.com/platform/tools/gpu/replay"
	"android.googlesource.com/platform/tools/gpu/service"
	"android.googlesource.com/platform/tools/gpu/service/path"
)

// BuildReport generates a service.Report for the given capture and optional
// device.
type BuildReport struct {
	binary.Generate
	Capture *path.Capture
	Device  *path.Device // If non-null, the report should include replay information.
}

// BuildLazy writes to out the schema.Report resource resulting from the given ResolveReport request.
func (r *BuildReport) BuildLazy(c interface{}, d database.Database, l log.Logger) (interface{}, error) {
	mgr := c.(*Context).ReplayManager

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

	// Gather report items from the state mutator, and collect together all the
	// APIs in use.
	apis := map[gfxapi.ID]struct{}{}
	for i, a := range atoms {
		apis[a.API()] = struct{}{}
		mutate(i, a)
	}

	if r.Device != nil {
		// Request is for a replay report too.
		ctx := replay.Context{
			Capture: r.Capture.ID,
			Device:  r.Device.ID,
		}

		// Capture can use multiple APIs. Iterate the APIs in use looking for
		// those that support the QueryIssues interface. Call QueryIssues for each
		// of these APIs, and use reflect.Select to gather all the issues.
		issues := []reflect.SelectCase{}
		for id := range apis {
			if api := gfxapi.Find(id); api != nil {
				if qi, ok := api.(replay.QueryIssues); ok {
					c := qi.QueryIssues(ctx, mgr)
					issues = append(issues, reflect.SelectCase{
						Dir:  reflect.SelectRecv,
						Chan: reflect.ValueOf(c),
					})
				}
			}
		}

		for {
			_, v, ok := reflect.Select(issues)
			if !ok {
				break // All issues gathered
			}
			issue := v.Interface().(replay.Issue)
			report.Items = append(report.Items, service.ReportItem{
				Severity: issue.Severity,
				Message:  issue.Error.Error(),
				Atom:     uint64(issue.Atom),
			})
		}

		// Items are now all out of order. Sort them.
		sort.Sort(reportSorter(report.Items))
	}

	return report, nil
}

// Used for sorting report items
type reportSorter []service.ReportItem

func (s reportSorter) Len() int           { return len(s) }
func (s reportSorter) Less(i, j int) bool { return s[i].Atom < s[j].Atom }
func (s reportSorter) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
