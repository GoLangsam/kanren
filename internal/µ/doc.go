package µ

//port "github.com/GoLangsam/kanren/internal/µ/bind"
import "github.com/GoLangsam/kanren/internal/µ/pipe"

import "github.com/GoLangsam/sexpr"

// V represents a logic variable
type V = X // *sexpr.Variable

// X represents a symbolic expression
type X = *sexpr.Expression

type S = pipe.S

type SI interface {
	Clone() pipe.S
	Reify(V) X
	Unify(v, w V) bool // used by Equal(X, Y) Goal
	// V() V              // Used by Fresh()

}

var _ SI = pipe.NewS()

type StreamOfStates = pipe.StreamOfStates

var (
	NewS = pipe.NewS
	Unit = pipe.Unit
	ZERO = pipe.ZERO
)

/*
func Unit(s S) StreamOfStates {
	return pipe.Unit(s.(pipe.S))
}

// g.And(h) must use "return g(s).Bind(goal(h))"
func goal(g Goal) func(pipe.S) StreamOfStates {
	return func(s pipe.S) StreamOfStates {
		return g(s)
	}
}
*/
// =============================================================================

// ForEver is a goal that keeps evaluating itself forever.
//
// Inspired by `any` as on page 14 in Byrd_indiana_0093A_10344.pdf
//
// It's implementation highlights the difference between
// a lazy-eval language such as Scheme and
// a Call-by-Value language such as Go.
func (g Goal) ForEver() Goal {
	return func(s S) StreamOfStates {
		return pipe.ForEver(s, g)
	}
}

// =============================================================================
