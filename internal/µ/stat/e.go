package stat

import "github.com/GoLangsam/kanren/internal/µ/bind"
import "github.com/GoLangsam/kanren/internal/µ/vari"

type Constraints struct{}

// TODO:
// vari.Able pool implements Fresh & V();
// bind.Ings implements newName and uses x := NewVariable, thus producing a X directly

// E represents a State: Variables, Bindings, Constraints
//
// use as `stat.E` (pun intended)
type E struct {
	*vari.Able // known variables
	*bind.Ings
	Constraints
}

// Init provides an empty state
func Init() *E {
	return &E{
		vari.Ables(),
		bind.New(),
		Constraints{},
	}
}

func (s *E) Clone() *E {
	return &E{
		s.Able.Clone(),
		s.Ings.Clone(),
		s.Constraints,
	}
}
