package state

//go:generate codergen -go

import (
	"android.googlesource.com/platform/tools/gpu/atom"
	"android.googlesource.com/platform/tools/gpu/binary"
	"android.googlesource.com/platform/tools/gpu/memory"
)

// FramebufferAttachment values indicate the type of frame buffer attachment.
type FramebufferAttachment uint32

const (
	FramebufferAttachmentColor FramebufferAttachment = iota
	FramebufferAttachmentDepth
	FramebufferAttachmentStencil
)

// State represents the graphics state across all contexts.
type State struct {
	binary.Object
	// Memory holds the memory state of the application.
	Memory memory.Memory
	// Contexts holds the per-context states.
	Contexts map[atom.ContextID]Context
}

// New returns a new and initialized State.
func New() *State {
	return &State{
		Contexts: make(map[atom.ContextID]Context),
	}
}

// StateMutator is the interface of types that can mutate a State.
type Mutator interface {
	Mutate(*State) error
}

// Mutate mutates the State using the atom a.
func (s *State) Mutate(a atom.Atom) error {
	switch a := a.(type) {
	case Mutator:
		return a.Mutate(s)
	case *atom.Observation:
		s.Memory.Slice(a.Range).Write(memory.ResourceData(a.ResourceID, a.Range.Size))
	}
	return nil
}

// Context represents the graphics state for a single graphics context.
type Context interface {
	binary.Object
	// GetFramebufferAttachmentSize returns the width and height of the framebuffer at the given attachment.
	GetFramebufferAttachmentSize(attachment FramebufferAttachment) (width uint32, height uint32, err error)
}
