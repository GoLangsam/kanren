package bind

// Ings represents bindings (or "substitutions"):
// any logic variable may be bound to some symbolic expression
// representing its current value.
//
// Use as `bind.Ings` (pun intended).
//
// The zero value is not useful - initialize with `bind.New()`.
type Ings struct {
	bindings
	count int // used to create anonymus variables during reify
}

// New creates fresh and empty mapping of/for bind.Ings and returns a pointer.
func New() *Ings {
	return &Ings{
		bindings: bindings{bound: make(map[V]X)},
	}
}

// Clone provides a clone of b.
func (b *Ings) Clone() *Ings {
	clone := New()
	clone.bindings = *(b.clone())
	return clone
}

// Reify ...
func (b *Ings) Reify(x X) *Ings {
	xx := b.walkX(x)
	u, isVariable := xx.AsVariable()
	if isVariable { // bind u(=xx) to new fresh var
		b.Bind(u, newVar(b.nextName()))
	} else if xx.IsPair() {
		b.Reify(xx.Car()).Reify(xx.Cdr())
	}
	return b
}
