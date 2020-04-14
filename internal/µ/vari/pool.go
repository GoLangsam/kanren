// package vari provides a tiny factory
// for named and anonymous variables.
//
// use as `vari.Able` (pun intended)
package vari

import "strconv"

// Able - a factory for variables: `vari.Able` (pun intended)
//
//   u := vars.Fresh("u")
//   v := vars.V()      // anonymous
//   r := vars.Var("u") // re-use or Fresh
//
// The zero value is useful - lazy init.
type Able struct {
	ables map[string]X // holds the variables seen so far
}

func Ables() *Able {
	return &Able{
		ables: make(map[string]X),
	}
}

/*
func (v V) String() string {
	return "<lvar " + v.Name + strconv.Itoa(int(v.index)) + ">"
}
*/

// ableInit provides lazy initialization
func (p *Able) ableInit() *Able {
	if p == nil {
		p = Ables()
	}
	return p
}

// V returns a fresh variable (with an autogenerated name)
func (vari *Able) V() V {
	vari = vari.ableInit()
	return vari.Fresh("~." + strconv.Itoa(len(vari.ables)))
}

// Fresh returns a fresh variable
// (and updates the internal pool).
func (vari *Able) Fresh(name string) V {
	vari = vari.ableInit()

	_, known := vari.ables[name]
	if known {
		panic("Fresh: duplicate name for variable")
	}

	x := newVar(name)
	vari.ables[name] = x

	return x
}

// Clone returns a copy full of Fresh variables (with same names).
func (p *Able) Clone() *Able {
	c := Ables()
	for name, _ := range p.ables {
		c.Fresh(name)
	}
	return c
}
