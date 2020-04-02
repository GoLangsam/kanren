package micro

// was first part of goal.go

import (
//	"strconv"

	"github.com/awalterschulze/gominikanren/sexpr/ast"
)

/*
// State is a product of a list of substitutions and a variable counter.
type State struct {
	Substitutions
	Counter int
}

// String returns a string representation of State.
func (s *State) String() string {
	if s.Substitutions == nil {
		return "(" + "()" + " . " + strconv.Itoa(s.Counter) + ")"
	}
	return "(" + s.Substitutions.String() + " . " + strconv.Itoa(s.Counter) + ")"
}
*/

// EmptyState returns an empty state.
func EmptyState() *State {
	return &State{}
}

// Substitutions is a list of substitutions represented by a sexprs pair.
// type Substitutions []*Substitution

func (s Substitutions) String() string {
	ss := deriveFmapSs(func(s *Substitution) *Expression {
		return ast.Cons(&ast.SExpr{Atom: &ast.Atom{Var: &ast.Variable{Name: s.Var}}}, s.Value)
	}, []*Substitution(s))
	l := ast.NewList(ss...).String()
	return l[1 : len(l)-1]
}

// Substitution represents a variable and a value.
type Substitution struct {
	Var   string
	Value *Expression
}

func (s Substitution) String() string {
	return ast.Cons(&ast.SExpr{Atom: &ast.Atom{Var: &ast.Variable{Name: s.Var}}}, s.Value).String()
}
