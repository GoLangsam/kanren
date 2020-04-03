package kanren

import "github.com/GoLangsam/kanren/internal/µ/bind"

type Constraints struct{}

// S represents a State: Variables, Bindings, Constraints
type S struct {
	pool // known variables
	*bind.Bindings
	Constraints
}

// Init provides an empty state
func Init() *S {
	return &S{
		newPool(),
		bind.New(),
		Constraints{},
	}
}
