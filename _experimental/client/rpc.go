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

package client

import (
	"fmt"
	"sync"

	"android.googlesource.com/platform/tools/gpu/atom"
	"android.googlesource.com/platform/tools/gpu/binary/registry"
	"android.googlesource.com/platform/tools/gpu/binary/schema"
	"android.googlesource.com/platform/tools/gpu/log"
	"android.googlesource.com/platform/tools/gpu/service"
	"android.googlesource.com/platform/tools/gpu/service/path"
)

type rpc struct {
	logger log.Logger
	client service.Client
	ready  <-chan struct{}
}

func (r *rpc) init(logger log.Logger, client service.Client, constants map[string]schema.ConstantSet) {
	ready := make(chan struct{})

	r.logger = logger
	r.ready = ready

	go func() {
		// make the decoder namespace try the global namespace before the schema one
		schemaNamespace := registry.NewNamespace()
		namespace := registry.NewNamespace(registry.Global, schemaNamespace)

		s, err := client.GetSchema(r.logger)
		if err != nil {
			log.Errorf(r.logger, "Error resolving schema: %v", err)
			return
		}
		log.Infof(r.logger, "Schema with %d classes, %d constant sets", len(s.Classes), len(s.Constants))
		atoms := 0
		for _, class := range s.Classes {
			// Find the atom metadata, if present
			if meta := atom.FindMetadata(class); meta != nil {
				atoms++
				schemaNamespace.Add(NewAtomClass(class, meta))
			} else {
				schemaNamespace.Add(class)
			}
		}
		log.Infof(r.logger, "Schema with %d atoms", atoms)
		for _, s := range s.Constants {
			constants[s.Type.String()] = s
		}
		// Replace the current RPC
		r.client = service.NewClient(client.Multiplexer(), namespace)
		close(ready)
	}()
}

func (r *rpc) beginRPC(name string) log.Logger {
	<-r.ready
	return log.Enter(log.Fork(r.logger), name)
}

func (r *rpc) GetCaptures() ([]capture, error) {
	l := r.beginRPC("GetCaptures")
	ids, err := r.client.GetCaptures(l)
	if err != nil {
		return nil, err
	}

	wg := sync.WaitGroup{}
	wg.Add(len(ids))
	captures := make([]capture, len(ids))
	for i := range ids {
		i := i
		go func() {
			defer wg.Done()
			captures[i].path = ids[i]
			if c, err := service.GetCapture(captures[i].path, r.client, l); err == nil {
				captures[i].info = c
			}
		}()
	}
	wg.Wait()
	return captures, nil
}

func (r *rpc) GetDevices() ([]device, error) {
	l := r.beginRPC("GetDevices")
	ids, err := r.client.GetDevices(l)
	if err != nil {
		return nil, err
	}

	wg := sync.WaitGroup{}
	wg.Add(len(ids))
	devices := make([]device, len(ids))
	for i := range ids {
		i := i
		go func() {
			defer wg.Done()
			devices[i].path = ids[i]
			if d, err := service.GetDevice(devices[i].path, r.client, l); err == nil {
				devices[i].info = d
			}
		}()
	}
	wg.Wait()
	return devices, nil
}

func (r *rpc) Import(name string, data []byte) (*path.Capture, error) {
	l := r.beginRPC("Import")

	return r.client.Import(name, data, l)
}

func (r *rpc) LoadCapture(p *path.Capture) (service.Capture, error) {
	l := r.beginRPC("LoadCapture")

	if capture, err := service.GetCapture(p, r.client, l); err == nil {
		return service.Capture{}, err
	} else {
		return *capture, nil
	}
}

func (r *rpc) LoadAtoms(p *path.Atoms) ([]atom.Atom, error) {
	l := r.beginRPC("LoadAtoms")

	stream, err := r.client.Get(p, l)
	if err != nil {
		return nil, err
	}

	return stream.(*atom.List).Atoms, nil
}

func (r *rpc) LoadHierarchy(p *path.Hierarchy) (atom.Group, error) {
	l := r.beginRPC("LoadHierarchy")

	root, err := r.client.Get(p, l)
	if err != nil {
		return atom.Group{}, err
	}

	return *root.(*atom.Group), nil
}

func (r *rpc) LoadTiming(device *path.Device, capture *path.Capture, flags service.TimingFlags) (service.TimingInfo, error) {
	l := r.beginRPC("LoadTiming")

	p, err := r.client.GetTimingInfo(device, capture, flags, l)
	if err != nil {
		return service.TimingInfo{}, err
	}

	timingInfo, err := service.GetTimingInfo(p, r.client, l)
	if err != nil {
		return service.TimingInfo{}, err
	}

	return *timingInfo, nil
}

func (r *rpc) LoadReport(capture *path.Capture) (service.Report, error) {
	l := r.beginRPC("LoadReport")

	report, err := r.client.Get(capture.Report(), l)
	if err != nil {
		return service.Report{}, nil
	}

	return *report.(*service.Report), nil
}

func (r *rpc) LoadState(state *path.State) (*schema.Object, error) {
	l := r.beginRPC("LoadState")

	object, err := r.client.Get(state, l)
	if err != nil {
		return nil, err
	}

	return object.(*schema.Object), nil
}

func (r *rpc) LoadMemory(rng *path.MemoryRange) (service.MemoryInfo, error) {
	l := r.beginRPC("LoadMemory")

	object, err := r.client.Get(rng, l)
	if err != nil {
		return service.MemoryInfo{}, err
	}

	return *object.(*service.MemoryInfo), nil
}

func (r *rpc) RequestColorBuffer(device *path.Device, after *path.Atom, settings service.RenderSettings) (w, h int, d []byte, e error) {
	l := r.beginRPC("RequestColorBuffer")

	imageID, err := r.client.GetFramebufferColor(device, after, settings, l)
	if err != nil {
		return 0, 0, nil, err
	}

	imageInfo, err := service.GetImageInfo(imageID, r.client, l)
	if err != nil {
		return 0, 0, nil, err
	}

	data, err := service.GetBlob(imageInfo.Data, r.client, l)
	if err != nil {
		return 0, 0, nil, err
	}

	if imageInfo.Width <= 0 || imageInfo.Height <= 0 {
		return 0, 0, nil, fmt.Errorf("Invalid image dimensions %dx%d", imageInfo.Width, imageInfo.Height)
	}
	return int(imageInfo.Width), int(imageInfo.Height), data, nil
}

func (r *rpc) RequestDepthBuffer(device *path.Device, after *path.Atom) (w, h int, d []byte, e error) {
	l := r.beginRPC("RequestDepthBuffer")

	imageID, err := r.client.GetFramebufferDepth(device, after, l)
	if err != nil {
		return 0, 0, nil, err
	}

	imageInfo, err := service.GetImageInfo(imageID, r.client, l)
	if err != nil {
		return 0, 0, nil, err
	}

	data, err := service.GetBlob(imageInfo.Data, r.client, l)
	if err != nil {
		return 0, 0, nil, err
	}

	if imageInfo.Width <= 0 || imageInfo.Height <= 0 {
		return 0, 0, nil, fmt.Errorf("Invalid image dimensions %dx%d", imageInfo.Width, imageInfo.Height)
	}
	return int(imageInfo.Width), int(imageInfo.Height), data, nil
}

func (r *rpc) Change(p path.Path, v interface{}) (path.Path, error) {
	l := r.beginRPC("Change")
	log.I(l, "%v -> %v", p, v)
	p, err := r.client.Set(p, v, l)
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (r *rpc) Follow(p path.Path) (path.Path, error) {
	l := r.beginRPC("Follow")
	p, err := r.client.Follow(p, l)
	if err != nil {
		return nil, err
	}
	return p, nil
}
