// Package kanren implements relational symbolic logic
package kanren

//go:generate doc

import (
	"github.com/GoLangsam/kanren/internal/µ"
	"github.com/GoLangsam/sexpr"
)

type X = µ.X

type V = µ.V

type Goal = µ.Goal // func(S) StreamOfStates

type S = µ.S
type StreamOfStates = µ.StreamOfStates

var (
	FAIL Goal = µ.Failure() // FAIL represents Failure.
	GOAL Goal = µ.Success() // GOAL represents Success.

	NewS                = µ.NewS // only used in test programs
	Unit                = µ.Unit
	ZERO StreamOfStates = µ.ZERO // used by Equal-relation

	cons    = sexpr.Cons
	NewList = sexpr.NewList
)
