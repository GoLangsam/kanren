#

## Infix Notation

( ,x == 5 )
( ,N == Tom )

( G && H ) || ( G1 && G2 )

( G * H + G1 * G2 )

## Goals

Let G, H, I ... be Variables of Type `Goal`

### Binary Goal-Rel

`<=>` - `==` - `=` : Equivalence <=> G(s).Unify(H(s))

### Binary Goal-Ops

#### OR - || - *
#### AND  - && - + :

#### Operator Precedence
G & H | I == G & (H|I)  

### Goal algebra

#### Assoziativ
`G || ( H || I )` <=> `( G ||  H ) || I`	: oor is associativ 
`G && ( H && I )` <=> `( G &&  H ) && I`	: and is associativ  

#### Kommutativ
`G || H` <=> `H || G`			: oor is kommutativ
`G && H` <=> `H && G`			: and is kommutativ

#### Distributiv
`( G && H1 ) || ( G && H2 )`
<=> `( G || ( H1 && H2 ) )`		: distributiv

#### ZERO
`ZERO && G` <=> `ZERO`			: See? It's logic! For numbers, it would be G !!!
`ZERO || G` <=> `G`

#### UNIT
`UNIT && G` <=> `G

----
callRun(g) == g.Eval() -> StreamOfStates {return g(EmptyState)}
```go
func callFresh(f func (X) Goal) Goal {
	return func (s S) StreamOfStates {
		// s = s.Clone() // TODO: need??
		x := s.Fresh()	
		return f(x)(s)
	}

}
```

----
## StreamOfStates as a Monad

Unit(s) is the 'Return(s)'-function of the StreamOfStates-Monad

UNIT is the 'multiplicative neutral element' of the Goal-Algebra.