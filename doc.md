// go doc put to good use	 
				
-------------------------------------------------------------------------------
## go doc .  Goal		
var (
	FAIL Goal = µ.Failure() // FAIL represents Failure.
	GOAL Goal = µ.Success() // GOAL represents Success.

	NewS                = µ.NewS // only used in test programs
	Unit                = µ.Unit
	ZERO StreamOfStates = µ.ZERO // used by Equal-relation

)
type Goal = µ.Goal // func(S) StreamOfStates

func Append(aHead, aTail, aList X) Goal
func AppendAWS(l, t, out X) Goal
func CallFresh(f func(V) Goal) Goal
func Car(list, head X) Goal
func Cdr(list, tail X) Goal
func Conjunction(gs ...Goal) Goal
func Cons(car, cdr, pair X) Goal
func Disjoint(gs ...Goal) Goal
func EitherOr(THIS, THAT Goal) Goal
func Equal(x, y X) Goal
func Fresh1(f func(V) Goal) Goal
func Fresh2(f func(V, V) Goal) Goal
func Fresh3(f func(V, V, V) Goal) Goal
func Fresh4(f func(V, V, V, V) Goal) Goal
func Fresh5(f func(V, V, V, V, V) Goal) Goal
func Fresh6(f func(V, V, V, V, V, V) Goal) Goal
func Fresh7(f func(V, V, V, V, V, V, V) Goal) Goal
func Fresh8(f func(V, V, V, V, V, V, V, V) Goal) Goal
func IfThenElse(IF, THEN, ELSE Goal) Goal
func Null(x X) Goal
func Once(g Goal) Goal
				
-------------------------------------------------------------------------------
## go doc -all		
package kanren // import "github.com/GoLangsam/kanren"

Package kanren implements relational symbolic logic

VARIABLES

var (
	FAIL Goal = µ.Failure() // FAIL represents Failure.
	GOAL Goal = µ.Success() // GOAL represents Success.

	NewS                = µ.NewS // only used in test programs
	Unit                = µ.Unit
	ZERO StreamOfStates = µ.ZERO // used by Equal-relation

)

TYPES

type Goal = µ.Goal // func(S) StreamOfStates

func Append(aHead, aTail, aList X) Goal
    Append is the relation: append(aHead, aTail) == aList.

func AppendAWS(l, t, out X) Goal
    AppendAWS is the relation: append(l, t) == out.

func CallFresh(f func(V) Goal) Goal
    CallFresh expects a function f that returns a Goal given an eXpression.

    CallFresh returns the Goal which, when evaluated, applies f to a fresh
    anonymous variable and evaluates the resulting Goal.

    CallFresh allows to introduce a host-language-symbol as a free variable when
    constructing some Goal, e.g. in order to model some relation. See `Append`,
    for example.

func Car(list, head X) Goal
    Car is the relation: Car(list) == head.

func Cdr(list, tail X) Goal
    Cdr is the relation: Cdr(list) == tail.

func Conjunction(gs ...Goal) Goal
    Conjunction is a goal that returns a logical AND of the input goals.

func Cons(car, cdr, pair X) Goal
    Cons is the relation: Cons(car, cdr) == pair.

func Disjoint(gs ...Goal) Goal
    Disjoint is a goal that returns a logical OR of the input goals.

func EitherOr(THIS, THAT Goal) Goal
    EitherOr is a goal that behaves like the THIS Goal unless THIS fails, when
    it behaves like the THAT Goal.

func Equal(x, y X) Goal
    Equal is a relation: it reports whether x unifies with y.

    Note: In Scheme, Equal is often spelled "==".

func Fresh1(f func(V) Goal) Goal
func Fresh2(f func(V, V) Goal) Goal
func Fresh3(f func(V, V, V) Goal) Goal
func Fresh4(f func(V, V, V, V) Goal) Goal
func Fresh5(f func(V, V, V, V, V) Goal) Goal
func Fresh6(f func(V, V, V, V, V, V) Goal) Goal
func Fresh7(f func(V, V, V, V, V, V, V) Goal) Goal
func Fresh8(f func(V, V, V, V, V, V, V, V) Goal) Goal
func IfThenElse(IF, THEN, ELSE Goal) Goal
    IfThenElse is a goal that upon evaluation probes the IF goal and, using a
    clone of the state, evaluates the THEN goal, if IF evaluates successful and
    evaluates the ELSE goal otherwise.

func Null(x X) Goal
    Null is the relation: x == nil.

func Once(g Goal) Goal
    Once is a goal that returns the first success of g, if any, and discards
    further results, if any.

type S = µ.S

type StreamOfStates = µ.StreamOfStates

type V = µ.V

type X = µ.X

