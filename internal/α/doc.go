// Package α implements relational logic using numinal symbolic expressions
package α

import (
	"github.com/GoLangsam/kanren/internal/µ"
	// "github.com/GoLangsam/kanren/internal/α"
	"github.com/GoLangsam/sexpr"
)

type Expression = sexpr.Expression

type Variable = sexpr.Variable

//type Symbol = ast.Symbol
type Atom = sexpr.Atom

func Parse(s string) (*Expression, error) {
	return sexpr.Parse(s)
}

type Element = µ.Element
type Node = µ.Node
type T = µ.T // Term ?
type V = µ.V // Variable

// in Test-Programmen will/man oft speziell erzeugen:

func NewSymbol(s string) *Expression {
	return sexpr.NewSymbol(s)
}

func NewVariable(s string) *Expression {
	return sexpr.NewVariable(s)
}

func NewList(ss ...*Expression) *Expression {
	return sexpr.NewList(ss...)
}

/*

   func Cons(car *SExpr, cdr *SExpr) *SExpr
   func NewFloat(f float64) *Expression
   func NewInt(i int64) *Expression
   func NewString(s string) *Expression
*/
