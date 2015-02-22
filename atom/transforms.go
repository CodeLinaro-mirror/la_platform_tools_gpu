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

// Add is a convenience function for appending the list of Transformer functions
// f to the end of the Transforms list.
func (l *Transforms) AddFunc(f ...func(id Id, atom Atom, output Writer)) {
	for _, f := range f {
		*l = append(*l, transformerFunc(f))
	}
}

type transformerFunc func(id Id, atom Atom, output Writer)

func (f transformerFunc) Transform(id Id, atom Atom, output Writer) {
	f(id, atom, output)
}

type processWriter struct {
	t Transformer
	o Writer
}

func (p processWriter) Write(id Id, a Atom) { p.t.Transform(id, a, p.o) }
