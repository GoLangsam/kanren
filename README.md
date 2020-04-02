[:toc:]

# kanren
Logic: Programming it is about :-)

## some Nomenklatura
- **reify**:
    to convert mentally into a thing; to materialize.
    
- **unify**:
    to make, form into, or cause to become one;
    to combine (two or more) in one;
	to join (one or more) to or with another or others so as to form one whole or unit;
	to unite, consolidate. 
	
- **expression**
    The action of expressing or representing (a meaning, thought, state of things) in words or symbols; the utterance (of feelings, intentions, etc.). 
	The action or process of manifesting (qualities or feelings) by action, appearance or other evidences or tokens. 

## Logical Statement

A logical statement is built from **state**s, **variable**s, and **goal**s.

Create an initial _state_, then obtain some _variables_ (and resulting _state_) from it.

Construct a _goal_ consisting of _variable bindings_ and logical operations (`AND`, `OR`), or _predicates_.

Then evaluate the goal using the state resulting from making the variables.

Evaluating a goal returns all possible solutions to the statement in the form of a number of states containing variable bindings.

---

Note: Most often, the `EmptyState` is used as initial _state_.

A _predicate_ is what -given some variable- evaluates to `#t` or `#f` (with/in a given state).

---
Any _value_ used in a state must be **unifiable**.

Unifying two values produces zero or more possible states, where variables that may be contained in the values may be bound in various combinations.

---
## µKanren

The entire system is comprised of a handful of functions for
- the implementation of variables,
- streams themselves,
- the interface for creating and combining streams, and
- four primitive goal constructors.

### Variable, Substitution, Unification

#### `V` - Variables

A logical _Variable_ is represented by some variable index, an `int`.
Equality is determined by coincidence of their index numbers.

```scheme
(define (var c) (vector c))
(define (var? x) (vector? x))
(define (var=? x 1 x 2 ) (= (vector-ref x 1 0) (vector-ref x 2 0)))
```

```go
type index int
type V index
func (v V)Eq(x V) bool
func (*S)V() V // extend State and return a fresh Variable
func IsV(some interface{}) bool
```

#### `T` -Term

In our context. this is what a Variable may represent, what a Variable may be bound to. 
Sometimes it is called *Value* - as in  "*Value* of a *Variable*".

#### bind V <->T

How to bind a Variable and/with/to some Value/Term/Expression?

A _Binding_ consists of :

- the _Variable_ being bound, and

- some _Value/Term/Expression_ related to it, given to it, assigned to it.

The _Value/Term/Expression_ substitutes the _Variable_, so to say.
Every occurrence of the _Variable_ (being bound) shall be replaced/substituted by the _Value/Term/Expression_ (the _Variable_ is bound to/with.)

#### bond

A bond serves as a de-reference of the pair: `( _Variable_ . _Term_ )`.

Any _Variable_ may have one (and no more than one) bond.
(Some _Variable_ without bond is (currently) unbound.)

Such bond points to / is index to/of some _Term_ / _Value_ / _Expression_ to which the _Variable_ is currently considered to be bound to.

### Logical State - Substitution

A _logical state_ is a collection of variable _bindings_, sometimes called: Substitution.

A _Substitution_ `S` is a mapping/association between logic variables `V` and their values `T` (also called `terms`).
Note: Some `T`  may itself be a logic Variable: e.g. `(y . 5)(x . y)`

() is the _empty Substitution_

```go
S.LookUp( V ) -> ( t T, found bool )
S.Walk  ( V )
S.Unify ( a, b T ) -> ( S, bool )
```
    `Unify` unifies terms a and b with respect to Substitution S and
    returns a (potentially extended) Substitution, or #f (fails)

#### Substitution Representations:

- triangular
- idempotent

### Reification

Reification is the process of turning some _term_ into another  _term_ that does not contain any _logic variable_.

```go
S.Reify( T ) -> T
```

`reify` within a substitution takes an arbitrary value, perhaps containing variables, and returns its reified value.

## miniKanren

```scheme
==     unifies two Values
cond
exist

run
```

## Logical Goals and Goal Constructors

Goals are used to specify logical statements.

A _goal_ `g` is a function that maps a _substitution_ `s` to an ordered sequence of zero or more values.
(These values are almost always substitutions.)

The sequence of values may be infinite. Thus, it is not a list but as a special kind of stream.

```go
type Goal func(S) ValueStream:
```

Evaluate a _Goal_ to produce zero or more `State`s, or collections of variable bindings/Substitutions.

### ValueStream

`(mZero)` represents the empty stream of values.

`(Unit a)` represents the stream containing just a, if a is a value.

`(Choice a f)` represents a non-empty stream, where

- a is the first value in the stream, and where
- f is a function of zero arguments.
  Invoking the function f produces the remainder of the stream.

`type Cons struct {head, tail}`is an alternative name & implementation for `Choice`

### Goal Constructors

A _Goal_ does either

- fail, or
- succeed

Evaluating a `Fail` goal always results in zero states.

Evaluating a `UnifyVal` goal attempts to unify a variable and a value.

Evaluating a `UnifyVar` goal attempts to unify the variables.

A `Conjunction` goal evaluates its sub-goal `a` using a given state, then evaluates sub-goal `b` using the results.

Evaluating a `Disjunction` goal returns all the possible states of evaluating `a` and `b`.

Evaluating a `Predicate` goal returns the given state only if the function returns `true`.

## αKanren

αKanren extends core miniKanren with operators for **nominal** logic programming

### `fresh`

introduces _noms_ ("Names" / "Atoms")

### `hash` `#` - a _term_ constructor`

used to limit the scope of a _nom_ within a _term_ ...

### `tie` - a _term_ constructor

A term constructed using `tie`  ◃▹ is called _binder_. In the term created by the expression
`(◃▹ a t)`, all occurrences of the _nom_ `a` within _term_ `t` are considered bound.

We refer to the term `t` as the body of `(◃▹ a t)`, and to the _nom_ `a` as being in _binding position_.

The ◃▹ constructor does not create _noms_; rather, it **delimits the scope** of _noms_, already introduced using `fresh`.

# References

**[µKanren](https://github.com/jasonhemann/microKanren)** is a reference implementation in Scheme.

**[rslogic](https://github.com/kulibali/rslogic)** is a logic programming framework for Rust inspired by [µKanren](https://github.com/jasonhemann/microKanren).

**[Fairness](http://okmij.org/ftp/Computation/monads.html#fair-bt-stream)** mentions early Kanren.


