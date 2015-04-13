package replay

import (
	"fmt"

	"android.googlesource.com/platform/tools/gpu/atom"
	"android.googlesource.com/platform/tools/gpu/gfxapi/state"
	"android.googlesource.com/platform/tools/gpu/replay/builder"
)

// Replayer is the interface that wraps the basic Replay method.
type Replayer interface {
	// Replay issues replay operations to the replay builder b for the given atom
	// with identifier id, and graphics API state s. If the replay action will
	// have an effect on the graphics driver state, then the call to Replay should
	// also apply the corresponding changes to the state s. If postback is true
	// then the replay instructions should include postback of all outputs of the
	// action.
	Replay(id atom.ID, s *state.State, b *builder.Builder, postback bool)
}

// Replay issues replay operations to the replay builder b for the given atom a
// with identifier id, and graphics API state s. If replaying the Atom will have
// an effect on the graphics driver state, then the call to Replay will also
// apply the corresponding changes to the state s. If postback is true then
// the replay instructions should include postback of all outputs of the Atom.
func Replay(id atom.ID, a atom.Atom, s *state.State, b *builder.Builder, postback bool) {
	switch a := a.(type) {
	case *atom.Observation:
		b.Observation(a.Range, a.ResourceID)
	case Replayer:
		a.Replay(id, s, b, postback)
	default:
		panic(fmt.Errorf("Atom %d of type %T does not conform to the Replayer interface", id, a))
	}
}
