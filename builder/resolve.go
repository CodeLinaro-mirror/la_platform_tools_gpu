package builder

import (
	"fmt"
	"reflect"

	"android.googlesource.com/platform/tools/gpu/atom"
	"android.googlesource.com/platform/tools/gpu/database"
	"android.googlesource.com/platform/tools/gpu/gfxapi"
	"android.googlesource.com/platform/tools/gpu/log"
	"android.googlesource.com/platform/tools/gpu/service"
	"android.googlesource.com/platform/tools/gpu/service/path"
)

// Resolve resolves and returns the object, value or memory at the path p.
func Resolve(p path.Path, d database.Database, l log.Logger) (interface{}, error) {
	n := path.Flatten(p)
	v := make([]interface{}, len(n))

	for i, p := range n {
		switch p := p.(type) {
		case *path.Capture:
			capture, err := service.ResolveCapture(service.CaptureId{ID: p.ID}, d, l)
			if err != nil {
				return nil, err
			}
			v[i] = capture

		case *path.Atoms:
			capture := v[i-1].(service.Capture)
			atoms, err := loadAtoms(capture.Atoms, d, l)
			if err != nil {
				return nil, err
			}
			v[i] = atoms

		case *path.Atom:
			atoms := v[i-1].(atom.List)
			if p.Index >= uint64(len(atoms)) {
				return nil, fmt.Errorf("Atom index (%d) is out of bounds [0-%d]", p.Index, len(atoms)-1)
			}
			v[i] = atoms[p.Index]

		case *path.State:
			atoms := v[i-2].(atom.List)
			api := gfxapi.Find(atoms[p.After.Index].API())
			s := gfxapi.NewState()
			for _, a := range atoms[:p.After.Index] {
				a.Mutate(s, d, l)
			}
			res, found := s.APIs[api]
			if !found {
				return nil, fmt.Errorf("No state for API '%v'", api.Name())
			}
			v[i] = res

		case *path.Field:
			s := reflect.ValueOf(v[i-1])
			switch s.Kind() {
			case reflect.Struct:
				v[i] = s.FieldByName(p.Name).Interface()

			default:
				return nil, fmt.Errorf("Cannot access fields of type %T", v)
			}

		case *path.ArrayIndex:
			a := reflect.ValueOf(v[i-1])
			switch a.Kind() {
			case reflect.Array, reflect.Slice, reflect.String:
				v[i] = a.Index(int(p.Index)).Interface()

			default:
				return nil, fmt.Errorf("Cannot array-index type %T", v)
			}

		case *path.MapIndex:
			m := reflect.ValueOf(v[i-1])
			switch m.Kind() {
			case reflect.Map:
				v[i] = m.MapIndex(reflect.ValueOf(p.Key)).Interface()

			default:
				return nil, fmt.Errorf("Cannot map-index type %T", v)
			}

		default:
			return nil, fmt.Errorf("Unknown path type %T", p)
		}
	}

	return v[len(v)-1], nil
}
