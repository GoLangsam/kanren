package bind

import "strconv"

// nextName provides a new unique name
// useful to create a new fresh anonymous variable
// nextName is a helper for b.Reify.
func (b *Ings) nextName() string {
	n := b.count // len(b.bound)
	b.count++
	return "_." + strconv.Itoa(n)
}
