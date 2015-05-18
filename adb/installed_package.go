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
	"errors"
	"fmt"
	"strings"
)

// https://android.googlesource.com/platform/bionic/+/master/libc/include/sys/system_properties.h#38
// Actually 32, but that includes null-terminator.
const maxPropName = 31

type InstalledPackage struct {
	Name   string  // Name of the package.
	Path   string  // Path to the package's APK.
	Device *Device // The device this package is installed on.
}

// WrapProperties returns the list of wrap-properties for the given installed
// package.
func (p *InstalledPackage) WrapProperties() ([]string, error) {
	list, err := p.Device.Command("getprop", p.wrapPropName()).Call()
	return strings.Fields(list), err
}

// WrapProperties sets the list of wrap-properties for the given installed
// package.
func (p *InstalledPackage) SetWrapProperties(props ...string) error {
	arg := strings.Join(props, " ")
	return p.Device.Command("setprop", p.wrapPropName(), arg).Run()
}

// Action represents an Android action that can be sent as an intent.
type Action struct {
	Name       string   // Example: android.intent.action.MAIN
	Component  string   // Example: com.foo.bar/.FooBarActivity
	Categories []string // Example: android.intent.category.LAUNCHER
}

// Actions returns all the actions supported by the specified package.
func (p *InstalledPackage) Actions() ([]Action, error) {
	str, err := p.Device.Command("dumpsys", "package", p.Name).Call()
	if err != nil {
		return nil, err
	}
	return parseActions(str)
}

// String returns the package name.
func (p *InstalledPackage) String() string {
	return p.Name
}

func (p *InstalledPackage) wrapPropName() string {
	name := "wrap." + p.Name
	if len(name) > maxPropName {
		name = name[:maxPropName]
	}
	return name
}

// InstalledPackages returns the list of installed packages on the device
func (d *Device) InstalledPackages() ([]*InstalledPackage, error) {
	str, err := d.Command("pm", "list", "packages", "-f").Call()
	if err != nil {
		return nil, err
	}
	return parsePackages(str, d)
}

func parsePackages(str string, device *Device) ([]*InstalledPackage, error) {
	lines := strings.Split(str, "\n")
	packages := make([]*InstalledPackage, 0, len(lines))
	for _, line := range lines {
		line = strings.TrimRight(line, "\r")
		segments := strings.SplitAfter(line, "package:")
		if len(segments) != 2 {
			continue
		}
		fields := strings.Split(segments[1], "=")
		if len(fields) != 2 {
			return nil, errors.New("Could not parse package list")
		}
		pkg := &InstalledPackage{
			Path:   fields[0],
			Name:   fields[1],
			Device: device,
		}
		packages = append(packages, pkg)
	}
	return packages, nil
}

type treeNode struct {
	text     string
	children []*treeNode
	parent   *treeNode
	depth    int
}

func parseTabbedTree(str string) *treeNode {
	head := &treeNode{depth: -1}
	for _, line := range strings.Split(str, "\n") {
		line = strings.TrimRight(line, "\r")
		if line == "" {
			continue
		}

		// Calculate the line's depth
		depth := 0
		for i, r := range line {
			if r == ' ' {
				depth++
			} else {
				line = line[i:]
				break
			}
		}

		// Find the insertion point
		for {
			if head.depth >= depth {
				head = head.parent
			} else {
				node := &treeNode{text: line, depth: depth, parent: head}
				head.children = append(head.children, node)
				head = node
				break
			}
		}
	}
	for head.parent != nil {
		head = head.parent
	}
	return head
}

func unquote(s string) string {
	return strings.Trim(s, `"`)
}

// Currently parses only the non-data actions.
func parseActions(str string) ([]Action, error) {
	actions := []Action{}
	for _, root := range parseTabbedTree(str).children {
		if root.text == "Activity Resolver Table:" {
			for _, node := range root.children {
				if node.text == "Non-Data Actions:" {
					for _, node := range node.children {
						for _, node := range node.children {
							if action, err := parseAction(node); err == nil {
								actions = append(actions, action)
							} else {
								return nil, err
							}
						}
					}
				}
			}
		}
	}

	return actions, nil
}

func parseAction(node *treeNode) (Action, error) {
	action := Action{}

	// 43178558 com.google.foo/.FooActivity filter 431d7db8
	fields := strings.Fields(node.text)
	if len(fields) != 4 || fields[2] != "filter" {
		return action, fmt.Errorf("Could not parse component: '%v'", node.children[0].text)
	}

	action.Component = fields[1]

	for _, detail := range node.children {
		fields = strings.Fields(detail.text)
		if len(fields) != 2 {
			continue
		}
		switch fields[0] {
		case "Action:":
			action.Name = unquote(fields[1])

		case "Category:":
			action.Categories = append(action.Categories, unquote(fields[1]))

		default:
			return action, fmt.Errorf("Unknown field: '%v'", fields[0])
		}
	}

	return action, nil
}
