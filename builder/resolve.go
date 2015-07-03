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
	cache := map[service.CaptureId]atom.List{}
	load := func(id service.CaptureId) (atom.List, error) {
		atoms, found := cache[id]
		if found {
			return atoms, nil
		}
		capture, err := service.ResolveCapture(id, d, l)
		if err != nil {
			return nil, err
		}

		if atoms, err = loadAtoms(capture.Atoms, d, l); err != nil {
			return nil, err
		}

		cache[id] = atoms
		return atoms, nil
	}

	var resolve func(interface{}) (interface{}, error)

	resolve = func(p interface{}) (interface{}, error) {
		switch p := p.(type) {
		case *path.Atoms:
			atoms, err := load(service.CaptureId{ID: p.Capture.ID})
			if err != nil {
				return nil, err
			}
			return atoms, err

		case *path.Atom:
			a, err := resolve(p.Atoms)
			if err != nil {
				return nil, err
			}
			atoms := a.(atom.List)
			if p.Index >= uint64(len(atoms)) {
				return nil, fmt.Errorf("Atom index (%d) is out of bounds [0-%d]", p.Index, len(atoms)-1)
			}
			return atoms[p.Index], err

		case *path.State:
			a, err := resolve(p.After.Atoms)
			if err != nil {
				return nil, err
			}
			atoms := a.(atom.List)
			if p.After.Index >= uint64(len(atoms)) {
				return nil, fmt.Errorf("Atom index (%d) is out of bounds [0-%d]", p.After.Index, len(atoms)-1)
			}
			api := gfxapi.Find(atoms[p.After.Index].API())
			s := gfxapi.NewState()
			for _, a := range atoms[:p.After.Index] {
				a.Mutate(s, d, l)
			}
			res, found := s.APIs[api]
			if !found {
				return nil, fmt.Errorf("No state for API '%v'", api.Name())
			}
			return res, nil

		case *path.Field:
			o, err := resolve(p.Struct)
			if err != nil {
				return nil, err
			}
			v := reflect.ValueOf(o)
			switch v.Kind() {
			case reflect.Struct:
				return v.FieldByName(p.Name).Interface(), nil

			default:
				return nil, fmt.Errorf("Cannot access fields of type type %T", o)
			}

		case *path.ArrayIndex:
			o, err := resolve(p.Array)
			if err != nil {
				return nil, err
			}
			v := reflect.ValueOf(o)
			switch v.Kind() {
			case reflect.Array, reflect.Slice, reflect.String:
				return v.Index(int(p.Index)).Interface(), nil

			default:
				return nil, fmt.Errorf("Cannot array-index type %T", o)
			}

		case *path.MapIndex:
			o, err := resolve(p.Map)
			if err != nil {
				return nil, err
			}
			k, err := resolve(p.Key)
			if err != nil {
				return nil, err
			}
			v := reflect.ValueOf(o)
			switch v.Kind() {
			case reflect.Map:
				return v.MapIndex(reflect.ValueOf(k)).Interface(), nil

			default:
				return nil, fmt.Errorf("Cannot map-index type %T", o)
			}

		default:
			return p, nil
		}
	}

	return resolve(p)
}
