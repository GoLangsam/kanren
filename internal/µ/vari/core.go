package vari

import "github.com/GoLangsam/sexpr"

// V represents a logic variable
type V = *sexpr.Variable

// X represents a symbolic expression
type X = *sexpr.Expression

var (
	NewVariable = sexpr.NewVariable
)
