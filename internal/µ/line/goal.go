package micro

// import "strconv"

type State struct{}

func EmptyState() *State {
	return &State{}
}

func (State) Unify(x, y X) (xUy X, ok bool) {
	return
}

// end of fakes

// Goal is a function that takes a state and returns a stream of states.
type Goal func(*State) StreamOfStates

/*
RunGoal calls a goal with an emptystate and n possible resulting states.

scheme code:

	(define (run-goal n g)
		(takeInf n (g empty-s))
	)

If n == -1 then all possible states are returned.
*/
func RunGoal(n int, g Goal) []*State {
	ss := g(EmptyState())
	return takeStream(n, ss)
}

// SuccessO is a goal that always returns the input state in the resulting stream of states.
func SuccessO() Goal {
	return func(s *State) StreamOfStates {
		return NewSingletonStream(s)
	}
}

// FailureO is a goal that always returns an empty stream of states.
func FailureO() Goal {
	return func(s *State) StreamOfStates {
		return nil
	}
}

/*
EqualO returns a Goal that unifies the input expressions in the output stream.

scheme code:

	(define (= u v)
		(lambda_g (s/c)
			(let ((s (unify u v (car s/c))))
				(if s (unit `(,s . ,(cdr s/c))) mzero)
			)
		)
	)
*/
func EqualO(u, v X) Goal {
	return func(s *State) StreamOfStates {
		ss, sok := s.Unify(u, v)
		if sok {
			s, _ := &State{}, ss // TODO: &State{Substitutions: ss, Counter: s.Counter}
			return NewSingletonStream(s)
		}
		return nil
	}
}

/*
NeverO is a Goal that returns a never ending stream of suspensions.

scheme code:

	(define (nevero)
		(lambda (s)
			(lambda ()
				((nevero) s)
			)
		)
	)
*/
func NeverO() Goal {
	return func(s *State) StreamOfStates {
		return Suspension(func() StreamOfStates {
			return NeverO()(s)
		})
	}
}

/*
AlwaysO is a goal that returns a never ending stream of success.

scheme code:

	(define (alwayso)
		(lambda (s)
			(lambda ()
				(
					(disj
						S
						(alwayso)
					)
					s
				)
			)
		)
	)
*/
func AlwaysO() Goal {
	return func(s *State) StreamOfStates {
		return Suspension(func() StreamOfStates {
			return DisjointO(
				SuccessO(),
				AlwaysO(),
			)(s)
		})
	}
}
