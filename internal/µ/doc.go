package µ

import "github.com/GoLangsam/kanren/internal/µ/stat"
import "github.com/GoLangsam/kanren/internal/µ/pipe"

import "github.com/GoLangsam/sexpr"

// V represents a logic variable
type V = X // *sexpr.Variable

// X represents a symbolic expression
type X = *sexpr.Expression

var (
	EmptyState = stat.Init
)

type S = pipe.S
type StreamOfStates = pipe.StreamOfStates

var (
	Suspend = pipe.Suspend
	Unit    = pipe.Unit
	Zero    = pipe.Zero
)
