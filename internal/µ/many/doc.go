package many

import "github.com/GoLangsam/kanren/internal/µ/stat"
/*
import "github.com/GoLangsam/sexpr"

// V represents a logic variable
type V = *sexpr.Variable

// X represents a symbolic expression
type X = *sexpr.Expression

var (
	Parse = sexpr.Parse
)
*/
type S = *stat.E

var (
	newState = stat.Init
)
