// Package α implements relational logic using nominal symbolic expressions
package α

import (
	"github.com/GoLangsam/kanren/internal/µ"
	// "github.com/GoLangsam/kanren/internal/α"
	"github.com/GoLangsam/sexpr"
)

var (
	Parse = sexpr.Parse
)

type V = µ.V
type X = µ.X
