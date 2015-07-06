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

	"android.googlesource.com/platform/tools/gpu/atom"
	"android.googlesource.com/platform/tools/gpu/database"
	"android.googlesource.com/platform/tools/gpu/log"
	"android.googlesource.com/platform/tools/gpu/service"
	"android.googlesource.com/platform/tools/gpu/service/path"
)

// BuildLazy creates a copy of the capture referenced by the request's path, but
// with the object, value or memory at p replaced with v. The path returned is
// identical to p, but with the base changed to refer to the new capture.
func (request *Set) BuildLazy(c interface{}, d database.Database, l log.Logger) (interface{}, error) {
	p := path.Flatten(request.Path)
	v, err := resolveChain(p, d, l)
	if err != nil {
		return nil, err
	}

	// Start by replacing the target value
	v[len(v)-1] = request.Value

	// Propagate changes back down to the root
	for i := len(v) - 1; i >= 0; i-- {
		switch p := p[i].(type) {
		case *path.Capture:
			id, err := database.Store(v[i], d, l)
			if err != nil {
				return nil, err
			}
			// Return the path to the new capture
			newPath := request.Path.Clone()
			for p := newPath; p != nil; p = p.Base() {
				if root, ok := p.(*path.Capture); ok {
					root.ID = id
					return newPath, nil
				}
			}
			panic("Unreachable")

		case *path.Atoms:
			stream := v[i].(*service.AtomStream)
			streamID, err := service.StoreAtomStream(stream, d, l)
			if err != nil {
				return nil, err
			}

			reportID, err := getBuildReport(streamID, d, l)
			if err != nil {
				return nil, err
			}

			capture := v[i-1].(*service.Capture)
			capture.Name = capture.Name + "*"
			capture.Atoms = streamID
			capture.Report = reportID
			v[i-1] = capture

		case *path.Atom:
			if v[i] == nil {
				return nil, fmt.Errorf("Atom cannot be nil")
			}
			stream := v[i-1].(*service.AtomStream).Clone()
			stream.Atoms[p.Index] = v[i].(atom.Atom)
			v[i-1] = stream

		case *path.State:
			return nil, fmt.Errorf("State can not currently be mutated")

		case *path.Field:
			obj, err := clone(reflect.ValueOf(v[i-1]))
			if err != nil {
				return nil, err
			}
			s := obj
			for s.Kind() == reflect.Ptr {
				s = s.Elem() // Deref
			}
			s.FieldByName(p.Name).Set(reflect.ValueOf(v[i]))
			v[i-1] = obj.Interface()

		case *path.ArrayIndex:
			a, err := clone(reflect.ValueOf(v[i-1]))
			if err != nil {
				return nil, err
			}
			a.Index(int(p.Index)).Set(reflect.ValueOf(v[i]))
			v[i-1] = a.Interface()

		case *path.MapIndex:
			m, err := clone(reflect.ValueOf(v[i-1]))
			if err != nil {
				return nil, err
			}
			m.MapIndex(reflect.ValueOf(p.Key)).Set(reflect.ValueOf(v[i]))
			v[i-1] = m.Interface()

		default:
			return nil, fmt.Errorf("Unknown path type %T", p)
		}
	}

	panic("Unreachable")
}

func clone(v reflect.Value) (reflect.Value, error) {
	o := reflect.New(v.Type()).Elem()
	return o, shallowCopy(o, v)
}

func shallowCopy(dst, src reflect.Value) error {
	switch dst.Kind() {
	case reflect.Ptr, reflect.Interface:
		if !src.IsNil() {
			o := reflect.New(src.Elem().Type())
			shallowCopy(o.Elem(), src.Elem())
			dst.Set(o)
		}

	case reflect.Slice, reflect.Array:
		reflect.Copy(dst, src)

	case reflect.Map:
		for _, k := range src.MapKeys() {
			dst.SetMapIndex(k, src.MapIndex(k))
		}

	default:
		dst.Set(src)
	}
	return nil
}
