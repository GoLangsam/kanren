package bind

// TODO: (b *Ings) String() string
// see subs.go; we do not need FMapSs, but we shall import sexpr

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
