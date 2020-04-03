package micro

import (
	"github.com/GoLangsam/sexpr"
)

// from SExpr

type V = *sexpr.Variable
type X = *sexpr.Expression

//pe Atom = sexpr.Atom
//pe Pair = sexpr.Pair

var (
	Parse = sexpr.Parse
)
