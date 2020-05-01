// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pipe

// ===========================================================================
// Beg of FanIn2 simple binary Fan-In

// FanIn2 returns a channel to receive
// all from both `from` and `inp2`
// before close.
func (from StreamOfStates) FanIn2(inp2 StreamOfStates) (out StreamOfStates) {
	cha := from.New()
	go cha.fanIn2(from, inp2)
	return cha
}

/* not used - kept for reference only.
// fanIn2 as seen in Go Concurrency Patterns
func fanIn2(out chan<- anyThing, inp, inp2 <-chan anyThing) {
	for {
		select {
		case e := <-inp:
			out <- e
		case e := <-inp2:
			out <- e
		}
	}
} */

func (into StreamOfStates) fanIn2(inp, inp2 StreamOfStates) {

	defer func() {
		inp.Drop()
		inp2.Drop()
		into.Close()
	}()

	_, in1 := inp.From()
	_, in2 := inp2.From()

	var (
		closed bool // we found a chan closed
		ok     bool // did we read successfully?
		e      S    // what we've read
	)

	done, done1, done2 := into.Done(), inp.Done(), inp2.Done()

	for !closed {
		select {
		case e, ok = <-in1:
			if ok {
				into.Provide(e)
			} else {
				in1 = in2     // swap inp2 into inp
				closed = true // break out of the loop
			}
		case e, ok = <-in2:
			if ok {
				into.Provide(e)
			} else {
				closed = true // break out of the loop				}
			}
		case <-done:
			return // abort
		case <-done1:
			return // abort
		case <-done2:
			return // abort
		}
	}

	// inp might not be closed yet. Drain it.
	for e := range in1 {
		into.Provide(e)
	}
}

// End of FanIn2 simple binary Fan-In
// ===========================================================================
