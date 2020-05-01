package pipe

import "github.com/GoLangsam/kanren/internal/µ/bind"
import "github.com/GoLangsam/kanren/internal/µ/pipe/xxs"

type S = bind.Ings
type StreamOfStates struct {
	pipe.StreamOfStates
}

var (
	NewS = bind.New
)

func New(done <-chan struct{}) StreamOfStates {
	return StreamOfStates{pipe.StreamOfStatesMakeChan(done)}
}

func (s StreamOfStates) New() StreamOfStates {
	return StreamOfStates{s.StreamOfStates.New()}
}
