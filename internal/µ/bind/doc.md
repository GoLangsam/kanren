// go doc put to good use	 
				
-------------------------------------------------------------------------------
## go doc .  Ings		
type Ings struct {
	// Has unexported fields.
}
    Ings represents bindings (or "substitutions"): any logic variable may be
    bound to some symbolic expression representing its current value.

    use as "bind.Ings" (pun intended)

func New() *Ings
func (b *Ings) Bind(v V, x X) *bindings
func (b *Ings) Clone() *Ings
func (b *Ings) Drop(v V) (x X, wasBound bool)
func (b *Ings) IsBound(v V) (isBound bool)
func (b *Ings) Occurs(v V, x X) bool
func (b *Ings) Reify(x X) *Ings
func (b *Ings) Subs(v V) (x X, hasSubs bool)
func (b *Ings) Unify(x, y X) bool
func (b *Ings) Walk(x X) X
				
-------------------------------------------------------------------------------
## go doc -u Ings		
type Ings struct {
	bindings
	count int // used to create anonymus variables during reify
}
    Ings represents bindings (or "substitutions"): any logic variable may be
    bound to some symbolic expression representing its current value.

    use as "bind.Ings" (pun intended)

func New() *Ings
func (b *Ings) Bind(v V, x X) *bindings
func (b *Ings) Clone() *Ings
func (b *Ings) Drop(v V) (x X, wasBound bool)
func (b *Ings) IsBound(v V) (isBound bool)
func (b *Ings) Occurs(v V, x X) bool
func (b *Ings) Reify(x X) *Ings
func (b *Ings) Subs(v V) (x X, hasSubs bool)
func (b *Ings) Unify(x, y X) bool
func (b *Ings) Walk(x X) X
func (b *Ings) clone() *bindings
func (b *Ings) exts(v V, x X) bool
func (b *Ings) nextName() string
func (b *Ings) walkV(v V) X
func (b *Ings) walkX(x X) X
				
-------------------------------------------------------------------------------
## go doc -all		
package bind // import "github.com/GoLangsam/kanren/internal/µ/bind"


TYPES

type Ings struct {
	// Has unexported fields.
}
    Ings represents bindings (or "substitutions"): any logic variable may be
    bound to some symbolic expression representing its current value.

    use as "bind.Ings" (pun intended)

func New() *Ings
    New creates fresh and empty mapping of/for bind.Ings and returns a pointer.

func (b *Ings) Bind(v V, x X) *bindings
    Bind binds x to v, so v is bound to x. Thus, (v . x) resembles a
    substitution pair. Note: Bind does not avoid circular bindings.

func (b *Ings) Clone() *Ings
    Clone provides a clone of b.

func (b *Ings) Drop(v V) (x X, wasBound bool)
    Drop makes v unbound, reports whether v was bound, and returns the
    expression (if any) v was previously bound with.

func (b *Ings) IsBound(v V) (isBound bool)
    IsBound reports whether v is bound or not

func (b *Ings) Occurs(v V, x X) bool
    Occurs reports whether v occurs in x.

func (b *Ings) Reify(x X) *Ings
    Reify ...

func (b *Ings) Subs(v V) (x X, hasSubs bool)
    Subs returns the expression to which v is bound, if any.

    This expression shall substitute the variable, which shall thus become
    substituted by its value.

func (b *Ings) Unify(x, y X) bool
    Unify returns either (ok = false) or the substitutions extended with zero or
    more associations, where cycles in substitutions can lead to (ok = false)

func (b *Ings) Walk(x X) X
    Walk ... some call it `walkstar` or `walk*`

type V = *sexpr.Variable
    V represents a logic variable

type X = *sexpr.Expression
    X represents a symbolic expression

