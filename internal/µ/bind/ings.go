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

func (bind *Ings) Clone() *Ings {
	clone := New()
	for k, v := range bind.bound {
		clone.bound[k] = v
	}
	return clone
}

// Bind binds x to v, so v is bound to x.
// Thus, (v . x) resembles a substitution pair.
// Note: Bind does not avoid circular bindings.
// Use Occurs to check beforehand.
// Bind is a nOp if v or x are nil or v is not a Variable.
func (bind *Ings) Bind(v V, x X) *Ings {
	if v != nil && v.Atom.Var != nil && x != nil {
		bind.bound[v] = x
	}
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

// Resolve the eXpression
// along the bindings
// down to the first non-Variable eXpression
// or down to the first unbound eXpression
func (bind *Ings) Resolve(x X) X {
	if x.IsVariable() && bind.bound[x] != nil {
		return bind.Resolve(bind.bound[x])
	}
	return x
}

// Walk ... some call it `walkstar` or `walk*`
func (bind *Ings) Walk(x X) X {
	xx := bind.Resolve(x)
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
	x = bind.Resolve(x)
	u, isVariable := x.AsVariable()
	if isVariable {
		return u.Equal(v.Atom.Var)
	}
	if x.IsPair() {
		return bind.Occurs(v, x.Car()) || bind.Occurs(v, x.Cdr())
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
	x = bind.Resolve(x)
	y = bind.Resolve(y)

	vx, isXxVariable := x.AsVariable()
	vy, isYyVariable := y.AsVariable()

	if isXxVariable && isYyVariable && vx.Equal(vy) {
		return true
	}

	if isXxVariable {
		return bind.exts(x, y)
	}
	if isYyVariable {
		return bind.exts(y, x)
	}

	if x.IsPair() && y.IsPair() {
		uniCars := bind.Unify(x.Car(), y.Car())
		if uniCars {
			return bind.Unify(x.Cdr(), y.Cdr())
		}
		return false
	}
	return x.Equal(y)
}
