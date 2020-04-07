package bind

// Ings represents bindings (or "substitutions"):
// any logic variable may be bound to some symbolic expression
// representing its current value.
//
// Use as `bind.Ings` (pun intended).
//
// The zero value is not useful - initialize with `bind.New()`.
type Ings struct {
	bound map[V]X
}

// New creates fresh and empty mapping of/for bind.Ings and returns a pointer.
func New() *Ings {
	return &Ings{
		bound: make(map[V]X),
	}
}

/*
// Clone provides a clone of b.
func (bind *Ings) Clone() *Ings {
	clone := New()
	for v, x := range bind.bound {
		clone.bound[v] = x
	}
	return clone
}
*/

// Bind binds x to v, so v is bound to x.
// Thus, (v . x) resembles a substitution pair.
// Note: Bind does not avoid circular bindings.
func (bind *Ings) Bind(v V, x X) *Ings {
	bind.bound[v] = x
	return bind
}

// Drop makes v unbound, reports whether v was bound, and
// returns the expression (if any) v was previously bound with.
func (bind *Ings) Drop(v V) (x X, wasBound bool) {
	x, wasBound = bind.bound[v]
	if wasBound {
		delete(bind.bound, v)
	}
	return
}

// IsBound reports whether v is bound or not
func (bind *Ings) IsBound(v V) (isBound bool) {
	_, isBound = bind.bound[v]
	return
}

// Subs returns the expression to which v is bound, if any.
//
// This expression shall substitute the variable,
// which shall thus become substituted by its value.
func (bind *Ings) Subs(v V) (x X, hasSubs bool) {
	x, hasSubs = bind.bound[v]
	return
}

// walkV ...
func (bind *Ings) walkV(v V) X {
	xx, found := bind.Subs(v)
	if !found {
		return v.Expr()
	}
	vx, isXxVariable := xx.AsVariable()
	if !isXxVariable {
		return xx
	}
	return bind.walkV(vx)
}

// walkX ...
func (bind *Ings) walkX(x X) X {
	vx, isXxVariable := x.AsVariable()
	if !isXxVariable {
		return x
	}
	xx, found := bind.Subs(vx)
	if !found {
		return x
	}
	return bind.walkX(xx)
}

// Walk ... some call it `walkstar` or `walk*`
func (bind *Ings) Walk(x X) X {
	xx := bind.walkX(x)
	if xx.IsVariable() {
		return xx
	}
	if xx.IsPair() {
		return cons(
			bind.Walk(xx.Pair.Car),
			bind.Walk(xx.Pair.Cdr),
		)
	}
	return xx
}

// Occurs reports whether v occurs in x.
func (bind *Ings) Occurs(v V, x X) bool {
	xx := bind.walkX(x)
	u, isVariable := xx.AsVariable()
	if isVariable {
		return u.Equal(v)
	}
	if xx.IsPair() {
		return bind.Occurs(v, xx.Car()) || bind.Occurs(v, xx.Cdr())
	}
	return false
}

// exts attempts to bind v with x and fails
// if such bind would introduce a circle
// due to v occurs in x.
func (bind *Ings) exts(v V, x X) bool {
	if bind.Occurs(v, x) {
		return false
	}
	bind.Bind(v, x)
	return true
}

// Unify extends the bind.Ings with zero or more associations
// in an attempt to see whether the given eXpressions are equal
// and reports its success.
// Circular bindings imply failure (ok = false).
func (bind *Ings) Unify(x, y X) bool {
	xx := x
	vx, isXxVariable := x.AsVariable()
	if isXxVariable {
		xx = bind.walkV(vx)
	}

	yy := y
	vy, isYyVariable := y.AsVariable()
	if isYyVariable {
		yy = bind.walkV(vy)
	}

	if isXxVariable && isYyVariable && vx.Equal(vy) {
		return true
	}

	if isXxVariable {
		return bind.exts(vx, yy)
	}
	if isYyVariable {
		return bind.exts(vy, xx)
	}

	if xx.IsPair() && yy.IsPair() {
		sok := bind.Unify(xx.Car(), yy.Car())
		if !sok {
			return false
		}
		return bind.Unify(xx.Cdr(), yy.Cdr())
	}
	return xx.Equal(yy)
}
