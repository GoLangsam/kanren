package bind

// TODO: does exts need to clone?

// bindings implements the essence.
// Bindings extends bindings with Reifiy, which needs to know how to create anonymous variables.
type bindings struct {
	bound map[V]X
}

// new creates fresh and empty bindings with a ready-to-use bound-map and returns a pointer.
func new() *bindings {
	return &bindings{
		bound: make(map[V]X),
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

// Bind binds x to v, so v is bound to x.
// Thus, (v . x) resembles a substitution pair.
// Note: Bind does not avoid circular bindings.
func (b *bindings) Bind(v V, x X) *bindings {
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

// Subs returns the expression to which v is bound, if any.
//
// This expression shall substitute the variable,
// which shall thus become substituted by its value.
func (b *bindings) Subs(v V) (x X, hasSubs bool) {
	x, hasSubs = b.bound[v]
	return
}

// walkV ...
func (b *bindings) walkV(v V) X {
	xx, found := b.Subs(v)
	if !found {
		return v.Expr()
	}
	vx, isXxVariable := xx.AsVariable()
	if !isXxVariable {
		return xx
	}
	return b.walkV(vx)
}

// walkX ...
func (b *bindings) walkX(x X) X {
	vx, isXxVariable := x.AsVariable()
	if !isXxVariable {
		return x
	}
	xx, found := b.Subs(vx)
	if !found {
		return x
	}
	return b.walkX(xx)
}

// Walk ... some call it `walkstar` or `walk*`
func (b *bindings) Walk(x X) X {
	xx := b.walkX(x)
	if xx.IsVariable() {
		return xx
	}
	if xx.IsPair() {
		return cons(
			b.Walk(xx.Pair.Car),
			b.Walk(xx.Pair.Cdr),
		)
	}
	return xx
}

// Occurs reports whether v occurs in x.
func (b *bindings) Occurs(v V, x X) bool {
	xx := b.walkX(x)
	u, isVariable := xx.AsVariable()
	if isVariable {
		return u.Equal(v)
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
	b.Bind(v, x)
	// append([]*Substitution{&Substitution{Var: v.Name, Value: x}}, s...), true
	return true

}

// Unify returns either (ok = false) or the substitutions extended with zero or more associations,
// where cycles in substitutions can lead to (ok = false)
func (b *bindings) Unify(x, y X) bool {
	xx := x
	vx, isXxVariable := x.AsVariable()
	if isXxVariable {
		xx = b.walkV(vx)
	}

	yy := y
	vy, isYyVariable := y.AsVariable()
	if isYyVariable {
		yy = b.walkV(vy)
	}

	if isXxVariable && isYyVariable && vx.Equal(vy) {
		return true
	}

	if isXxVariable {
		return b.exts(vx, yy)
	}
	if isYyVariable {
		return b.exts(vy, xx)
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
