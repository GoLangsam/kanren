package µ

import "github.com/GoLangsam/kanren/internal/µ/bind"
import "github.com/GoLangsam/kanren/internal/µ/vari"

type Constraints struct{}

// S represents a State: Variables, Bindings, Constraints
type S struct {
	vari.Able // known variables
	*bind.Bindings
	Constraints
}

// Init provides an empty state
func Init() *S {
	return &S{
		vari.Ables(),
		bind.New(),
		Constraints{},
	}
}
