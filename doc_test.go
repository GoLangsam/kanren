package kanren

import (
	"github.com/GoLangsam/kanren/internal/µ"
	"github.com/GoLangsam/sexpr"
)

var (
	EmptyState = µ.NewS

	Parse = sexpr.Parse

	// NewVariable = sexpr.NewVariable must NOT be used! Use Fresh instead!

	NewString = sexpr.NewString
	NewSymbol = sexpr.NewSymbol
	NewInt    = sexpr.NewInt
	NewFloat  = sexpr.NewFloat
	NewList   = sexpr.NewList
)
