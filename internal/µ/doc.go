package µ

import "github.com/GoLangsam/kanren/internal/µ/pipe"

import "github.com/GoLangsam/sexpr"

// V represents a logic variable
type V = X // *sexpr.Variable

// X represents a symbolic expression
type X = *sexpr.Expression

type S = pipe.S
type StreamOfStates = pipe.StreamOfStates

var (
	NewS = pipe.NewS
	Unit = pipe.Unit
	Zero = pipe.Zero
)
