// go doc put to good use	 
				
-------------------------------------------------------------------------------
## go doc .  Ings		
type Ings struct {
	smap.SMap // SMap
}
    Ings represents bindings (or "substitutions" or ""): any logic variable may
    be bound to some symbolic expression representing its current value.

    Use as `bind.Ings` (pun intended).

    The zero value is not useful - initialize with `bind.New()`.

func New() Ings
func (bind Ings) Bind(v V, x X)
func (bind Ings) Clone() Ings
func (bind Ings) Occurs(v V, x X) bool
func (bind Ings) Resolve(v X) X
func (bind Ings) SafeBind(v V, x X) bool
func (bind Ings) String() string
func (bind Ings) Unify(x, y X) bool
func (bind Ings) Walk(x X) X
				
-------------------------------------------------------------------------------
## go doc -u Ings		
type Ings struct {
	smap.SMap // SMap
}
    Ings represents bindings (or "substitutions" or ""): any logic variable may
    be bound to some symbolic expression representing its current value.

    Use as `bind.Ings` (pun intended).

    The zero value is not useful - initialize with `bind.New()`.

func New() Ings
func (bind Ings) Bind(v V, x X)
func (bind Ings) Clone() Ings
func (bind Ings) Occurs(v V, x X) bool
func (bind Ings) Resolve(v X) X
func (bind Ings) SafeBind(v V, x X) bool
func (bind Ings) String() string
func (bind Ings) Unify(x, y X) bool
func (bind Ings) Walk(x X) X
				
-------------------------------------------------------------------------------
## go doc -all		
package bind // import "github.com/GoLangsam/kanren/internal/µ/bind"


TYPES

type Ings struct {
	smap.SMap // SMap
}
    Ings represents bindings (or "substitutions" or ""): any logic variable may
    be bound to some symbolic expression representing its current value.

    Use as `bind.Ings` (pun intended).

    The zero value is not useful - initialize with `bind.New()`.

func New() Ings
    New creates fresh and empty mapping of/for bind.Ings and returns a pointer.

func (bind Ings) Bind(v V, x X)
    Bind binds x to v, so v is bound to x. Thus, (v . x) resembles a
    substitution pair.

    Bind is a noOp if v or x are nil or v is not a Variable.

    Note: Bind does not attempt to avoid circular bindings. Use Occurs to check
    beforehand.

func (bind Ings) Clone() Ings
    Clone returns a shallow copy.

func (bind Ings) Occurs(v V, x X) bool
    Occurs reports whether v occurs in x.

func (bind Ings) Resolve(v X) X
    Resolve the eXpression by chasing along the bindings recurring down to the
    first non-Variable eXpression or down to the first unbound eXpression

func (bind Ings) SafeBind(v V, x X) bool
    exts attempts to bind v with x and fails if such bind would introduce a
    circle due to v occurs in x.

func (bind Ings) String() string
    String returns a string of the symbolic map sorted by key.

func (bind Ings) Unify(x, y X) bool
    Unify extends the bind.Ings with zero or more associations in an attempt to
    see whether the given eXpressions are equal and reports its success.
    Circular bindings imply failure (ok = false).

func (bind Ings) Walk(x X) X
    Walk ... some call it `walkstar` or `walk*`

type SMap interface {
	Clone() smap.SMap
	String() string

	Load(V) (X, bool)
	Store(V, X)

	Delete(X) // only for tests - TODO: remove
}
    SMap documents the behaviour required from a substitution map.

type V = X // *sexpr.Variable
    V is an eXpression which represents a logic variable

type X = *sexpr.Expression
    X represents a symbolic expression

