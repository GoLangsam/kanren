package pipe

import "github.com/GoLangsam/kanren/internal/µ/bind"
import "github.com/GoLangsam/sexpr"

type S = bind.Ings

var (
	NewS = bind.New
)

var _ = sexpr.NewVariable("q")

type aS interface {
	String() string
	Unify(x, y bind.X) bool
	V() bind.X
	Walk(v bind.V) bind.X
}
