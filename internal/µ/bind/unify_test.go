package bind

import "fmt"

import "github.com/GoLangsam/sexpr"

func ExampleIngs_Bind() {
	b := New()

	five := sexpr.NewInt(5)
	x := sexpr.NewVariable("x")
	// vx, _ := x.AsVariable()

	b.Bind(x, five)

	fmt.Println(b)

	// Output:
	// ((,x . 5))
}

func ExampleIngs_Unify_yes() {
	b := New()

	five := sexpr.NewInt(5)
	x := sexpr.NewVariable("x")
	y := sexpr.NewVariable("y")
	z := sexpr.NewVariable("z")

	b.Bind(x, five)
	b.Bind(y, five)

	fmt.Println(b.Unify(x, y))

	b.Bind(z, y)

	// ((,x . 5)(,y . 5)(,z . ,y))

	fmt.Println(b.Unify(x, z))

	b.Drop(y)

	fmt.Println(b.Unify(x, z))
	yX, ok := b.Subs(y)
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

	four := sexpr.NewInt(4)
	five := sexpr.NewInt(5)
	x := sexpr.NewVariable("x")
	y := sexpr.NewVariable("y")

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
