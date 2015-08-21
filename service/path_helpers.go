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

package service

import (
	"fmt"

	"android.googlesource.com/platform/tools/gpu/atom"
	"android.googlesource.com/platform/tools/gpu/binary"
	"android.googlesource.com/platform/tools/gpu/database"
	"android.googlesource.com/platform/tools/gpu/image"
	"android.googlesource.com/platform/tools/gpu/log"
	"android.googlesource.com/platform/tools/gpu/service/path"
)

// GetBlob calls s.Get with p and then safely casts the result to a []byte.
func GetBlob(p *path.Blob, s Service, l log.Logger) ([]byte, error) {
	if v, err := s.Get(p, l); err != nil {
		return nil, err
	} else if r, ok := v.([]byte); !ok {
		return nil, fmt.Errorf("path %s gave %T, expected []byte", p, v)
	} else {
		return r, nil
	}
}

// ResolveAtomList resolves an AtomsID and then safely casts the result to a *atom.List.
func ResolveAtomList(id AtomsID, d database.Database, l log.Logger) (*atom.List, error) {
	if v, err := database.Resolve(binary.ID(id), d, l); err != nil {
		return nil, err
	} else if r, ok := v.(*atom.List); !ok {
		return nil, fmt.Errorf("ID %s gave %T, expected *atom.List", id, v)
	} else {
		return r, nil
	}
}

// GetAtomList resolves an AtomsID and then safely casts the result to a *atom.List.
func GetAtomList(p *path.Atoms, s Service, l log.Logger) (*atom.List, error) {
	if v, err := s.Get(p, l); err != nil {
		return nil, err
	} else if r, ok := v.(*atom.List); !ok {
		return nil, fmt.Errorf("path %s gave %T, expected *atom.List", p, v)
	} else {
		return r, nil
	}
}

// ResolveCapture resolves a binary.ID and then safely casts the result to a *Capture.
func ResolveCapture(id binary.ID, d database.Database, l log.Logger) (*Capture, error) {
	if v, err := database.Resolve(id, d, l); err != nil {
		return nil, err
	} else if r, ok := v.(*Capture); !ok {
		return nil, fmt.Errorf("ID %s gave %T, expected Capture", id, v)
	} else {
		return r, nil
	}
}

// GetCapture calls s.Get with p and then safely casts the result to a *Capture.
func GetCapture(p *path.Capture, s Service, l log.Logger) (*Capture, error) {
	if v, err := s.Get(p, l); err != nil {
		return nil, err
	} else if r, ok := v.(*Capture); !ok {
		return nil, fmt.Errorf("path %s gave %T, expected Capture", p, v)
	} else {
		return r, nil
	}
}

// ResolveDevice resolves a binary.ID and then safely casts the result to a *Device.
func ResolveDevice(id binary.ID, d database.Database, l log.Logger) (*Device, error) {
	if v, err := database.Resolve(id, d, l); err != nil {
		return nil, err
	} else if r, ok := v.(*Device); !ok {
		return nil, fmt.Errorf("ID %s gave %T, expected Device", id, v)
	} else {
		return r, nil
	}
}

// GetDevice calls s.Get with p and then safely casts the result to a *Device.
func GetDevice(p *path.Device, s Service, l log.Logger) (*Device, error) {
	if v, err := s.Get(p, l); err != nil {
		return nil, err
	} else if r, ok := v.(*Device); !ok {
		return nil, fmt.Errorf("path %s gave %T, expected Device", p, v)
	} else {
		return r, nil
	}
}

// ResolveImageInfo resolves a binary.ID and then safely casts the result to a *image.Info.
func ResolveImageInfo(id binary.ID, d database.Database, l log.Logger) (*image.Info, error) {
	if v, err := database.Resolve(id, d, l); err != nil {
		return nil, err
	} else if r, ok := v.(*image.Info); !ok {
		return nil, fmt.Errorf("ID %s gave %T, expected image.Info", id, v)
	} else {
		return r, nil
	}
}

// GetImageInfo calls s.Get with p and then safely casts the result to a *image.Info.
func GetImageInfo(p *path.ImageInfo, s Service, l log.Logger) (*image.Info, error) {
	if v, err := s.Get(p, l); err != nil {
		return nil, err
	} else if r, ok := v.(*image.Info); !ok {
		return nil, fmt.Errorf("path %s gave %T, expected image.Info", p, v)
	} else {
		return r, nil
	}
}

// ResolveTimingInfo resolves a binary.ID and then safely casts the result to a *TimingInfo.
func ResolveTimingInfo(id binary.ID, d database.Database, l log.Logger) (*TimingInfo, error) {
	if v, err := database.Resolve(id, d, l); err != nil {
		return nil, err
	} else if r, ok := v.(*TimingInfo); !ok {
		return nil, fmt.Errorf("ID %s gave %T, expected TimingInfo", id, v)
	} else {
		return r, nil
	}
}

// GetTimingInfo calls s.Get with p and then safely casts the result to a *TimingInfo.
func GetTimingInfo(p *path.TimingInfo, s Service, l log.Logger) (*TimingInfo, error) {
	if v, err := s.Get(p, l); err != nil {
		return nil, err
	} else if r, ok := v.(*TimingInfo); !ok {
		return nil, fmt.Errorf("path %s gave %T, expected TimingInfo", p, v)
	} else {
		return r, nil
	}
}
