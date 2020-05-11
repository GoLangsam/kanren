//https://github.com/swannodette/logic-tutorial#zebras
package main

import "context"
import "fmt"
import "strconv"
import "time"

import . "github.com/GoLangsam/kanren"
import "github.com/GoLangsam/sexpr"

var NewSymbol = sexpr.NewSymbol

var Unify = Equal
var And = Conjunction
var Or = Disjoint

// Bunch of all houses
func B(H1, H2, H3, H4, H5 V) X {
	return NewList(H1, H2, H3, H4, H5)
}

// Q represents a Query
type Q struct {
	X          // Query variable
	V func() V // Paginator reifies fresh free variables
}

func (q Q) righto(a, b X) Goal {
	v := q.V
	return Or(
		Unify(B(a, b, v(), v(), v()), q.X),
		Unify(B(v(), a, b, v(), v()), q.X),
		Unify(B(v(), v(), a, b, v()), q.X),
		Unify(B(v(), v(), v(), a, b), q.X))
}

func (q Q) nexto(a, b X) Goal {
	return Or(
		q.righto(a, b),
		q.righto(b, a))
}

func (q Q) first_o(a X) Goal {
	v := q.V
	return Unify(B(a, v(), v(), v(), v()), q.X)
}

func (q Q) middleo(a X) Goal {
	v := q.V
	return Unify(B(v(), v(), a, v(), v()), q.X)
}

func (q Q) membero(a X) Goal {
	v := q.V
	return Or(
		Unify(B(a, v(), v(), v(), v()), q.X),
		Unify(B(v(), a, v(), v(), v()), q.X),
		Unify(B(v(), v(), a, v(), v()), q.X),
		Unify(B(v(), v(), v(), a, v()), q.X),
		Unify(B(v(), v(), v(), v(), a), q.X))
}

func Zebra(s S, qv V, mode uint8) Goal {

	Nat := [5]string{"englishmen", "norwegian", "spaniard", "ukranian", "japanese"}
	Smo := [5]string{"kools", "lucky strikes", "parliaments", "oldgolds", "chesterfields"}
	Dri := [5]string{"orange juice", "tea", "milk", "coffee", "water"}
	Pet := [5]string{"dog", "fox", "zebra", "horse", "snails"}
	Col := [5]string{"red", "blue", "ivory", "green", "yello"}

	S5 := func(s [5]string) (S1, S2, S3, S4, S5 X) {
		return NewSymbol(s[0]), NewSymbol(s[1]), NewSymbol(s[2]), NewSymbol(s[3]), NewSymbol(s[4])
	}

	eng, nor, spa, rus, jap := S5(Nat)
	koo, l_s, par, old, che := S5(Smo)
	ora, tea, mil, cof, wat := S5(Dri)
	dog, fox, zeb, hor, sna := S5(Pet)
	red, blu, ivo, gre, yel := S5(Col)

	g := GOAL

	// House
	H := func(nat, smo, dri, pet, col X) X {

		if nat.IsVariable() {
			g = g.And(OneOf(nat, eng, nor, spa, rus, jap))
		}
		if smo.IsVariable() {
			g = g.And(OneOf(smo, koo, l_s, par, old, che))
		}
		if dri.IsVariable() {
			g = g.And(OneOf(dri, ora, tea, mil, cof, wat))
		}
		if pet.IsVariable() {
			g = g.And(OneOf(pet, dog, fox, zeb, hor, sna))
		}
		if col.IsVariable() {
			g = g.And(OneOf(col, red, blu, ivo, gre, yel))
		}

		return NewList(nat, smo, dri, pet, col)
	}

	v := s.V      // factory for reified fresh variables
	q := Q{qv, v} // Querulant

	goal := And(

		// https://en.wikipedia.org/wiki/Zebra_Puzzle

		q.membero(H(eng, v(), v(), v(), red)), // The Englishman lives in the red house.
		q.membero(H(spa, v(), v(), dog, v())), // The Spaniard owns the dog.
		q.membero(H(v(), v(), cof, v(), gre)), // Coffee is drunk in the green house.
		q.membero(H(rus, v(), tea, v(), v())), // The Ukrainian drinks tea.
		q.righto(
			H(v(), v(), v(), v(), ivo), // The green house is immediately to the right of the ivory house.
			H(v(), v(), v(), v(), gre)),
		q.membero(H(v(), old, v(), sna, v())), // The Old Gold smoker owns snails.
		q.membero(H(v(), koo, v(), v(), yel)), // Kools are smoked in the yellow house.
		q.middleo(H(v(), v(), mil, v(), v())), // Milk is drunk in the middle house.
		q.first_o(H(nor, v(), v(), v(), v())), // The Norwegian lives in the first house.

		q.nexto(
			H(v(), v(), v(), fox, v()), // The man who smokes Chesterfields lives in the house next to the man with the fox.
			H(v(), che, v(), v(), v())),

		q.nexto(
			H(v(), v(), v(), hor, v()), // Kools are smoked in the house next to the house where the horse is kept.
			H(v(), koo, v(), v(), v())),

		q.membero(H(v(), l_s, ora, v(), v())), // The Lucky Strike smoker drinks orange juice.
		q.membero(H(jap, par, v(), v(), v())), // The Japanese smokes Parliaments.

		q.nexto(
			H(nor, v(), v(), v(), v()), // The Norwegian lives next to the blue house.
			H(v(), v(), v(), v(), blu)),

		q.membero(H(v(), v(), wat, v(), v())), // Who drinks water?
		q.membero(H(v(), v(), v(), zeb, v())), // Who owns the zebra?

	)

	switch {
	case mode == 1:
		goal = goal.And(g)
	case mode == 2:
		goal = g.And(goal)
	}

	return goal
}

func TacTic(title string) (tac func(string), tic func(s string) string) {
	var t time.Time
	tick := func(s string) string { t = time.Now(); return s }
	tack := func(s string) {
		see(title, s,
			tab, "finished",
			tab, "Time:",
			tab, time.Since(t).String())
		tick(s)
	}
	return tack, tick
}

func zebra(ctx context.Context, mode uint8) {

	done := ctx.Done()

	defer fmt.Println("=========================================================================")
	tack, tick := TacTic("Zebra")
	defer tack(tick(strconv.Itoa(int(mode))))

	s := NewS()
	q := s.Fresh("q")
	g := Zebra(s, q, mode)
	c := g(s)
	tack("build")
	for head, ok := c.Head(); ok; head, ok = c.Head() {
		fmt.Println("=========================================================================")
		//t.Println(head)
		fmt.Println(head.Reify(q))
		select {
		case <-done:
			return
		default:
		}
	}
	/*
		(
			(norwegian kools water fox yello)
			(ukranian chesterfields tea horse blue)
			(englishmen oldgolds milk snails red)
			(spaniard lucky strikes orange juice dog ivory)
			(japanese parliaments coffee zebra green)
		)
	*/

	if false {
		r := c.Reify(q)
		for i := range r {
			fmt.Println(i)
		}
	}
}
