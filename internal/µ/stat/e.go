package stat

import "github.com/GoLangsam/kanren/internal/µ/reif"

type Constraints struct{}

// E represents a State: Variables, Bindings, Constraints
//
// use as `stat.E` (pun intended)
//
// Note: the zero value is not useful.
// Init provides an empty state.
type E struct {
	reif.Y
	Constraints
}

// Init provides an empty state.
func Init() *E {
	return &E{
		reif.Ier(),
		Constraints{},
	}
}

func (s *E) Clone() *E {
	return &E{
		s.Y.Clone(),
		s.Constraints,
	}
}
