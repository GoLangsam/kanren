package micro

/*
exts either extends a substitution s with an association between the variable x and the value v ,
or it produces (ok = false)s
if extending the substitution with the pair `(,x . ,v) would create a cycle.

scheme code:

	(define (ext-s x v s)
		(cond
			(
				(occurs? x v s)
				#f
			)
			(else
				(cons
					`(,x . ,v)
					s
				)
			)
		)
	)
*/
func (s Substitutions)exts(v V, x *Expression) (Substitutions, bool) {
	if s.occurs(v, x) {
		return nil, false
	}
	return append([]*Substitution{&Substitution{Var: v.Name, Value: x}}, s...), true
}

/*
scheme code:

	(define (occurs? x v s)
		(let
			(
				(v
					(walk v s)
				)
			)
			(cond
				(
					(var? v)
					(eqv? v x)
				)
				(
					(pair? v)
					(or
						(occurs? x (car v) s)
						(occurs? x (cdr v) s)
					)
				)
				(else
					#f
				)
			)
		)
	)
*/
func (s Substitutions) occurs(v V, x *Expression) bool {
	xx := x
	if x.IsVariable() {
		xx = s.walk(x.Atom.Var)
	}
	if xx.IsVariable() {
		return xx.Atom.Var.Equal(v)
	}
	if xx.IsPair() {
		return s.occurs(v, xx.Car()) || s.occurs(v, xx.Cdr())
	}
	return false
}
