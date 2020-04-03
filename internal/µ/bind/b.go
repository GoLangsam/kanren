package bind

// TODO: does exts need to clone?

// bindings implements the essence.
// Bindings extends bindings with Reifiy, which needs to know how to create anonymous variables.
type bindings struct {
	bound map[V]X
}

// Bindings represents bindings (or "substitutions"):
// any logic variable may be bound to some symbolic expression
// representing its current value.
type Bindings struct {
	bindings
	count int // used to create anonymus variables during reify
}

// new creates fresh and empty bindings with a ready-to-use bound-map and returns a pointer.
func new() *bindings {
	return &bindings{
		bound: make(map[V]X),
	}
}

// New creates fresh and empty Bindings and returns a pointer.
func New() *Bindings {
	return &Bindings{
		bindings: bindings{bound: make(map[V]X)},
	}
}

// clone provides a clone of b.
func (b *bindings) clone() *bindings {
	clone := new()
	for v, x := range b.bound {
		clone.bound[v] = x
	}
	return clone
}

// Clone provides a clone of b.
func (b *Bindings) Clone() *Bindings {
	cb := b.clone()
	clone := New()
	clone.bindings = *cb
	return clone
}

// bind binds x to v, so v is bound to x.
// Thus, (v . x) resembles a substitution pair.
// Note: Bind does not avoid circular bindings.
func (b *bindings) bind(v V, x X) *bindings {
	b.bound[v] = x
	return b
}

// Drop makes v unbound, reports whether v was bound, and
// returns the expression (if any) v was previously bound with.
func (b *bindings) Drop(v V) (x X, wasBound bool) {
	x, wasBound = b.bound[v]
	if wasBound {
		delete(b.bound, v)
	}
	return
}

// IsBound reports whether v is bound or not
func (b *bindings) IsBound(v V) (isBound bool) {
	_, isBound = b.bound[v]
	return
}

// Subs return the expression to which v is bound, if any.
func (b *bindings) Subs(v V) (x X, hasSubs bool) {
	x, hasSubs = b.bound[v]
	return
}

// walk ...
func (b *bindings) walkV(v V) X {
	x, ok := b.Subs(v)
	if !ok {
		return v.Expr()
	}
	if !x.IsVariable() {
		return x
	}
	return b.walkV(x.Atom.Var)
}

// walkX - not used
func (b *bindings) walkX(x X) X {
	if !x.IsVariable() {
		return x
	} 
	xx, found := b.Subs(x.Atom.Var)
	if !found {
		return x
	}
	return b.walkX(xx)
}

// Walk ...
func (b *bindings) Walk(x X) X {
	xx := x
	if x.IsVariable() {
		xx = b.walkV(x.Atom.Var)
	}
	if xx.IsVariable() {
		return xx
	}
	if xx.IsPair() {
		return Cons(
			b.Walk(xx.Pair.Car),
			b.Walk(xx.Pair.Cdr),
		)
	}
	return xx
}

// occurs reports whether v occurs in x.
func (b *bindings) Occurs(v V, x X) bool {
	xx := x
	if x.IsVariable() {
		xx = b.walkV(x.Atom.Var)
	}
	if xx.IsVariable() {
		return xx.Atom.Var.Equal(v)
	}
	if xx.IsPair() {
		return b.Occurs(v, xx.Car()) || b.Occurs(v, xx.Cdr())
	}
	return false
}

// exts attempts to bind v with x and fails
// if such bind would introduce a circle
// due to v occurs in x.
func (b *bindings) exts(v V, x X) bool {
	if b.Occurs(v, x) {
		return false
	}
	b.bind(v, x)
	// append([]*Substitution{&Substitution{Var: v.Name, Value: x}}, s...), true
	return true

}

// Unify returns either (ok = false) or the substitutions extended with zero or more associations,
// where cycles in substitutions can lead to (ok = false)
func (b *bindings) Unify(x, y X) bool {
	xx := x
	if x.IsVariable() {
		xx = b.walkV(x.Atom.Var)
	}
	yy := y
	if y.IsVariable() {
		yy = b.walkV(y.Atom.Var)
	}

	if xx.IsVariable() && yy.IsVariable() && xx.Atom.Var.Equal(xx.Atom.Var) {
		return true
	}

	if xx.IsVariable() {
		return b.exts(xx.Atom.Var, yy)
	}
	if yy.IsVariable() {
		return b.exts(yy.Atom.Var, xx)
	}

	if xx.IsPair() && yy.IsPair() {
		sok := b.Unify(xx.Car(), yy.Car())
		if !sok {
			return false
		}
		return b.Unify(xx.Cdr(), yy.Cdr())
	}
	if xx.Equal(yy) {
		return true
	}
	return false
}

// Reify ...
func (b *Bindings) Reify(x X) *Bindings {
	xx := x
	if x.IsVariable() {
		xx = b.walkV(x.Atom.Var)
	}
	if xx.IsVariable() {
		b.bind(xx.Atom.Var, b.newV())
	} else if xx.IsPair() {
		b.Reify(xx.Car()).Reify(xx.Cdr())
	}
	return b
}
