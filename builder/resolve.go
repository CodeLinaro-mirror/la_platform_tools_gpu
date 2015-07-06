package builder

import (
	"fmt"
	"reflect"

	"android.googlesource.com/platform/tools/gpu/database"
	"android.googlesource.com/platform/tools/gpu/gfxapi"
	"android.googlesource.com/platform/tools/gpu/log"
	"android.googlesource.com/platform/tools/gpu/service"
	"android.googlesource.com/platform/tools/gpu/service/path"
)

// Resolve resolves and returns the object, value or memory at the path p.
func Resolve(p path.Path, d database.Database, l log.Logger) (interface{}, error) {
	v, err := resolveChain(path.Flatten(p), d, l)
	if err != nil {
		return nil, err
	}
	return v[len(v)-1], nil
}

func resolveChain(paths []path.Path, d database.Database, l log.Logger) ([]interface{}, error) {
	v := make([]interface{}, len(paths))
	for i, p := range paths {
		switch p := p.(type) {
		case *path.Capture:
			capture, err := service.ResolveCapture(service.CaptureId{ID: p.ID}, d, l)
			if err != nil {
				return nil, err
			}
			v[i] = &capture

		case *path.Atoms:
			capture := v[i-1].(*service.Capture)
			atoms, err := loadAtoms(capture.Atoms, d, l)
			if err != nil {
				return nil, err
			}
			v[i] = &service.AtomStream{Atoms: atoms}

		case *path.Atom:
			atoms := v[i-1].(*service.AtomStream).Atoms
			if p.Index >= uint64(len(atoms)) {
				return nil, fmt.Errorf("Atom at %s is out of bounds [0-%d]",
					p.Path(), len(atoms)-1)
			}
			v[i] = atoms[p.Index]

		case *path.State:
			atoms := v[i-2].(*service.AtomStream).Atoms
			api := gfxapi.Find(atoms[p.After.Index].API())
			if api == nil {
				return nil, fmt.Errorf("Atom at %s has no API",
					paths[i-1].Path())
			}
			s := gfxapi.NewState()
			for _, a := range atoms[:p.After.Index] {
				a.Mutate(s, d, l)
			}
			res, found := s.APIs[api]
			if !found {
				return nil, fmt.Errorf("No state for API '%v' after %v",
					api.Name(), paths[i-1].Path())
			}
			v[i] = res

		case *path.Field:
			s := reflect.ValueOf(v[i-1])
		deref:
			for {
				switch s.Kind() {
				case reflect.Struct:
					f := s.FieldByName(p.Name)
					if !f.IsValid() {
						return nil, fmt.Errorf("Struct of type %v at %s does not have field '%s'",
							s.Type(), paths[i-1].Path(), p.Name)
					}
					v[i] = s.FieldByName(p.Name).Interface()
					break deref
				case reflect.Interface, reflect.Ptr:
					if s.IsNil() {
						return nil, fmt.Errorf("Pointer at %s is nil", paths[i-1].Path())
					}
					s = s.Elem()
				default:
					return nil, fmt.Errorf("Type %v at %s is not a struct",
						s.Type(), paths[i-1].Path())
				}
			}

		case *path.ArrayIndex:
			a := reflect.ValueOf(v[i-1])
			switch a.Kind() {
			case reflect.Array, reflect.Slice, reflect.String:
				if int(p.Index) >= a.Len() {
					return nil, fmt.Errorf("Index at %s is out of bounds [0-%d]", p.Path(), a.Len()-1)
				}
				v[i] = a.Index(int(p.Index)).Interface()

			default:
				return nil, fmt.Errorf("Type %v at %s is not an array, slice or string",
					a.Type(), paths[i-1].Path())
			}

		case *path.MapIndex:
			m := reflect.ValueOf(v[i-1])
			switch m.Kind() {
			case reflect.Map:
				key := reflect.ValueOf(p.Key)
				if key.Type() != m.Type().Key() {
					return nil, fmt.Errorf("Map at %s has key of type %v, got type %v",
						paths[i-1].Path(), m.Type().Key(), key.Type())
				}
				val := m.MapIndex(reflect.ValueOf(p.Key))
				if !val.IsValid() {
					return nil, fmt.Errorf("Map at %s does not contain key %v",
						paths[i-1].Path(), p.Key)
				}
				v[i] = val.Interface()

			default:
				return nil, fmt.Errorf("Type %v at %s is not a map",
					m.Type(), paths[i-1].Path())
			}

		default:
			return nil, fmt.Errorf("Unknown path type %T", p)
		}
	}

	return v, nil
}
