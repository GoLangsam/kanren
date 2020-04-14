// Package kanren implements relational symbolic logic
package kanren

import (
	"github.com/GoLangsam/kanren/internal/µ"
	"github.com/GoLangsam/sexpr"
)

type X = µ.X

//type V = µ.V

type Goal func(S) StreamOfStates

type S = µ.S
type StreamOfStates = µ.StreamOfStates

var (
	Unit = µ.Unit
	Zero = µ.Zero
	cons = sexpr.Cons
)
