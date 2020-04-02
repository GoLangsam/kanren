package redblack

// Element represents the constituent of any Node
type Element interface {
	Key() int
	Merge(Element) Element
}

type color bool

const (
	red   color = false
	black color = true
)

// Node is what the tree is made of.
type Node struct {
	color   color
	element Element
	left    *Node
	right   *Node
}

// Root returns (the root of) a fresh Red-Black Tree
func Root(e Element) *Node {
	return &Node{red, e, nil, nil}
}

func (color color) isRedL(n *Node) bool {
	if color == black && n != nil && n.color == red {
		return red == n.left.color
	}
	return false
}

func (color color) isRedR(n *Node) bool {
	if color == black && n != nil && n.color == red {
		return red == n.right.color
	}
	return false

}

func tree(a *Node, x Element, b *Node, y Element, c *Node, z Element, d *Node) *Node {
	return &Node{red, y, &Node{black, x, a, b}, &Node{black, z, c, d}}
}

func balance(x Element, color color, a *Node, y Element, b *Node) *Node {
	if color.isRedL(a) {

		z, d := y, b
		y, l, c := a.element, a.left, a.right
		x, a, b := l.element, l.left, l.right

		return tree(a, x, b, y, c, z, d)

	} else if color.isRedR(a) {

		z, d := y, b
		x, a, r := a.element, a.left, a.right
		y, b, c := r.element, r.left, r.right

		return tree(a, x, b, y, c, z, d)

	} else if color.isRedL(b) {

		z, l, d := b.element, b.left, b.right
		y, b, c := l.element, l.left, l.right

		return tree(a, x, b, y, c, z, d)

	} else if color.isRedR(b) {

		y, b, r := b.element, b.left, b.right
		z, c, d := r.element, r.left, r.right

		return tree(a, x, b, y, c, z, d)

	} else {

		return &Node{color, y, a, b}

	}
}

func (tree *Node) ins(x Element) (here *Node) {
	if tree == nil {
		here = Root(x)
	} else {
		color := tree.color
		y, a, b := tree.element, tree.left, tree.right
		x_key, y_key := x.Key(), y.Key()

		if y_key > x_key {
			here = balance(x, color, a.ins(x), y, b)
		} else if x_key == y_key {
			return &Node{color, y.Merge(x), a, b}
		} else {
			here = balance(x, color, a, y, b.ins(x))
		}
	}
	return
}

func make_black(s *Node) *Node {
	if s.color == red {
		return s
	} else {
		return &Node{black, s.element, s.left, s.right}
	}
}

// Insert an Element
func (tree *Node) Insert(e Element) *Node {
	return make_black(tree.ins(e))
}

// Locate an Element given it's Key()
func (tree *Node) Locate(n int) (Element, bool) {
	for {
		if tree == nil {
			return nil, false
		} else {
			x, l, r := tree.element, tree.left, tree.right
			if x.Key() == n {
				return x, true
			} else if x.Key() > n {
				tree = l
			} else {
				tree = r
			}
		}
	}
}
