package reif

import "github.com/GoLangsam/sexpr"

// V is an eXpression which represents a logic variable
type V = X // *sexpr.Variable

// X represents a symbolic expression
type X = *sexpr.Expression
