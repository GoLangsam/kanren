package micro

/*
unify returns either (ok = false) or the substitution s extended with zero or more associations,
where cycles in substitutions can lead to (ok = false)

scheme code:

	(define (unify u v s)
		(let ((u (walk u s)) (v (walk v s)))
			(cond
				((and (var? u) (var? v) (var=? u v)) s)
				((var? u) (ext-s u v s))
				((var? v) (ext-s v u s))
				((and (pair? u) (pair? v))
					(let
						((s (unify (car u) (car v) s)))
						(and s (unify (cdr u) (cdr v) s))
					)
				)
				(else (and (eqv? u v) s))
			)
		)
	)
*/
func (s Substitutions)unify(u, v *Expression) (Substitutions, bool) {
	uu := u
	if u.IsVariable() {
		uu = s.walk(u.Atom.Var)
	}
	vv := v
	if v.IsVariable() {
		vv = s.walk(v.Atom.Var)
	}
	if uu.IsVariable() && vv.IsVariable() && uu.Atom.Var.Equal(uu.Atom.Var) {
		return s, true
	}
	if uu.IsVariable() {
		return s.exts(uu.Atom.Var, vv)
	}
	if vv.IsVariable() {
		return s.exts(vv.Atom.Var, uu)
	}
	if uu.IsPair() && vv.IsPair() {
		scar, sok := s.unify(uu.Car(), vv.Car())
		if !sok {
			return nilSubs(), false
		}
		return scar.unify(uu.Cdr(), vv.Cdr())
	}
	if uu.Equal(vv) {
		return s, true
	}
	return nilSubs(), false
}
