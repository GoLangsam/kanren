package kanren

type Constraints struct{}

// S represents a State: Variables, Bindings, Constraints
type S struct {
	pool // known variables
	*Bindings
	Constraints
}

// Init provides an empty state
func Init() *S {
	return &S{
		newPool(),
		newBind(),
		Constraints{},
	}
}
