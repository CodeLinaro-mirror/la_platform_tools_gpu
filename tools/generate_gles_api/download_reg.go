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

package main

import (
	"encoding/xml"
	"fmt"
	"strings"
)

const url = "https://cvs.khronos.org/svn/repos/ogl/trunk/doc/registry/public/api/gl.xml"

// Download the khronos XML registry file
func DownloadReg() *Registry {
	bytes := Download(url)
	reg := &Registry{}
	if err := xml.Unmarshal(bytes, reg); err != nil {
		panic(err.Error())
	}
	return reg
}

type KhronosAPI string
type Version string

const GLES1API = KhronosAPI("gles1")
const GLES2API = KhronosAPI("gles2") // Includes GLES 3.0 and later

func (v Version) String() string { return fmt.Sprintf("%s", string(v)) }

type Registry struct {
	Group     []Group            `xml:"groups>group"`
	Enums     []Enums            `xml:"enums"`
	Command   []Command          `xml:"commands>command"`
	Feature   []Feature          `xml:"feature"`
	Extension []ExtensionElement `xml:"extensions>extension"`
}

type NamedElementList []NamedElement
type NamedElement struct {
	Name string `xml:"name,attr"`
}

type Group struct {
	NamedElement
	Enum NamedElementList `xml:"enum"`
}

type Enums struct {
	Namespace string `xml:"namespace,attr"`
	Group     string `xml:"group,attr"`
	Type      string `xml:"type,attr"` // "bitmask"
	Comment   string `xml:"comment,attr"`
	Enum      []Enum `xml:"enum"`
}

type Enum struct {
	NamedElement
	Value string     `xml:"value,attr"`
	Type  string     `xml:"type,attr"` // "u" or "ull"
	API   KhronosAPI `xml:"api,attr"`
	Alias string     `xml:"alias,attr"`
}

type Command struct {
	Proto ProtoOrParam   `xml:"proto"`
	Param []ProtoOrParam `xml:"param"`
	Alias NamedElement   `xml:"alias"`
}

type ProtoOrParam struct {
	InnerXML string `xml:",innerxml"`
	Chardata string `xml:",chardata"`
	Group    string `xml:"group,attr"`
	Length   string `xml:"len,attr"`
	Ptype    string `xml:"ptype"`
	Name     string `xml:"name"`
}

type Feature struct {
	NamedElement
	API     KhronosAPI          `xml:"api,attr"`
	Number  Version             `xml:"number,attr"`
	Require RequireOrRemoveList `xml:"require"`
	Remove  RequireOrRemoveList `xml:"remove"`
}

type ExtensionElement struct {
	NamedElement
	Supported string              `xml:"supported,attr"`
	Require   RequireOrRemoveList `xml:"require"`
	Remove    RequireOrRemoveList `xml:"remove"`
}

type RequireOrRemoveList []RequireOrRemove
type RequireOrRemove struct {
	API     KhronosAPI       `xml:"api,attr"` // for extensions only
	Profile string           `xml:"profile,attr"`
	Comment string           `xml:"comment,attr"`
	Enum    NamedElementList `xml:"enum"`
	Command NamedElementList `xml:"command"`
}

func (l NamedElementList) Contains(name string) bool {
	for _, v := range l {
		if v.Name == name {
			return true
		}
	}
	return false
}

func (r *RequireOrRemove) Contains(name string) bool {
	return r.Enum.Contains(name) || r.Command.Contains(name)
}

func (l RequireOrRemoveList) Contains(name string) bool {
	for _, v := range l {
		if v.Contains(name) {
			return true
		}
	}
	return false
}

func (e *ExtensionElement) IsSupported(api KhronosAPI) bool {
	for _, v := range strings.Split(e.Supported, "|") {
		if KhronosAPI(v) == api {
			return true
		}
	}
	return false
}

func (c Command) Name() string {
	return c.Proto.Name
}

func (p ProtoOrParam) Type() string {
	name := p.InnerXML
	name = name[:strings.Index(name, "<name>")]
	name = strings.Replace(name, "<ptype>", "", 1)
	name = strings.Replace(name, "</ptype>", "", 1)
	name = strings.TrimSpace(name)
	return name
}

// Add the return value as param "result" at slot -1
func (cmd *Command) ParamsAndResult() map[int]ProtoOrParam {
	result := cmd.Proto
	result.Name = "result"
	results := map[int]ProtoOrParam{-1: result}
	for i, param := range cmd.Param {
		results[i] = param
	}
	return results
}

// Return sorted list of versions which support the given symbol
func (r *Registry) GetVersions(api KhronosAPI, name string) []Version {
	version, found := Version(""), false
	for _, feature := range r.Feature {
		if feature.API == api {
			if feature.Require.Contains(name) {
				if found {
					panic(fmt.Errorf("redefinition of %s", name))
				}
				version, found = feature.Number, true
			}
			if feature.Remove != nil {
				// not used in GLES
				panic(fmt.Errorf("remove tag is not supported"))
			}
		}
	}
	if found {
		switch version {
		case "1.0":
			return []Version{"1.0", "1.1"}
		case "2.0":
			return []Version{"2.0", "3.0", "3.1"}
		case "3.0":
			return []Version{"3.0", "3.1"}
		case "3.1":
			return []Version{"3.1"}
		default:
			panic(fmt.Errorf("Uknown GLES version: %v", version))
		}
	} else {
		return nil
	}
}

// Return extensions which define the given symbol
func (r *Registry) GetExtensions(api KhronosAPI, name string) []string {
	var extensions []string
ExtensionLoop:
	for _, extension := range r.Extension {
		if extension.IsSupported(api) {
			for _, require := range extension.Require {
				if require.API == "" || require.API == api {
					if require.Contains(name) {
						extensions = append(extensions, extension.Name)
						// sometimes the extension repeats definition - ignore
						continue ExtensionLoop
					}
				}
			}
			if extension.Remove != nil {
				// not used in GLES
				panic(fmt.Errorf("remove tag is not supported"))
			}
		}
	}
	return extensions
}
