package atom

// Transformer is the interface that wraps the basic Transform method.
//
// Transform takes a given atom and identifier and Writes out a new atom and
// identifier to the output atom Writer. Transform must not modify the atom in
// any way.
type Transformer interface {
	Transform(id ID, atom Atom, output Writer)
}
