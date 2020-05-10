// go doc put to good use	 
				
-------------------------------------------------------------------------------
## go doc .  StreamOfStates	
package pipe // import "github.com/GoLangsam/kanren/internal/µ/pipe/xxs"

type StreamOfStates struct {
	abort.Aborter
	// Has unexported fields.
}
    StreamOfStates is a supply channel

func StreamOfStatesMakeBuff(done <-chan struct{}, cap int) StreamOfStates
func StreamOfStatesMakeChan(done <-chan struct{}) StreamOfStates
func (c StreamOfStates) Cap() int
func (into StreamOfStates) Close()
func (from StreamOfStates) Drop()
func (from StreamOfStates) FanIn2(inp2 StreamOfStates) (out StreamOfStates)
func (from StreamOfStates) From() (req chan<- struct{}, rcv <-chan S)
func (from StreamOfStates) Get() (val S, open bool)
func (into StreamOfStates) Into() (req <-chan struct{}, snd chan<- S)
func (c StreamOfStates) Len() int
func (c StreamOfStates) New() StreamOfStates
func (into StreamOfStates) Next() (ok bool)
func (into StreamOfStates) NextGetFrom(from StreamOfStates) (val S, ok bool)
func (into StreamOfStates) Provide(val S) (ok bool)
func (into StreamOfStates) Put(val S) (ok bool)
func (from StreamOfStates) Receive() (val S, open bool)
func (c StreamOfStates) Self() StreamOfStates
func (into StreamOfStates) Send(val S) (sent bool)
				
-------------------------------------------------------------------------------
## go doc -all		

package pipe // import "github.com/GoLangsam/kanren/internal/µ/pipe/xxs"


VARIABLES

var (
	NewS = bind.New
)

TYPES

type S = bind.Ings

type StreamOfStates struct {
	abort.Aborter
	// Has unexported fields.
}
    StreamOfStates is a supply channel

func StreamOfStatesMakeBuff(done <-chan struct{}, cap int) StreamOfStates
    StreamOfStatesMakeBuff returns a (pointer to a) fresh buffered supply
    channel (with capacity=`cap`).

func StreamOfStatesMakeChan(done <-chan struct{}) StreamOfStates
    StreamOfStatesMakeChan returns a (pointer to a) fresh unbuffered supply
    channel.

func (c StreamOfStates) Cap() int
    Cap reports the capacity of the underlying value channel.

func (into StreamOfStates) Close()
    Close is to be called by a producer when finished sending. The value channel
    is closed in order to broadcast this.

    In order to avoid deadlock, pending requests are drained.

func (from StreamOfStates) Drop()
    Drop is to be called by a consumer when finished requesting. The request
    channel is closed in order to broadcast this.

    In order to avoid deadlock, pending sends are drained.

func (from StreamOfStates) FanIn2(inp2 StreamOfStates) (out StreamOfStates)
    FanIn2 returns a channel to receive all from both `from` and `inp2` before
    close.

func (from StreamOfStates) From() (req chan<- struct{}, rcv <-chan S)
    From returns the handshaking channels (for use e.g. in `select` statements)
    to receive values:

        `req` to send a request `req <- struct{}{}` and
        `rcv` to reveive such requested value from.

func (from StreamOfStates) Get() (val S, open bool)
    Get is the comma-ok multi-valued form to receive from the channel and
    reports whether a value was received from an open channel or not (as it has
    been closed).

    Get blocks until the request is accepted and value `val` has been received
    from `from` (or until abort is broadcast).

    Get includes housekeeping: If `from` has been closed, `from` is dropped.

func (into StreamOfStates) Into() (req <-chan struct{}, snd chan<- S)
    Into returns the handshaking channels (for use e.g. in `select` statements)
    to send values:

        `req` to receive a request `<-req` and
        `snd` to send such requested value into.

func (c StreamOfStates) Len() int
    Len reports the length of the underlying value channel.

func (c StreamOfStates) New() StreamOfStates
    New returns a new similar channel.

    Useful e.g. when embedded anonymously.

func (into StreamOfStates) Next() (ok bool)
    Next is the request method. It blocks until a request is received and
    reports whether the request channel was open.

    A successful Next is to be followed by one Send(v).

func (into StreamOfStates) NextGetFrom(from StreamOfStates) (val S, ok bool)
    NextGetFrom `from` for `into` and report success.

    Follow it with `into.Send( f(val) )`, if ok.

    NextGetFrom includes housekeeping: If `into` has been dropped or `from` has
    been closed, `from` is dropped and `into` is closed.

func (into StreamOfStates) Provide(val S) (ok bool)
    Provide is the low-level send-upon-request method - aka "myAnyChan <-
    myAny".

    Note: Provide is low-level - its cousin `Put` includes housekeeping: `Put`
    closes the channel upon nok.

    Hint: Provide is useful in constructors together with `defer into.Close()`.

func (into StreamOfStates) Put(val S) (ok bool)
    Put is the send-upon-request method - aka "myAnyChan <- myAny".

    Put blocks until requested to send value `val` into `into` and reports
    whether the request channel was open.

    Put is a convenience for

        if Next() { Send(v) } else { Close() }

    Put includes housekeeping: If `into` has been dropped, `into` is closed.

func (from StreamOfStates) Receive() (val S, open bool)
    Receive is the comma-ok multi-valued form to receive from the channel and
    reports whether a value was received from an open channel or not (as it has
    been closed).

    Receive blocks until the request is accepted and value `val` has been
    received from `from` (or until abort is broadcast).

    Note: Receive is low-level - its cousin `Get` includes housekeeping: `Get`
    drops the channel upon nok.

func (c StreamOfStates) Self() StreamOfStates
    Self returns itself.

    Useful e.g. when embedded anonymously and e.g. wrappers for multi-value
    methods are required.

func (into StreamOfStates) Send(val S) (sent bool)
    Send is to be used after a successful Next()

				
-------------------------------------------------------------------------------
