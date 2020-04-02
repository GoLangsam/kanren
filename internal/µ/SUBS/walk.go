package micro

/*
walk has been modified to assume it is getting a variable.

Here is the original scheme code:

	(define (walk u s)
		(let (
			(pr (and
					(var? u)
					(assp (λ (v) (var=? u v)) s)
			)))
			(if pr (walk (cdr pr) s) u)
		)
	)
*/
func (s Substitutions) walk(v V) *Expression {
	a, ok := assv(v, s)
	if !ok {
		return &Expression{Atom: &Atom{Var: v}}
	}
	if !a.Value.IsVariable() {
		return a.Value
	}
	return s.walk(a.Value.Atom.Var)
}

/*
assv either produces the first association in s that has v as its car using eqv,
or produces ok = false if l has no such association.

for example:

	assv v s <==> (assp (λ (v) (var=? u v)) s))
*/
func assv(v V, ss *Substitutions) (*Substitution, bool) {
	if ss == nil {
		return nil, false
	}
	for i, s := range ss {
		if v.Name == s.Var {
			return ss[i], true
		}
	}
	return nil, false
}

/*
scheme code:

	(define (walkStar v s)
		(let
			(
				(v (walk v s))
			)
			(cond
				(
					(var? v)
					v
				)
				(
					(pair? v)
					(cons
						(walkStar (car v) s)
						(walkStar (cdr v) s)
					)
				)
				(else v)
			)
		)
	)
*/
func (s Substitutions) walkStar(v *Expression) *Expression {
	vv := v
	if v.IsVariable() {
		vv = s.walk(v.Atom.Var)
	}
	if vv.IsVariable() {
		return vv
	}
	if vv.IsPair() {
		carv := vv.Pair.Car
		cdrv := vv.Pair.Cdr
		wcar := s.walkStar(carv)
		wcdr := s.walkStar(cdrv)
		w := ast.Cons(wcar, wcdr)
		return w
	}
	return vv
}
