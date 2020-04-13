package reif

import (
	"github.com/GoLangsam/sexpr"
)

var (
	Cons  = sexpr.Cons
	Parse = sexpr.Parse

	// NewVariable = sexpr.NewVariable must NOT be used! Use Fresh instead!

	NewString = sexpr.NewString
	NewSymbol = sexpr.NewSymbol
	NewInt    = sexpr.NewInt
	NewFloat  = sexpr.NewFloat
	NewList   = sexpr.NewList
)
