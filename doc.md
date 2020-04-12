// go doc put to good use	 
				
-------------------------------------------------------------------------------
## go doc .  Goal		
type Goal = µ.Goal

func Always() Goal
func Append(l, t, out X) Goal
func CallFresh(f func(X) Goal) Goal
func Car(list, head X) Goal
func Conjunction(gs ...Goal) Goal
func Cons(car, cdr, pair X) Goal
func Disjoint(gs ...Goal) Goal
func Equal(x, y X) Goal
func Failure() Goal
func IfThenElse(IF, THEN, ELSE Goal) Goal
func Never() Goal
func Null(x X) Goal
func Once(g Goal) Goal
func Success() Goal
				
-------------------------------------------------------------------------------
## go doc -all		
package kanren // import "github.com/GoLangsam/kanren"

Package kanren implements relational symbolic logic

VARIABLES

var (
	Unit       = µ.Unit
	Zero       = µ.Zero
	EmptyState = µ.EmptyState
)
var Fail = Failure
    Fail is an alias for Failure.


TYPES

type Goal = µ.Goal

func Always() Goal
    Always is a goal that always returns a never ending stream of success.

func Append(l, t, out X) Goal
    Append is the relation: append(l, t) == out.

func CallFresh(f func(X) Goal) Goal
    CallFresh expects a function f that returns a Goal given an eXpression.

    CallFresh returns the Goal which, when evaluated, applies f to a fresh
    anonymous variable and evaluates the resulting Goal.

    CallFresh allows to introduce host-language-symbols as free variables when
    constructing some Goal, e.g. in order to model some relation. See Append,
    for example.

func Car(list, head X) Goal
    Car is the relation: Car(list) == head.

func Conjunction(gs ...Goal) Goal
    Conjunction is a goal that returns a logical AND of the input goals.

func Cons(car, cdr, pair X) Goal
    Cons is the relation: Cons(car, cdr) == pair.

func Disjoint(gs ...Goal) Goal
    Disjoint is a goal that returns a logical OR of the input goals.

func Equal(x, y X) Goal
    Equal is a relation: it reports whether x unifies with y.

    Note: In Scheme, Equal is often spelled "==".

func Failure() Goal
    Failure is a goal that always returns an empty stream of states.

func IfThenElse(IF, THEN, ELSE Goal) Goal
    IfThenElse is a goal that evaluates the THEN goal if the IF goal is
    successful, otherwise it evaluates the ELSE goal.

func Never() Goal
    Never is a Goal that returns a never ending evaluation of itself.

    Note: This is a joke. Use on Your own risk!

func Null(x X) Goal
    Null is the relation: x == nil.

func Once(g Goal) Goal
    Once is a goal that returns the first success of g, if any, and discards
    further results, if any.

func Success() Goal
    Success is a goal that always returns the input state in the resulting
    stream of states.

type S = µ.S

type StreamOfStates = µ.StreamOfStates

type V = µ.V

type X = µ.X

