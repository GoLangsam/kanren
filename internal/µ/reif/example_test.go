package reif

import "fmt"

import "github.com/GoLangsam/sexpr"

func ExampleY_Reify() {
	b := Ier()

	four := sexpr.NewInt(4)
	five := sexpr.NewInt(5)
	x := b.Fresh("x")
	y := b.Fresh("y")
	z := b.Fresh("z")

	b.Bind(x, four)
	b.Bind(y, five)
	b.Bind(z, y)
	fmt.Println(b.Ings)

	l := sexpr.NewList(x, y, z, b.Fresh("u"), sexpr.Cons(x, y))
	fmt.Println(l)
	b = b.Reify(l)
	fmt.Println(b.Ings)

	// Output:
	// true
}
