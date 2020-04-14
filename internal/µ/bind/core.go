package bind

import "github.com/GoLangsam/sexpr"

// V is an eXpression which represents a logic variable
type V = X // *sexpr.Variable

// X represents a symbolic expression
type X = *sexpr.Expression

var (
	cons    = sexpr.Cons    // used by b.String & b.Walk
	newList = sexpr.NewList // used by b.String
)
