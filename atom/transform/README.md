# transform
--
    import "android.googlesource.com/platform/tools/gpu/atom/transform"

Package transform provides implementations of the atom Transformer interface.

## Usage

#### type EarlyTerminator

```go
type EarlyTerminator struct {
}
```

EarlyTerminator is an implementation of Transformer that will consume all atoms
(except for the EOS atom) once all the atoms passed to Add have passed through
the transformer.

#### func (*EarlyTerminator) Add

```go
func (t *EarlyTerminator) Add(id atom.ID)
```
Add adds the atom with identifier id to the set of atoms that must be seen
before the EarlyTerminator will consume all atoms (excluding the EOS atom).

#### func (*EarlyTerminator) Flush

```go
func (t *EarlyTerminator) Flush(out atom.Writer)
```

#### func (*EarlyTerminator) Transform

```go
func (t *EarlyTerminator) Transform(id atom.ID, a atom.Atom, out atom.Writer)
```

#### type Injector

```go
type Injector struct {
}
```

Injector is an implementation of Transformer that can inject atoms into the atom
stream.

#### func (*Injector) Flush

```go
func (t *Injector) Flush(out atom.Writer)
```

#### func (*Injector) Inject

```go
func (t *Injector) Inject(after atom.ID, a atom.Atom)
```
Inject emits the atom a with identifier id after the atom with identifier after.

#### func (*Injector) Transform

```go
func (t *Injector) Transform(id atom.ID, a atom.Atom, out atom.Writer)
```

#### type SkipDrawCalls

```go
type SkipDrawCalls struct {
}
```

SkipDrawCalls is an implementation of Transformer that skips all draw calls that
have not been explicitly requested.

#### func (*SkipDrawCalls) Draw

```go
func (t *SkipDrawCalls) Draw(id atom.ID)
```
Draw adds an exception to allow all draw calls up to and including the atom with
identifier id for the frame holding the atom.

#### func (*SkipDrawCalls) Flush

```go
func (t *SkipDrawCalls) Flush(out atom.Writer)
```

#### func (*SkipDrawCalls) Transform

```go
func (t *SkipDrawCalls) Transform(id atom.ID, a atom.Atom, out atom.Writer)
```

#### type Trace

```go
type Trace struct {
	Logger log.Logger
}
```

Trace is an implementation of Transformer that records each atom id and atom
value that passes through Trace to Logger. Atoms passing through Trace are
written to the output Writer unaltered.

#### func (Trace) Flush

```go
func (t Trace) Flush(out atom.Writer)
```

#### func (Trace) Transform

```go
func (t Trace) Transform(id atom.ID, a atom.Atom, out atom.Writer)
```
