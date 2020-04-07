// go doc put to good use	 
				
-------------------------------------------------------------------------------
## go doc .  Ings		
type Ings struct {
	// Has unexported fields.
}
    Ings represents bindings (or "substitutions"): any logic variable may be
    bound to some symbolic expression representing its current value.

    Use as `bind.Ings` (pun intended).

    The zero value is not useful - initialize with `bind.New()`.

func New() *Ings
func (bind *Ings) Bind(v V, x X) *Ings
func (bind *Ings) Drop(v V) (x X, wasBound bool)
func (bind *Ings) IsBound(v V) (isBound bool)
func (bind *Ings) Occurs(v V, x X) bool
func (bind *Ings) Subs(v V) (x X, hasSubs bool)
func (bind *Ings) Unify(x, y X) bool
func (bind *Ings) Walk(x X) X
				
-------------------------------------------------------------------------------
## go doc -u Ings		
type Ings struct {
	bound map[V]X
}
    Ings represents bindings (or "substitutions"): any logic variable may be
    bound to some symbolic expression representing its current value.

    Use as `bind.Ings` (pun intended).

    The zero value is not useful - initialize with `bind.New()`.

func New() *Ings
func (bind *Ings) Bind(v V, x X) *Ings
func (bind *Ings) Drop(v V) (x X, wasBound bool)
func (bind *Ings) IsBound(v V) (isBound bool)
func (bind *Ings) Occurs(v V, x X) bool
func (bind *Ings) Subs(v V) (x X, hasSubs bool)
func (bind *Ings) Unify(x, y X) bool
func (bind *Ings) Walk(x X) X
func (bind *Ings) exts(v V, x X) bool
func (bind *Ings) walkV(v V) X
func (bind *Ings) walkX(x X) X
				
-------------------------------------------------------------------------------
## go doc -all		
package bind // import "github.com/GoLangsam/kanren/internal/µ/bind"


TYPES

type Ings struct {
	// Has unexported fields.
}
    Ings represents bindings (or "substitutions"): any logic variable may be
    bound to some symbolic expression representing its current value.

    Use as `bind.Ings` (pun intended).

    The zero value is not useful - initialize with `bind.New()`.

func New() *Ings
    New creates fresh and empty mapping of/for bind.Ings and returns a pointer.

func (bind *Ings) Bind(v V, x X) *Ings
    Bind binds x to v, so v is bound to x. Thus, (v . x) resembles a
    substitution pair. Note: Bind does not avoid circular bindings.

func (bind *Ings) Drop(v V) (x X, wasBound bool)
    Drop makes v unbound, reports whether v was bound, and returns the
    expression (if any) v was previously bound with.

func (bind *Ings) IsBound(v V) (isBound bool)
    IsBound reports whether v is bound or not

func (bind *Ings) Occurs(v V, x X) bool
    Occurs reports whether v occurs in x.

func (bind *Ings) Subs(v V) (x X, hasSubs bool)
    Subs returns the expression to which v is bound, if any.

    This expression shall substitute the variable, which shall thus become
    substituted by its value.

func (bind *Ings) Unify(x, y X) bool
    Unify extends the bind.Ings with zero or more associations in an attempt to
    see whether the given eXpressions are equal and reports its success.
    Circular bindings imply failure (ok = false).

func (bind *Ings) Walk(x X) X
    Walk ... some call it `walkstar` or `walk*`

type V = *sexpr.Variable
    V represents a logic variable

type X = *sexpr.Expression
    X represents a symbolic expression

