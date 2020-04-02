package bind

import (
	"github.com/GoLangsam/sexpr"
	"strconv"
)

// from SExpr

//pe Expression = sexpr.Expression
//pe Pair = sexpr.Pair
//pe Atom = sexpr.Atom
//pe Variable = sexpr.Variable

// =========

type V = *sexpr.Variable
type X = *sexpr.Expression

var Cons = sexpr.Cons
var NewSymbol = sexpr.NewSymbol

// TODO: (b *Bind) String() string
// see subs.go; we do not need FMapSs, but we shall import ast

// vAsX returns a singleton expression composed only of v
// vAsX is a helper used by walk
func vAsX(v V) X {
	return &sexpr.Expression{Atom: &sexpr.Atom{Var: v}}
}

// newV provides an expression with nothings but a new anonymous variable.
// newV is a helper for b.Reify.
func (b *Bindings) newV() X {
	n := b.count / len(b.bound)
	s := "_" + strconv.Itoa(n)
	return NewSymbol(s)
}

/*

func (s Substitutions) String() string {
	ss := deriveFmapSs(func(s *Substitution) *ast.SExpr {
		x := ast.Cons(&ast.SExpr{Atom: &ast.Atom{Var: &ast.Variable{Name: s.Var}}}, s.Value)
		return x
	}, []*Substitution(s))
	l := ast.NewList(ss...).String()
	return l[1 : len(l)-1]
}

// Substitution represents a variable and a value.
type Substitution struct {
	Var   string
	Value *ast.SExpr
}

func (s Substitution) String() string {
	x := ast.Cons(&ast.SExpr{Atom: &ast.Atom{Var: &ast.Variable{Name: s.Var}}}, s.Value)
	return x.String()
}

func (b *Bind) asString(v V, x X) string {
ast.Cons(&ast.SExpr{Atom: &ast.Atom{Var: &ast.Variable{Name: s.Var}}}, s.Value)
ast.Cons(&ast.SExpr{Atom: &ast.Atom{Var: &ast.Variable{Name: s.Var}}}, s.Value)
*/
