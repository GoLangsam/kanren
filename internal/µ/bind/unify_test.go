package bind

import "fmt"

func ExampleIngs_Bind() {
	b := New()

	five := NewInt(5)
	x := NewVariable("x")
	// vx, _ := x.AsVariable()

	b.Bind(x, five)

	fmt.Println(b)

	// Output:
	// ((,x . 5))
}

func ExampleIngs_Unify_yes() {
	b := New()

	five := NewInt(5)
	x := NewVariable("x")
	y := NewVariable("y")
	z := NewVariable("z")

	b.Bind(x, five)
	b.Bind(y, five)

	fmt.Println(b.Unify(x, y))

	b.Bind(z, y)

	// ((,x . 5)(,y . 5)(,z . ,y))

	fmt.Println(b.Unify(x, z))

	b.Drop(y)

	fmt.Println(b.Unify(x, z))
	yX, ok := b.Bound(y)
	fmt.Println("y came back:", yX, ok)

	b.Drop(y)
	b.Drop(x)
	fmt.Println(b)

	// Output:
	// true
	// true
	// true
	// y came back: 5 true
	// ((,z . ,y))
}

func ExampleIngs_Unify_false() {
	b := New()

	four := NewInt(4)
	five := NewInt(5)
	x := NewVariable("x")
	y := NewVariable("y")

	b.Bind(x, four)
	b.Bind(y, five)

	// ((,x . 4)(,y . 5))
	fmt.Println(b.Unify(x, y))

	b.Drop(y)
	fmt.Println(b)

	// Output:
	// false
	// ((,x . 4))
}
