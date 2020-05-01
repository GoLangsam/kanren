// package abort provides an aborter.
package abort

// Aborter provides a simple mechanism to broadcast
// a request to cancel some ongoing work.
//
// Aborter is intended to be embedded in some struct,
// and to be initialised from some "context.Done()"
// or via WithCancel().
type Aborter <-chan struct{}

// Done returns a channel that's closed when work done on behalf of this
// context should be canceled. Done may return nil if this context can
// never be canceled. Successive calls to Done return the same value.
//
// Done is provided for use in select statements,
// and is intentionally similar to "context.Done()".
//
// See https://blog.golang.org/pipelines
// or https://github.com/GoLangsam/pipe
// for more examples of
// how to use a Done channel for cancelation.
func (sig Aborter) Done() (done <-chan struct{}) {
	return sig
}

// Ever reports whether to continue, or not.
//
// Ever is intended for for-loops:
//     `for foo.Ever() { ... }`
// (pun intended)
//
// Note: Due to the non-deterministic nature
// of the underlying mechanisms,
// sometimes false positives may be returned.
func (sig Aborter) Ever() (noNeedToStop bool) {
	select {
	case _, noNeedToStop = <-sig:
	default: // do not block
		noNeedToStop = true
	}
	return
}

// A CancelFunc tells an operation to abandon its work.
// A CancelFunc does not wait for the work to stop.
// After the first call, subsequent calls to a CancelFunc do nothing.
type CancelFunc func()

// WithCancel returns a new aborter and its cancel function.
func WithCancel() (Aborter, CancelFunc) {
	abort := make(chan struct{})
	cancel := func() { close(abort) }
	return abort, cancel
}
