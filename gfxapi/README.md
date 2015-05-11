# gfxapi
--
    import "android.googlesource.com/platform/tools/gpu/gfxapi"

Package gfxapi exposes the shared behavior of all graphics api's.

Copyright (C) 2015 The Android Open Source Project

Licensed under the Apache License, Version 2.0 (the "License"); you may not use
this file except in compliance with the License. You may obtain a copy of the
License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software distributed
under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR
CONDITIONS OF ANY KIND, either express or implied. See the License for the
specific language governing permissions and limitations under the License.

## Usage

#### func  Register

```go
func Register(api API)
```
Register adds an api to the understood set. It is illegal to register the same
name twice.

#### type API

```go
type API interface {
	// Name returns the official name of the api.
	Name() string

	// ID returns the unique API identifier.
	ID() ID

	// GetFramebufferAttachmentSize returns the width and height of the framebuffer at the given attachment.
	GetFramebufferAttachmentSize(state *State, attachment FramebufferAttachment) (width uint32, height uint32, err error)
}
```

API is the common interface to a graphics programming api.

#### func  Find

```go
func Find(id ID) API
```
Find looks up a graphics API by identifier. If the id has not been registered,
it returns nil.

#### type FramebufferAttachment

```go
type FramebufferAttachment uint32
```

FramebufferAttachment values indicate the type of frame buffer attachment.

```go
const (
	FramebufferAttachmentColor FramebufferAttachment = iota
	FramebufferAttachmentDepth
	FramebufferAttachmentStencil
)
```

#### type ID

```go
type ID binary.ID
```

ID is an API identifier

#### func (ID) Valid

```go
func (i ID) Valid() bool
```
Valid returns true if the id is not the default zero value.

#### type State

```go
type State struct {
	binary.Object

	// Memory holds the memory state of the application.
	Memory memory.Memory

	// APIs holds the per-API context states.
	APIs map[API]interface{}
}
```

State represents the graphics state across all contexts.
