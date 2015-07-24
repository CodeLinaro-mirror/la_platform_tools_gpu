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

package template

import "fmt"

type counter int

// Set assigns a value to the counter.
func (c *counter) Set(value int) string {
	*c = counter(value)
	return ""
}

// AddLen increments the counter by the length of the supplied string.
func (c *counter) AddLen(value string) string {
	*c += counter(len(value))
	return ""
}

func (c *counter) String() string {
	return fmt.Sprint(*c)
}

// Counter returns a counter by name. If the counter does not exist, it will be created
// and initialized to zero.
func (t *Templates) Counter(name string) *counter {
	c, ok := t.counters[name]
	if !ok {
		c = new(counter)
		t.counters[name] = c
	}
	return c
}
