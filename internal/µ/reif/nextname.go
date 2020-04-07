package reif

import "strconv"

// nextName provides a new unique name
// useful to create a new fresh anonymous variable
// nextName is a helper for y.Reify.
func (y *Y) nextName() string {
	n := y.count
	y.count++
	return "_." + strconv.Itoa(n)
}
