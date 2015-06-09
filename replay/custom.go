package replay

import (
	"android.googlesource.com/platform/tools/gpu/atom"
	"android.googlesource.com/platform/tools/gpu/binary"
	"android.googlesource.com/platform/tools/gpu/database"
	"android.googlesource.com/platform/tools/gpu/gfxapi"
	"android.googlesource.com/platform/tools/gpu/log"
	"android.googlesource.com/platform/tools/gpu/replay/builder"
)

// Custom must conform to the replay.Replayer interface.
var _ = Replayer(Custom(nil))

// Custom is an atom issuing custom replay operations to the replay builder b upon Replay().
type Custom func(i atom.ID, s *gfxapi.State, d database.Database, l log.Logger, b *builder.Builder) error

func (c Custom) Replay(i atom.ID, s *gfxapi.State, d database.Database, l log.Logger, b *builder.Builder) error {
	return c(i, s, d, l, b)
}

// atom.Atom compliance
func (Custom) API() gfxapi.API                                                 { return nil }
func (Custom) TypeID() atom.TypeID                                             { return 0 }
func (Custom) Flags() atom.Flags                                               { return 0 }
func (Custom) Observations() *atom.Observations                                { return &atom.Observations{} }
func (Custom) Mutate(s *gfxapi.State, d database.Database, l log.Logger) error { return nil }
func (Custom) Class() binary.Class                                             { return nil }
