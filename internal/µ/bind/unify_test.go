package bind

import "fmt"

func ExampleIngs_Bind() {
	b := New()

	four := NewInt(4)
	five := NewInt(5)
	x := b.Fresh("x")
	// vx, _ := x.AsVariable()

	b.Bind(x, four)
	fmt.Println(b)

	b.Bind(x, five)
	fmt.Println(b)

	// Output:
	// ((,x . 4))
	// ((,x . 5))
}

func ExampleIngs_Unify_yes() {
	b := New()

	five := NewInt(5)
	x := b.Fresh("x")
	y := b.Fresh("y")
	z := b.Fresh("z")

	b.Bind(x, five)
	b.Bind(y, five)

	fmt.Println(b.Unify(x, y))

	b.Bind(z, y)

	fmt.Println(b)
	fmt.Println(b.Unify(x, z))

	// Output:
	// true
	// ((,x . 5)(,y . 5)(,z . ,y))
	// true
}

func ExampleIngs_Unify_false() {
	b := New()

	four := NewInt(4)
	five := NewInt(5)
	x := b.Fresh("x")
	y := b.Fresh("y")

	b.Bind(x, four)
	b.Bind(y, five)

	fmt.Println(b)
	fmt.Println(b.Unify(x, y))

	// Output:
	// ((,x . 4)(,y . 5))
	// false
}
