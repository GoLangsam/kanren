[:toc:]

# kanren
Logic: Programming it is about :-)

## Logical Statement

A logical statement is built from **variable**s, **state**s, and **goal**s.

Create an initial _state_, then obtain some _variables_ (and resulting _state_) from it.

Construct a _goal_ consisting of _variable bindings_, logical operations (`AND`, `OR`), or _predicates_.

Then evaluate the goal using the state resulting from making the variables.

Evaluating a goal returns all possible solutions to the statement in the form of a number of states containing variable bindings.

----

Note: Most often, the `EmptyState` is used as initial _state_.

A _predicate_ is what -given some variable- evaluates to `#t` or `#f` (with/in a given state).

## miniKanren

```scheme
==     unifies two Values
cond
exist

run
```

## Variable, Substitution, Unification

A _logic Variable_ x V ...

### How to Bind a Variable and/with/to some Value/Term

A _Binding_ is a pair, and consists of :

- the _Variable_ being bound, and

- some _Value_ related to it, given to it, assigned to it.

  The _Value_ substitutes the _Variable_, so to say. Every occurrence of the _Variable_ (being bound) shall be replaced/substituted by the _Value_ (the _Variable_ is bound to/with.)  

### Logical State - Substitution

A _logical state_ is a collection of variable _bindings_, sometimes called: Substitution.

A _Substitution_ `S` is a mapping/association between logic variables V and values T (also called `terms`).
Note: The `rhs`  may itself be a logic Variable. E.g. `(y . 5)(x . y)`

() is the _empty Substitution_

S.LookUp( V )
S.Walk  ( V )

S.Unify ( a, b T ) -> ( S, bool )
    unifies two terms with respect to Substitution S,
    returns a (potentially extended) Substitution, or #f (fails)

#### Substitution Representations:

- triangular
- idempotent

### Reification

Reification is the process of turning a miniKanren _term_ into a Scheme value that does not contain any _logic variable_.

S.Reify( V ) -> V

`reify` takes a substitution s and an arbitrary value v, perhaps containing variables, and returns the reified value of v.

## Logical Goals and Goal Constructors

Goals are used to specify logical statements.

A _goal_ `g` is a function that maps a _substitution_ `s` to an ordered sequence of zero or more values.
(These values are almost always substitutions.)

The sequence of values may be infinite. Thus, it is not a list but as a special kind of stream.

type Goal func(Substitution) ValueStream:

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


