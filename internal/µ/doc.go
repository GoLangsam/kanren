package µ

import "github.com/GoLangsam/kanren/internal/µ/stat"
import "github.com/GoLangsam/kanren/internal/µ/many"

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

type S = many.S
type StreamOfStates = many.StreamOfStates

var (
	Suspend = many.Suspend
	Unit    = many.Unit
	mZero   = many.Zero
)
