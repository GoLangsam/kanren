package redblack

import "container/list"

type ducer func(interface{}, Element) (interface{}, bool)

func (tree *Node) Fold(init interface{}, f ducer) (interface{}, bool) {
	c := make(chan Element)
	go func() {
		stack := list.New()
		stack.PushFront(tree)
		for {
			if stack.Len() > 0 {
				a := stack.Front().Value
				x, ok := a.(Node)
				if !ok {
					panic("oh no")
				}
				c <- x.element
				if x.left != nil {
					stack.PushFront(x.left)
				}
				if x.right != nil {
					stack.PushFront(x.right)
				}

			} else {
				break
			}
		}
	}()
	r := init
	for item := <-c; item != nil; item = <-c {
		a, cont := f(r, item)
		if !cont {
			close(c)
			return a, cont
		}
		r = a
	}
	return r, true
}
