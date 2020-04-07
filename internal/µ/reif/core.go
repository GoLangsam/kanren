package reif

import "github.com/GoLangsam/sexpr"

// V represents a logic variable
type V = *sexpr.Variable

// X represents a symbolic expression
type X = *sexpr.Expression

var (
	//cons   = sexpr.Cons        // used by b.Walk
	newVar = sexpr.NewVariable // used by b.Reify
)
