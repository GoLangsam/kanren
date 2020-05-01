package bind

import "fmt"

func ExampleIngs_Reify() {
	b := New()

	four := NewInt(4)
	five := NewInt(5)
	x := b.Fresh("x")
	y := b.Fresh("y")
	z := b.Fresh("z")

	b.Bind(x, four)
	b.Bind(y, five)
	b.Bind(z, y)
	fmt.Println(b)

	fmt.Println("x =", b.Reify(x))
	fmt.Println("z =", b.Reify(z))

	l := NewList(x, y, z, b.Fresh("u"), Cons(x, y))
	fmt.Println(l)

	// Output:
	// ((,x . 4)(,y . 5)(,z . ,y))
	// x = 4
	// z = 5
	// (,x ,y ,z ,u (,x . ,y))
}
