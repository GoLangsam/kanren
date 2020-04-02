package micro

import (
	"github.com/GoLangsam/kanren/internal/µ/core"
	// "github.com/GoLangsam/kanren/internal/?"
	"github.com/GoLangsam/sexpr"
)

// from SExpr

type Expression = sexpr.Expression
//pe Pair = sexpr.Pair
type Atom = sexpr.Atom
type Variable = sexpr.Variable

var Parse = sexpr.Parse

// from Core

var Init = kanren.Init

type State = kanren.S
//type Goal func(*State) StreamOfStates

// from SExpr - für Test-Programme

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
