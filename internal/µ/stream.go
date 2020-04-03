package µ

type head = S
type tail func() *Stream

type Stream struct {
	head
	tail
}

func mZero() *Stream {
	return nil
}

func nilTail() func() *Stream {
	return func() *Stream {
		return mZero()
	}
}

func Unit(a S) *Stream {
	return &Stream{a, nilTail()}
}

func choice(a S, s func() *Stream) *Stream {
	return &Stream{a, s}
}

func (s *Stream) concat(tail func() *Stream) *Stream {
	if s == mZero() {
		return tail()
	} else {
		return choice(s.head, func() *Stream {
			return s.tail().concat(tail)
		})
	}
}

func (s *Stream) interleave(s2 *Stream) *Stream {
	if s == mZero() {
		return s2
	} else {
		return choice(s.head, func() *Stream {
			return s2.interleave(s.tail())
		})
	}
}
