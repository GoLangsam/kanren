package reif

import "fmt"

func ExampleY_Reify() {
	b := Ier()

	four := NewInt(4)
	five := NewInt(5)
	x := b.Fresh("x")
	y := b.Fresh("y")
	z := b.Fresh("z")

	b.Bind(x, four)
	b.Bind(y, five)
	b.Bind(z, y)
	fmt.Println(b.Ings)

	l := NewList(x, y, z, b.Fresh("u"), Cons(x, y))
	fmt.Println(l)
	b = b.Reify(l)
	fmt.Println(b.Ings)

	// Output:
	// true
}
