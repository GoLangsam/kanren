package µ

import rb "github.com/GoLangsam/kanren/internal/redblack"

type Element = rb.Element // Key() int; Merge(E)E
type Node = rb.Node       // RedBlackTree

func root(e Element) *Node {
	return rb.Root(e)
}

// T represents a term
type T interface{}

// ============================================================================

// subs_pair is an Element
type subs_pair struct {
	V // Variable
	T // Term (value)
}

func (p subs_pair) Key() int {
	if &p == nil {
		return -1
	}
	return int(p.V)
}

func (p subs_pair) Merge(e Element) Element {
	return p
}

// ============================================================================

// node is a Map
type node struct {
	tree *Node
	size int
}

func (n node) Val_at(v V) (T, bool) {
	if n.tree != nil {
		x, found := n.tree.Locate(int(v))
		if found {
			a, ok := x.(subs_pair)
			if ok {
				return a.T, true
			} else {
				panic("oh no")
			}
		}
	}
	return nil, false
}

func (n node) With(v V, t T) Map {
	return node{n.tree.Insert(subs_pair{v, t}), n.size + 1}
}

func (n node) Count() int {
	return n.size
}
