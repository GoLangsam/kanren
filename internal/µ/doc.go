package µ

import "github.com/GoLangsam/kanren/internal/µ/pipe"

import "github.com/GoLangsam/sexpr"

// V represents a logic variable
type V = X // *sexpr.Variable

// X represents a symbolic expression
type X = *sexpr.Expression

type S interface {
	Clone() pipe.S
	Walk(X) X // for Reify
	// Reify(V, ) []X

	Unify(v, w V) bool // used by Equal(X, Y) Goal
	V() V              // Used by Fresh()

}

type StreamOfStates = pipe.StreamOfStates

var (
	NewS = pipe.NewS
	//Unit = pipe.Unit
	Zero = pipe.Zero
)

func Unit(s S) StreamOfStates {
	return pipe.Unit(s.(pipe.S))
}

// g.And(h) must use "return g(s).Bind(goal(h))"
func goal(g Goal) func(pipe.S) StreamOfStates {
	return func(s pipe.S) StreamOfStates {
		return g(s)
	}
}
