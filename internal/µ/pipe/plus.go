package pipe

/*
scheme code:

	(define (mplus $1 $2)
		(cond
			((null? $1) $2)
			((procedure? $1) (?_$ () (mplus $2 ($1))))
			(else (cons (car $1) (mplus (cdr $1) $2)))
		)
	)

An alternative implementation could swap the goals for a more fair distribution:

scheme code:

	(define (mplus $1 $2)
		(cond
			((null? $1) $2)
			((procedure? $1) (?_$ () (mplus $2 ($1))))
			(else (cons (car $1) (mplus $2 (cdr $1))))
		)
	)
*/
func (s StreamOfStates) Plus(ss StreamOfStates) StreamOfStates {
	if s == nil {
		return ss
	}
	head := s.Head()
	if head != nil { // not a suspension => procedure? == false
		return Prepend(head, func() StreamOfStates {
			return ss.Plus(s)
		})
	}
	return Suspend(func() StreamOfStates {
		return ss.Plus(s)
	})
}
