package µ

import "github.com/GoLangsam/sexpr"

var (
	Parse = sexpr.Parse
)

type V = *sexpr.Variable
type X = *sexpr.Expression
