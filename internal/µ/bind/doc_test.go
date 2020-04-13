package bind

import (
	"github.com/GoLangsam/sexpr"
)

var (
	NewVariable = sexpr.NewVariable
	NewString   = sexpr.NewString
	NewSymbol   = sexpr.NewSymbol
	NewInt      = sexpr.NewInt
	NewFloat    = sexpr.NewFloat
	NewList     = sexpr.NewList
)
