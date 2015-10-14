package builder

import (
	"fmt"
	"reflect"

	"android.googlesource.com/platform/tools/gpu/atom"
	"android.googlesource.com/platform/tools/gpu/binary"
	"android.googlesource.com/platform/tools/gpu/database"
	"android.googlesource.com/platform/tools/gpu/gfxapi"
	"android.googlesource.com/platform/tools/gpu/image"
	"android.googlesource.com/platform/tools/gpu/log"
	"android.googlesource.com/platform/tools/gpu/memory"
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

// ResolveCapture resolves and returns the capture from the path p.
func ResolveCapture(p *path.Capture, d database.Database, l log.Logger) (service.Capture, error) {
	if res, err := Resolve(p, d, l); err == nil {
		return *res.(*service.Capture), nil
	} else {
		return service.Capture{}, err
	}
}

// ResolveAtoms resolves and returns the atom list from the path p.
func ResolveAtoms(p *path.Atoms, d database.Database, l log.Logger) ([]atom.Atom, error) {
	if res, err := Resolve(p, d, l); err == nil {
		return res.(*atom.List).Atoms, nil
	} else {
		return nil, err
	}
}

// ResolveAtom resolves and returns the atom from the path p.
func ResolveAtom(p *path.Atom, d database.Database, l log.Logger) (atom.Atom, error) {
	if res, err := Resolve(p, d, l); err == nil {
		return res.(atom.Atom), nil
	} else {
		return nil, err
	}
}

func resolveChain(paths []path.Path, d database.Database, l log.Logger) ([]interface{}, error) {
	v := make([]interface{}, len(paths))
	for i, p := range paths {
		switch p := p.(type) {
		case *path.Capture:
			capture, err := service.ResolveCapture(p.ID, d, l)
			if err != nil {
				return nil, err
			}
			v[i] = capture

		case *path.Device:
			device, err := service.ResolveDevice(p.ID, d, l)
			if err != nil {
				return nil, err
			}
			v[i] = device

		case *path.As:
			o := v[i-1]
			if c, ok := o.(path.Converter); ok {
				o, err := c.Convert(p, d, l)
				if err != nil {
					return nil, err
				}
				v[i] = o
			} else {
				return nil, fmt.Errorf("Object of type %T at %s does not support converting",
					o, paths[i-1])
			}

		case *path.ImageInfo:
			r, err := service.ResolveImageInfo(p.ID, d, l)
			if err != nil {
				return nil, err
			}
			v[i] = r

		case *path.TimingInfo:
			r, err := service.ResolveTimingInfo(p.ID, d, l)
			if err != nil {
				return nil, err
			}
			v[i] = r

		case *path.Blob:
			if blob, err := database.Resolve(p.ID, d, l); err != nil {
				return nil, err
			} else if data, ok := blob.([]byte); !ok {
				return nil, fmt.Errorf("Path %s gave %T, expected []byte", p, blob)
			} else {
				v[i] = data
			}

		case *path.Atoms:
			capture := v[i-1].(*service.Capture)
			atoms, err := database.Resolve(binary.ID(capture.Atoms), d, l)
			if err != nil {
				return nil, err
			}
			v[i] = atoms

		case *path.Report:
			report, err := database.Build(&BuildReport{Capture: p.Capture, Device: p.Device}, d, l)
			if err != nil {
				return nil, err
			}
			v[i] = report.(*service.Report)

		case *path.Hierarchy:
			root, err := database.Build(&GetHierarchy{Capture: p.Capture}, d, l)
			if err != nil {
				return nil, err
			}
			v[i] = root.(*atom.Group)

		case *path.Atom:
			atoms := v[i-1].(*atom.List).Atoms
			if p.Index >= uint64(len(atoms)) {
				return nil, fmt.Errorf("Atom at %s is out of bounds [0-%d]",
					p.Path(), len(atoms)-1)
			}
			v[i] = atoms[p.Index]

		case *path.State:
			atoms := v[i-2].(*atom.List).Atoms
			if p.After.Index >= uint64(len(atoms)) {
				return nil, fmt.Errorf("%v is out of bounds. [0-%d]", p, len(atoms)-1)
			}
			api := gfxapi.Find(atoms[p.After.Index].API())
			if api == nil {
				return nil, fmt.Errorf("Atom at %s has no API", paths[i-1].Path())
			}
			s := gfxapi.NewState()
			for _, a := range atoms[:p.After.Index+1] {
				a.Mutate(s, d, l)
			}
			res, found := s.APIs[api]
			if !found {
				return nil, fmt.Errorf("No state for API '%v' after %s",
					api.Name(), paths[i-1].Path())
			}
			v[i] = res

		case *path.Resources:
			resources, err := database.Build(&GetResources{Capture: p.Capture}, d, l)
			if err != nil {
				return nil, err
			}
			v[i] = resources.(*service.Resources)

		case *path.Resource:
			resource, err := database.Build(&GetResourceData{Path: p}, d, l)
			if err != nil {
				return nil, err
			}
			v[i] = resource

		case *path.Thumbnail:
			img, err := resolveThumbnail(v[i-1], p, d, l)
			if err != nil {
				return nil, err
			}
			v[i] = img

		case *path.MemoryRange:
			atoms := v[i-2].(*atom.List).Atoms
			if p.After.Index >= uint64(len(atoms)) {
				return nil, fmt.Errorf("%v is out of bounds. [0-%d]", p, len(atoms)-1)
			}

			res, err := ResolveMemoryRange(
				atoms,
				atom.ID(p.After.Index),
				memory.PoolID(p.Pool),
				memory.Range{Base: p.Address, Size: p.Size},
				d, l)
			if err != nil {
				return nil, err
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

		case *path.Slice:
			a := reflect.ValueOf(v[i-1])
			switch a.Kind() {
			case reflect.Array, reflect.Slice, reflect.String:
				if int(p.Start) >= a.Len() || int(p.End) > a.Len() {
					return nil, fmt.Errorf("Slice at %s is out of bounds [0-%d]", p.Path(), a.Len()-1)
				}
				v[i] = a.Slice(int(p.Start), int(p.End)).Interface()

			default:
				return nil, fmt.Errorf("Type %v at %s is not an array, slice or string",
					a.Type(), paths[i-1].Path())
			}

		case *path.MapIndex:
			m := reflect.ValueOf(v[i-1])
			switch m.Kind() {
			case reflect.Map:
				key, ok := convert(reflect.ValueOf(p.Key), m.Type().Key())
				if !ok {
					return nil, fmt.Errorf("Map at %s has key of type %v, got type %v",
						paths[i-1].Path(), m.Type().Key(), key.Type())
				}
				val := m.MapIndex(key)
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

func resolveThumbnail(v interface{}, p *path.Thumbnail, d database.Database, l log.Logger) (*image.Info, error) {
	t, ok := v.(image.Thumbnailer)
	if !ok {
		return nil, fmt.Errorf("Type %T does not support thumbnailing", v)
	}

	img, err := t.Thumbnail(p.DesiredMaxWidth, p.DesiredMaxHeight, d, l)
	if err != nil {
		return nil, err
	}

	if p.DesiredFormat != nil {
		// Convert the image to the desired format.
		if f, ok := p.DesiredFormat.(image.Format); ok {
			if img.Format.Key() != f.Key() {
				img, err = img.ConvertTo(f, d, l)
				if err != nil {
					return nil, err
				}
			}
		} else {
			return nil, fmt.Errorf("Cannot convert thumbnail to format %v", f)
		}
	}

	if _, ok := img.Format.(image.Resizer); !ok {
		return nil, fmt.Errorf("Image format %v does not support resizing", img.Format)
	}

	// Image format supports resizing. See if the image should be.
	scaleX, scaleY := float32(1), float32(1)
	if p.DesiredMaxWidth > 0 && img.Width > p.DesiredMaxWidth {
		scaleX = float32(p.DesiredMaxWidth) / float32(img.Width)
	}
	if p.DesiredMaxHeight > 0 && img.Height > p.DesiredMaxHeight {
		scaleY = float32(p.DesiredMaxHeight) / float32(img.Height)
	}
	scale := scaleX // scale := min(scaleX, scaleY)
	if scale > scaleY {
		scale = scaleY
	}

	targetWidth := uint32(float32(img.Width) * scale)
	targetHeight := uint32(float32(img.Height) * scale)

	// Prevent scaling to zero size.
	if targetWidth == 0 {
		targetWidth = 1
	}
	if targetHeight == 0 {
		targetHeight = 1
	}

	if targetWidth == img.Width && targetHeight == img.Height {
		// Image is already at requested target size.
		return img, err
	}

	return img.Resize(targetWidth, targetHeight, d, l)
}

func convert(val reflect.Value, ty reflect.Type) (reflect.Value, bool) {
	if valTy := val.Type(); valTy != ty {
		if valTy.ConvertibleTo(ty) {
			val = val.Convert(ty)
		} else {
			return val, false
		}
	}
	return val, true
}
