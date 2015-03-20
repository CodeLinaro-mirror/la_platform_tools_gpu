package atom

// Transforms is a list of Transformer objects.
type Transforms []Transformer

// Transform sequentially transforms the atoms by each of the transformers in
// the list, before writing the final output to the output atom Writer.
func (l Transforms) Transform(atoms List, out Writer) {
	chain := out
	for i := len(l) - 1; i >= 0; i-- {
		chain = processWriter{l[i], chain}
	}
	atoms.WriteTo(chain)
}

// Add is a convenience function for appending the list of Transformers t to the
// end of the Transforms list.
func (l *Transforms) Add(t ...Transformer) {
	*l = append(*l, t...)
}

// Transform is a helper for building simple Transformers that are implemented
// by function f. name is used to identify the transform when logging.
func Transform(name string, f func(id ID, atom Atom, output Writer)) Transformer {
	return transform{name, f}
}

type transform struct {
	N string                                // Transform name. Used for debugging.
	F func(id ID, atom Atom, output Writer) // The transform function.
}

func (t transform) Transform(id ID, atom Atom, output Writer) {
	t.F(id, atom, output)
}

type processWriter struct {
	t Transformer
	o Writer
}

func (p processWriter) Write(id ID, a Atom) { p.t.Transform(id, a, p.o) }
