package µ

// Map represents a substitution map for variable V and term T
type Map interface {
	Val_at(V) (T, bool)
	With(V, T) Map
	Count() int
}

// emptyMap is a Map
type emptyMap struct{}

func (s emptyMap) Val_at(v V) (T, bool) {
	return nil, false
}

func (s emptyMap) With(v V, t T) Map {
	//return &SubsT{v,t,nil}
	return node{root(subs_pair{v, t}), 1}
}

func (s emptyMap) Count() int {
	return 0
}

func New() Map {
	return emptyMap{}
}
