package bind

// Bindings represents bindings (or "substitutions"):
// any logic variable may be bound to some symbolic expression
// representing its current value.
type Bindings struct {
	bindings
	count int // used to create anonymus variables during reify
}

// New creates fresh and empty Bindings and returns a pointer.
func New() *Bindings {
	return &Bindings{
		bindings: bindings{bound: make(map[V]X)},
	}
}

// Clone provides a clone of b.
func (b *Bindings) Clone() *Bindings {
	cb := b.clone()
	clone := New()
	clone.bindings = *cb
	return clone
}

// Reify ...
func (b *Bindings) Reify(x X) *Bindings {
	xx := b.walkX(x)
	u, isVariable := xx.AsVariable()
	if isVariable {
		b.Bind(u, b.newV())
	} else if xx.IsPair() {
		b.Reify(xx.Car()).Reify(xx.Cdr())
	}
	return b
}
