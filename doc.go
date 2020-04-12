// Package kanren implements relational symbolic logic
package kanren

import (
	"github.com/GoLangsam/kanren/internal/µ"
	"github.com/GoLangsam/sexpr"
)

type X = µ.X
type V = µ.V
type Goal = µ.Goal
type S = µ.S
type StreamOfStates = µ.StreamOfStates

var (
	Unit       = µ.Unit
	Zero       = µ.Zero
	EmptyState = µ.EmptyState
	cons       = sexpr.Cons

/*
	NewString   = sexpr.NewString
	NewSymbol   = sexpr.NewSymbol
	NewInt      = sexpr.NewInt
	NewFloat    = sexpr.NewFloat
	NewVariable = sexpr.NewVariable

	NewList = sexpr.NewList
*/
)

// ============================================================================
