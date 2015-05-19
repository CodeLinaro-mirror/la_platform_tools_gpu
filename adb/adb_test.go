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

package adb

import (
	"reflect"
	"strings"
	"testing"
)

func TestParseDevices(t *testing.T) {
	for i, c := range []struct {
		s string
		e error
		d []Device
	}{
		{
			d: []Device{{"02a2a2de20d7f6de", Unauthorized}},
			e: nil,
			s: `* daemon not running. starting it now on port 5037
* * daemon started successfully
* List of devices attached
02a2a2de20d7f6de        unauthorized
`,
		},
		{
			d: []Device{},
			e: nil,
			s: `List of devices attached
`,
		},
	} {
		devices, err := parseDevices(c.s)
		if c.e != err {
			t.Errorf("(%d) Expected error: %v, got error: %v", i, c.e, err)
		}
		count := len(devices)
		if count < len(c.d) {
			count = len(c.d)
		}
		for j := 0; j < count; j++ {
			var got, expected Device
			if j < len(devices) {
				got = *devices[j]
			}
			if j < len(c.d) {
				expected = c.d[j]
			}
			if got != expected {
				t.Errorf("(%d) Device %d was not as expected. Expected: %v, got: %v", i, j, expected, got)
			}
		}
	}
}

func TestParsePackages(t *testing.T) {
	for i, c := range []struct {
		s string
		e error
		p []InstalledPackage
	}{
		{
			p: []InstalledPackage{
				{Name: "android", Path: "/system/framework/framework-res.apk"},
				{Name: "com.android.dreams.basic", Path: "/system/app/BasicDreams.apk"},
				{Name: "com.google.earth", Path: "/data/app/com.google.earth-1.apk"},
			},
			e: nil,
			s: "package:/system/framework/framework-res.apk=android\n" +
				"package:/system/app/BasicDreams.apk=com.android.dreams.basic\r\n" +
				"package:/data/app/com.google.earth-1.apk=com.google.earth\n",
		},
	} {
		pkgs, err := parsePackages(c.s, nil)
		if c.e != err {
			t.Errorf("(%d) Expected error: %v, got error: %v", i, c.e, err)
		}
		count := len(pkgs)
		if count < len(c.p) {
			count = len(c.p)
		}
		for j := 0; j < count; j++ {
			var got, expected InstalledPackage
			if j < len(pkgs) {
				got = *pkgs[j]
			}
			if j < len(c.p) {
				expected = c.p[j]
			}
			if got != expected {
				t.Errorf("(%d) Package %d was not as expected. Expected: %+v, got: %+v", i, j, expected, got)
			}
		}
	}
}

func (t treeNode) String() string {
	var s string
	if t.depth > 0 {
		s = strings.Repeat(" ", t.depth) + t.text
	} else {
		s = t.text
	}
	if cnt := len(t.children); cnt > 0 {
		c := make([]string, cnt)
		for i := range c {
			c[i] = t.children[i].String()
		}
		return s + "\n" + strings.Join(c, "\n")
	} else {
		return s
	}
}

func TestParseTabbedTree(t *testing.T) {
	expected := `
0
  00
    000
    001
  01
    010
    011
      0110
  02
1
  10`
	got := parseTabbedTree(expected).String()
	if got != expected {
		t.Errorf("Tree was not as expected.\nExpected: %v\nGot: %v", expected, got)
	}
}

func TestParseActions(t *testing.T) {
	for i, c := range []struct {
		s string
		e error
		a []Action
	}{
		{
			a: []Action{
				{
					Name:       "android.intent.action.MAIN",
					Component:  "com.google.foo/.FooActivity",
					Categories: []string{"android.intent.category.LAUNCHER"},
				}, {
					Name:       "com.google.android.FOO",
					Component:  "com.google.foo/.FooActivity",
					Categories: []string{"android.intent.category.DEFAULT"},
				}, {
					Name:       "android.intent.action.SEARCH",
					Component:  "com.google.foo/.FooActivity",
					Categories: []string{"android.intent.category.DEFAULT"},
				}},
			e: nil,
			s: `
Activity Resolver Table:
  Non-Data Actions:
      android.intent.action.MAIN:
        43178558 com.google.foo/.FooActivity filter 4327f110
          Action: "android.intent.action.MAIN"
          Category: "android.intent.category.LAUNCHER"
      com.google.android.FOO:
        43178558 com.google.foo/.FooActivity filter 431d7db8
          Action: "com.google.android.FOO"
          Category: "android.intent.category.DEFAULT"
      android.intent.action.SEARCH:
        43178558 com.google.foo/.FooActivity filter 4327cc40
          Action: "android.intent.action.SEARCH"
          Category: "android.intent.category.DEFAULT"`,
		},
	} {
		actions, err := parseActions(c.s)
		if c.e != err {
			t.Errorf("(%d) Expected error: %v, got error: %v", i, c.e, err)
		}
		count := len(actions)
		if count < len(c.a) {
			count = len(c.a)
		}
		for j := 0; j < count; j++ {
			var got, expected *Action
			if j < len(actions) {
				got = &actions[j]
			}
			if j < len(c.a) {
				expected = &c.a[j]
			}
			if !reflect.DeepEqual(got, expected) {
				t.Errorf("(%d) Action %d was not as expected. Expected: %+v, got: %+v", i, j, expected, got)
			}
		}
	}
}
