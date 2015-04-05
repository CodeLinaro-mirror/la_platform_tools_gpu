package atom

// Transformer is the interface that wraps the basic Transform method.
type Transformer interface {
	// Transform takes a given atom and identifier and Writes out a new atom and
	// identifier to the output atom Writer. Transform must not modify the atom in
	// any way.
	Transform(id ID, atom Atom, output Writer)
	// Flush is called at the end of an atom stream to cause Transformers that
	// cache atoms to send any they have stored into the output.
	Flush(output Writer)
}
