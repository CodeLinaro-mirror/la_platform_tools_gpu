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
	"android.googlesource.com/platform/tools/gpu/memory"
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

func (r *rpc) GetCaptures() (map[service.CaptureID]service.Capture, error) {
	l := r.beginRPC("GetCaptures")
	ids, err := r.client.GetCaptures(l)
	if err != nil {
		return nil, err
	}

	wg := sync.WaitGroup{}
	wg.Add(len(ids))
	captures := make([]*service.Capture, len(ids))
	for i := range ids {
		i := i
		go func() {
			defer wg.Done()
			if capture, err := r.client.ResolveCapture(ids[i], l); err == nil {
				captures[i] = &capture
			} else {
				log.E(l, "Failed to resolve capture %v: %v", ids[i], err)
			}
		}()
	}
	wg.Wait()

	m := make(map[service.CaptureID]service.Capture, len(ids))
	for i := range ids {
		if c := captures[i]; c != nil {
			m[ids[i]] = *c
		}
	}
	return m, nil
}

func (r *rpc) GetDevices() (map[service.DeviceID]service.Device, error) {
	l := r.beginRPC("GetDevices")
	ids, err := r.client.GetDevices(l)
	if err != nil {
		return nil, err
	}

	wg := sync.WaitGroup{}
	wg.Add(len(ids))
	devices := make([]*service.Device, len(ids))
	for i := range ids {
		i := i
		go func() {
			defer wg.Done()
			if device, err := r.client.ResolveDevice(ids[i], l); err == nil {
				devices[i] = &device
			} else {
				log.E(l, "Failed to resolve device %v: %v", ids[i], err)
			}
		}()
	}
	wg.Wait()

	m := make(map[service.DeviceID]service.Device, len(ids))
	for i := range ids {
		if d := devices[i]; d != nil {
			m[ids[i]] = *d
		}
	}
	return m, nil
}

func (r *rpc) Import(name string, data []byte) (service.CaptureID, error) {
	l := r.beginRPC("Import")

	id, err := r.client.Import(name, data, l)
	if err != nil {
		log.E(l, "Error importing capture: %v", err)
	}

	return id, err
}

func (r *rpc) LoadCapture(captureID service.CaptureID) (service.Capture, error) {
	l := r.beginRPC("LoadCapture")

	capture, err := r.client.Get(captureID.Path(), l)
	if err != nil {
		log.E(l, "Error getting capture: %v", err)
		return service.Capture{}, err
	}

	return *capture.(*service.Capture), nil
}

func (r *rpc) LoadAtoms(p *path.Atoms) ([]atom.Atom, error) {
	l := r.beginRPC("LoadAtoms")

	stream, err := r.client.Get(p, l)
	if err != nil {
		log.E(l, "Error getting atom stream: %v", err)
		return nil, err
	}

	return stream.(*service.AtomStream).Atoms, nil
}

func (r *rpc) LoadHierarchy(p *path.Hierarchy) (atom.Group, error) {
	l := r.beginRPC("LoadHierarchy")

	hierarchy, err := r.client.Get(p, l)
	if err != nil {
		log.E(l, "Error getting hierarchy: %v", err)
		return atom.Group{}, err
	}

	return hierarchy.(*service.Hierarchy).Root, nil
}

func (r *rpc) LoadTiming(device *path.Device, capture *path.Capture, flags service.TimingFlags) (service.TimingInfo, error) {
	l := r.beginRPC("LoadTiming")

	timingInfoID, err := r.client.GetTimingInfo(device, capture, flags, l)
	if err != nil {
		return service.TimingInfo{}, err
	}

	timingInfo, err := r.client.ResolveTimingInfo(timingInfoID, l)
	if err != nil {
		return service.TimingInfo{}, err
	}

	return timingInfo, nil
}

func (r *rpc) LoadReport(capture *path.Capture) (service.Report, error) {
	l := r.beginRPC("LoadReport")

	report, err := r.client.Get(capture.Report(), l)
	if err != nil {
		log.E(l, "Error loading report: %v", err)
		return service.Report{}, nil
	}

	return *report.(*service.Report), nil
}

func (r *rpc) LoadState(state *path.State) (*schema.Object, error) {
	l := r.beginRPC("LoadState")

	object, err := r.client.Get(state, l)
	if err != nil {
		log.E(l, "Error loading state: %v", err)
		return nil, err
	}

	return object.(*schema.Object), nil
}

func (r *rpc) RequestColorBuffer(device *path.Device, after *path.Atom, settings service.RenderSettings) (w, h int, d []byte, e error) {
	l := r.beginRPC("RequestColorBuffer")

	imageID, err := r.client.GetFramebufferColor(device, after, settings, l)
	if err != nil {
		return 0, 0, nil, err
	}

	imageInfo, err := r.client.ResolveImageInfo(imageID, l)
	if err != nil {
		return 0, 0, nil, err
	}

	data, err := r.client.ResolveBinary(imageInfo.Data, l)
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

	imageInfo, err := r.client.ResolveImageInfo(imageID, l)
	if err != nil {
		return 0, 0, nil, err
	}

	data, err := r.client.ResolveBinary(imageInfo.Data, l)
	if err != nil {
		return 0, 0, nil, err
	}

	if imageInfo.Width <= 0 || imageInfo.Height <= 0 {
		return 0, 0, nil, fmt.Errorf("Invalid image dimensions %dx%d", imageInfo.Width, imageInfo.Height)
	}
	return int(imageInfo.Width), int(imageInfo.Height), data, nil
}

func (r *rpc) RequestMemory(after *path.Atom, base uint64, size uint64) (service.MemoryInfo, error) {
	l := r.beginRPC("RequestMemory")

	rng := memory.Range{Base: base, Size: size}
	id, err := r.client.GetMemoryInfo(after, rng, l)
	if err != nil {
		return service.MemoryInfo{}, err
	}

	info, err := r.client.ResolveMemoryInfo(id, l)
	if err != nil {
		return service.MemoryInfo{}, err
	}

	return info, nil
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
		log.E(l, "%v", err)
		return nil, err
	}
	return p, nil
}
