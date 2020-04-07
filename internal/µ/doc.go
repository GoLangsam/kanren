package µ

import "github.com/GoLangsam/kanren/internal/µ/stat"
import "github.com/GoLangsam/kanren/internal/µ/pipe"

import "github.com/GoLangsam/sexpr"

// V represents a logic variable
type V = *sexpr.Variable

// X represents a symbolic expression
type X = *sexpr.Expression

var (
	Parse = sexpr.Parse
)

var (
	newState = stat.Init
)

type S = pipe.S
type StreamOfStates = pipe.StreamOfStates

var (
	Suspend = pipe.Suspend
	Unit    = pipe.Unit
	mZero   StreamOfStates
)
