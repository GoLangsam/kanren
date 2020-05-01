[:toc:]

# kanren
Logic: Programming it is about. - Relational Programming :-)

---
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

---
## Logical Statement

A logical statement is built along **state**s from **variable**s, and **goal**s.

Create an initial _state_, then obtain some _variables_ (and resulting _state_) from it.

Construct a _goal_ consisting of _variable bindings_ and logical operations (`AND`, `OR`), or _predicates_.

Then evaluate the goal using the state resulting from making the variables.

Evaluating a goal returns all possible solutions to the statement in the form of a number of states containing suitable variable bindings.

---

Note: Most often, the `EmptyState` is used as initial _state_.

A _predicate_ is what -given some variable- evaluates to `#t` or `#f` (with/in a given state).

---
Any _value_ used in a _state_ must be **unifiable**.

Unifying two values produces zero or more possible states, where variables that may be contained in the values may be bound in various combinations.

---
## Basics

Basic constructs of relational programming are

- the _Goals_
  - `Yeah` aka `#s` (aka `Succeed`) and
  - `Fail` aka `#u`,
- fresh logic _Variables_ `,u` `,v` `,w`...
- operators
  - `==`, aka `Eq`,
  - `Fresh` and
  - `Cond`e.

The laws of operators `fresh`, `==` and `cond`e are postulated:

### The law of `==`:

`(== v w)` is the same as `(== w v)`

### The law of `Fresh`:

If `x` is fresh, then `(== v x)` succeeds and associates `x` with `v`.

### The law of `Cond`e:

To get more values from `cond`e, pretend that the successful line has failed, refreshing all variables that got an association from that line.

---

## µKanren

The entire system is comprised of a handful of facilities and functions for
- the implementation of variables,
- streams themselves,
- the interface for creating and combining streams, and
- four primitive goal constructors.

### Variables & eXpressions

#### `V` - Variable

A logical _Variable_ may be named or anonymous.

- Method `Fresh(name string) V` creates a *named Variable*
- Method `V() V` produces a new *anonymous Variable*

Note: Internally, `V()` calls `Fresh`  with a auto-generated (and previously unused) name-string.

Equality is determined by coincidence of their name-string.

- `u.Equal(v) bool` reports it.

#### `X` -eXpression

In our context. this is what a Variable may represent, what a Variable may be bound to.

- Some call it _Value_ - as in  "_Value_ of a _Variable_".
- Some call it _Term_ - as in "_Term_ as part of an _eXpression_".

#### Circular reference

##### X->V->X

Some _eXpression_ `x` may be very basic, e.g consist of just one single thing (Atom).
And such, in turn, may just happen to be a *Variable*.

- `x.IsVariable() bool` reports this,
- `x.AsVariable() (V, bool)` reveals the _Variable_ (if any)
- `v.Expr() X` expresses a *Variable* `v` as *eXpression*.

##### V->X->V

Thus, we have a circular self reference:

- any *Variable* may camouflage as an *eXpression*, and
- some *eXpression* may reveal to be nothing but a single *Variable*.

`V->X->V->X->V->X->V->X->V->X->V->X->V->X->V->X->V` ...

This (indirect) self-reference is almost magical - and very important.

### Bind & Substitute

In the literature we meet a lot of overloaded terminology here and we need to take great care in order to avoid confusion.

[Constraint programming](https://en.wikipedia.org/wiki/Constraint_programming/) calls it: _Assignment_ or _assignation_ (or _model_). 

#### bind V<->X

How to bind a Variable and/with/to some eXpression?

A _bind_ consists of:

- some _Variable_ being bound, and
- some _eXpression(Value/Term)_ related to it, given to it, assigned to it.

The _eXpression(Value/Term)_ substitutes the _Variable_, so to say.

Thus: _bind_ establishes a relation between `V` and `X`. 

Within the notation of symbolic expressions we may write individual members of such relation as pair: _(x . y)_ or _(y . 5)_ or _(x . (y . 5))_.

The _Binding_ of some _Variable_ may change over time.

Any _Variable_  may be (or may become) unbound.

No _Variable_  can ever be bound to more than one _eXpression_ at the same time.

Note: Some bound _eXpression_  may itself be a logic _Variable_: e.g. `(y . 5)(x . y)`

#### Substitute

Once there is a `bind`, we may use it on any other _eXpression_ (think: formula).

Every occurrence of the _Variable_ (being bound) shall be replaced/substituted by the _eXpression_ (the _Variable_ is bound to/with).

Thus, we may need to _Walk_ the _eXpression_, and replace any occurrence of the _Variable_.

### `bind.Ings`
```go
b.Bind( V, X )
b.Bound( V )     -> ( value X, isBound bool )
b.IsBound( V )   -> ( isBound bool )

b.Resolve( X )   -> X
b.Walk( X )      -> X

b.Occurs( V, X ) -> bool

b.Unify(x, y X ) -> bool 
```

#### Representations:

- triangular
- idempotent

### Logical State
Mainly, it's just a `bind.Ings`. (Which some call _Substitution_).

Thus, a _logical state_ is a collection of variable _bindings_.

```
Some say: A _Substitution_ `S` is a mapping/association between logic variables `V` and their values `T` (also called `terms`).

Note: Some `T`  may itself be a logic Variable: e.g. `(y . 5)(x . y)`

() is the _empty Substitution_
```



### Unify & Reify

#### Unify

`b.Unify(a, b)` attempts to unify `a` and `b` (with respect to _binding_ `b`) and
returns a (potentially extended) Substitution, or #f (fails).

#### Reify

Any _eXpression_ containing some _logic variable_ may be considered "vague" due to their "real" or "appropriate" or "right" content / meaning / value not being known (yet).

In general, "_Reify_" means to convert mentally into a thing; to materialize it. In our context,Reification is the process of turning some _eXpression_ into another  _eXpression_ that does not contain any _logic variable_ any more.

```go
S.Reify( X ) -> X
```

`Reify` takes an arbitrary _eXpression_, perhaps containing variables, and returns it reified.

## Logical Goals and Goal Constructors

Goals are used to specify logical statements.

A _Goal_ `g` is a function that maps a _substitution_ `s` to an ordered sequence of zero or more values.
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

## Kanren Implementations

### S-Expression based languages

#### Lisp
./.

#### Scheme

**[miniKanren](https://github.com/miniKanren/miniKanren)** - a canonical implementation.

Implements the language described in the paper:

William E. Byrd, Eric Holk, and Daniel P. Friedman.
miniKanren, Live and Untagged: Quine Generation via Relational Interpreters (Programming Pearl).
To appear in the Proceedings of the 2012 Workshop on Scheme and Functional Programming, Copenhagen, Denmark, 2012.

Core µKanren: `==`(aka `Unify`), `fresh`, `conde`
Extensions: `=/=`, `symbol`o, `number`o, `absent`o

**[µKanren](https://github.com/jasonhemann/microKanren/)** - a reference implementation.

**[Kanren-Book](https://sourceforge.net/projects/kanren/)**

##### The triology

[The little Schemer](https://github.com/pkrumins/the-little-schemer/)
[The seasoned Schemer](https://github.com/pkrumins/the-seasoned-schemer/)
[The reasoned Schemer](https://github.com/pkrumins/the-reasoned-schemer/)

Code from the books by Peteris Krumins (peter@catonmat.net).
His blog is at http://www.catonmat.net/  --  good coders code, great reuse.


### Racket

### Clojure

**[core.logic)](https://github.com/clojure/core.logic/)**
**[core.logic.wiki](https://github.com/clojure/core.logic.wiki/)**

**[logic-tutorial](https://github.com/swannodette/logic-tutorial)**

### Rust

**[rslogic](https://github.com/kulibali/rslogic/)** is a logic programming framework for Rust inspired by [µKanren](https://github.com/jasonhemann/microKanren/).

### Go

**[gologic](https://github.com/hiredman/gologic/)** includes examples

**[gominikanren](https://github.com/awalterschulze/gominikanren/)**

**[ukanren-go](https://github.com/elliotdavies/ukanren-go/)**

## Bibliograpy

**[Fairness](http://okmij.org/ftp/Computation/monads.html#fair-bt-stream)** mentions early Kanren.

* [miniKanren Live and Untagged: Quine Generation via Relational Interpreters](http://webyrd.net/quines/quines.pdf)
* [cKanren: miniKanren with Constraints](http://scheme2011.ucombinator.org/papers/Alvis2011.pdf)
* [μKanren: A Minimal Functional Core for Relational Programming](http://webyrd.net/scheme-2013/papers/HemannMuKanren2013.pdf)
* [Relational Programming in miniKanren: Techniques, Applications, and Implementations](https://scholarworks.iu.edu/dspace/bitstream/handle/2022/8777/Byrd_indiana_0093A_10344.pdf)
* [Pure, Declarative, and Constructive Arithmetic Relations (Declarative Pearl)](http://okmij.org/ftp/Prolog/Arithm/arithm.pdf)
* [From Variadic Functions to Variadic Relations - A miniKanren Perspective](http://scheme2006.cs.uchicago.edu/12-byrd.pdf)

* [core.logic](https://github.com/clojure/core.logic)

## Notes todo

Plenty of other miniKanren use log-time persistent maps for their substitutions; core.logic (https://github.com/clojure/core.logic) and veneer (https://github.com/tca/veneer) certainly do.

---
[JaCoP](http://jacop.osolpro.com) is a finite domain solver written in
pure Java that has been in continuous development since 2001.
In the yearly [MiniZinc](http://www.minizinc.org) constraint challenges it has received the Silver award in the fixed category for the past three years.

Some basic testing seems to show that JaCoP is anywhere from 10X-100X faster than core.logic at solving Finite Domain problems.
While there is a considerable amount of work to be done to improve the performance
of core.logic's general constraint framework, it's unlikely we'll achieve JaCoP finite domain solving performance in the near future.
Thus JaCoP integration is attractive.

---

## miniKanren operators

`eq` is the basic goal constructor: it succeeds if its arguments unify, fails otherwise.

`conde` accepts two or more clauses made of lists of goals, and returns the logical disjunction of these clauses.

`fresh` accepts a list of one or more logic variables, and a block containing
one or more goals. The logic variables are bound into the lexical scope of the
block, and the logical conjuction of the goals is performed.

## Interface operators

`run` is similar to `run_all`, but returns at most `n` results.

`run_all` accepts a list of one or more logic variables, an optional module name for a constraint solver, and a block containing one or more goals.
The logic variables are bound into the lexical scope of the block, and the logical conjunction of the goals is performed.
All results of the conjunction are evaluated, and are returned in terms of the first logic variable given.


## Impure operators

ExKanren/lib/minikanren.ex

### `conda`
`conda` accepts lists of goal clauses similar to `conde`, but returns the result of at most one clause: the first clause to have its first goal succeed.

```ex
run_all([out, x, y]) do
	conda do
		[eq(1, 2), eq(out, "First clause")]
		[appendo(x, y, [1, 2, 3]), eq(out, {x, "Second clause"})]
		[eq(x, y), eq(x, 1), eq(out, {{x, y}, "Third clause"})]
	end
end

[{[], "Second clause"}, {[1], "Second clause"}, {[1, 2], "Second clause"}, {[1, 2, 3], "Second clause"}]
```

### `condu`

`condu` is similar to `conda`, but takes only a single result from the first goal of the first successful clause.

### `project`

`project` binds the associated value (if any) of one or more logic variables  into lexical scope and allows them to be operated on.

```ex
run_all([out, x, y]) do
	eq(x, 3)
	eq(y, 7)
	project([x, y]) do
		eq(x + y, out)
	end
end
[10]
```

##   Some common relations used in miniKanren.

`succeed` is a goal that always succeeds.
`succeed` is a goal that ignores its argument and always succeeds.

`fail` is a goal that always fails.
`fail` is a goal that ignores its argument and always fails.

`heado` relates `h` as the head of list `ls`.
`tailo` relates `t` as the tail of list `ls`.
`conso` relates `h` and `t` as the head and tail of list `ls`.
(Equivalent to `eq([h | t], ls)`.)

`membero` relates `a` as being a member of the list `ls`.

`appendo` relates the list `ls` as the result of appending lists `l1` and `l2`.

`emptyo` relates `a` to the empty list.

## Non-relational functions

`onceo` accepts a goal function, and allows it to succeed at most once.

`copy_term` copies the term associated with `x` to `y`, replacing any fresh logic variables in `x` with distinct fresh variables in `y`.

`is` projects its argument `b`, and associates `a` with the result of the unary operation `f.(b)`

`fresho` succeeds if its argument is a fresh logic variable, fails otherwise.
