// go doc put to good use	 
				
-------------------------------------------------------------------------------
## go doc . Bind		
type Bind struct {
	// Has unexported fields.
}
    Bind represents bindings: any logic variable may be bound to some symbolic
    expression representing the current value of such variable.

func New() *Bind
func (b *Bind) Bind(v V, x X)
func (b *Bind) Drop(v V)
func (b *Bind) IsBound(v V) (isBound bool)
func (b *Bind) Subs(v V) (x X, hasSubs bool)
func (b *Bind) Unify(x, y X) bool
func (b *Bind) Walk(x X) X
				
-------------------------------------------------------------------------------
## go doc -u Bind		
type Bind struct {
	bound map[V]bond // Variables bound via bond to one of the bonds
	bonds map[bond]X // bonds map (via bond) to the expression representing the current value.
	count int
}
    Bind represents bindings: any logic variable may be bound to some symbolic
    expression representing the current value of such variable.

func New() *Bind
func (b *Bind) Bind(v V, x X)
func (b *Bind) Drop(v V)
func (b *Bind) IsBound(v V) (isBound bool)
func (b *Bind) Subs(v V) (x X, hasSubs bool)
func (b *Bind) Unify(x, y X) bool
func (b *Bind) Walk(x X) X
func (b *Bind) exts(v V, x X) bool
func (b *Bind) occurs(v V, x X) bool
func (b *Bind) vAsX(v V) (x X)
func (b *Bind) walk(v V) X
				
-------------------------------------------------------------------------------
## go doc -all		
package bind // import "github.com/GoLangsam/kanren/internal/µ/bind"


VARIABLES

var Cons = sexpr.Cons

TYPES

type Bind struct {
	// Has unexported fields.
}
    Bind represents bindings: any logic variable may be bound to some symbolic
    expression representing the current value of such variable.

func New() *Bind
func (b *Bind) Bind(v V, x X)
    Bind binds x to v, so v is bound to x. Thus, (v . x) resembles a
    substitution pair. Note: Bind does not avoid circular bindings

func (b *Bind) Drop(v V)
    Drop makes v unbound. A existing Bind (if any) is discarded.

func (b *Bind) IsBound(v V) (isBound bool)
    IsBound reports whether v is bound or not

func (b *Bind) Subs(v V) (x X, hasSubs bool)
    Subs return the expression to which v is bound, if any.

func (b *Bind) Unify(x, y X) bool
    Unify returns either (ok = false) or the substitutions extended with zero or
    more associations, where cycles in substitutions can lead to (ok = false)

func (b *Bind) Walk(x X) X
    Walk ...

type V = *sexpr.Variable

type X = *sexpr.Expression

