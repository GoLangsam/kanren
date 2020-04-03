// go doc put to good use	 
				
-------------------------------------------------------------------------------
## go doc .   Bindings		
type Bindings struct {
	// Has unexported fields.
}
    Bindings represents bindings (or "substitutions"): any logic variable may be
    bound to some symbolic expression representing its current value.

func New() *Bindings
func (b *Bindings) Clone() *Bindings
func (b *Bindings) Drop(v V) (x X, wasBound bool)
func (b *Bindings) IsBound(v V) (isBound bool)
func (b *Bindings) Occurs(v V, x X) bool
func (b *Bindings) Reify(x X) *Bindings
func (b *Bindings) Subs(v V) (x X, hasSubs bool)
func (b *Bindings) Unify(x, y X) bool
func (b *Bindings) Walk(x X) X
				
-------------------------------------------------------------------------------
## go doc -u Bindings		
type Bindings struct {
	bindings
	count int // used to create anonymus variables during reify
}
    Bindings represents bindings (or "substitutions"): any logic variable may be
    bound to some symbolic expression representing its current value.

func New() *Bindings
func (b *Bindings) Clone() *Bindings
func (b *Bindings) Drop(v V) (x X, wasBound bool)
func (b *Bindings) IsBound(v V) (isBound bool)
func (b *Bindings) Occurs(v V, x X) bool
func (b *Bindings) Reify(x X) *Bindings
func (b *Bindings) Subs(v V) (x X, hasSubs bool)
func (b *Bindings) Unify(x, y X) bool
func (b *Bindings) Walk(x X) X
func (b *Bindings) bind(v V, x X) *bindings
func (b *Bindings) clone() *bindings
func (b *Bindings) exts(v V, x X) bool
func (b *Bindings) newV() X
func (b *Bindings) walkV(v V) X
func (b *Bindings) walkX(x X) X
				
-------------------------------------------------------------------------------
## go doc -all		
package bind // import "github.com/GoLangsam/kanren/internal/µ/bind"


VARIABLES

var Cons = sexpr.Cons
var NewSymbol = sexpr.NewSymbol

TYPES

type Bindings struct {
	// Has unexported fields.
}
    Bindings represents bindings (or "substitutions"): any logic variable may be
    bound to some symbolic expression representing its current value.

func New() *Bindings
    New creates fresh and empty Bindings and returns a pointer.

func (b *Bindings) Clone() *Bindings
    Clone provides a clone of b.

func (b *Bindings) Drop(v V) (x X, wasBound bool)
    Drop makes v unbound, reports whether v was bound, and returns the
    expression (if any) v was previously bound with.

func (b *Bindings) IsBound(v V) (isBound bool)
    IsBound reports whether v is bound or not

func (b *Bindings) Occurs(v V, x X) bool
    occurs reports whether v occurs in x.

func (b *Bindings) Reify(x X) *Bindings
    Reify ...

func (b *Bindings) Subs(v V) (x X, hasSubs bool)
    Subs return the expression to which v is bound, if any.

func (b *Bindings) Unify(x, y X) bool
    Unify returns either (ok = false) or the substitutions extended with zero or
    more associations, where cycles in substitutions can lead to (ok = false)

func (b *Bindings) Walk(x X) X
    Walk ...

type V = *sexpr.Variable

type X = *sexpr.Expression

