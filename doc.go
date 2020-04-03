// Package kanren implements relational symbolic logic
package kanren

import (
	"github.com/GoLangsam/kanren/internal/µ"
	"github.com/GoLangsam/sexpr"
)

type Expression = sexpr.Expression

type Variable = sexpr.Variable

//type Symbol = ast.Symbol
type Atom = sexpr.Atom

var (
	Parse = sexpr.Parse

	Cons = sexpr.Cons

	NewString   = sexpr.NewString
	NewSymbol   = sexpr.NewSymbol
	NewInt      = sexpr.NewInt
	NewFloat    = sexpr.NewFloat
	NewVariable = sexpr.NewVariable

	NewList = sexpr.NewList
)

// ============================================================================

type Goal = µ.Goal
