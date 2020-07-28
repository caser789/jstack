package udp

type udpPacketList struct {
	head *udpPacket
	tail *udpPacket
}

func (l *udpPacketList) Reset() {
	l.head = nil
	l.tail = nil
}

func (l *udpPacketList) Empty() bool {
	return l.head == nil
}

func (l *udpPacketList) Front() *udpPacket {
	return l.head
}

func (l *udpPacketList) Back() *udpPacket {
	return l.tail
}

func (l *udpPacketList) PushFront(e *udpPacket) {
	e.SetNext(l.head)
	e.SetPrev(nil)

	if l.head != nil {
		l.head.SetPrev(e)
	} else {
		l.tail = e
	}

	l.head = e
}

func (l *udpPacketList) PushBack(e *udpPacket) {
	e.SetNext(nil)
	e.SetPrev(l.tail)

	if l.tail != nil {
		l.tail.SetNext(e)
	} else {
		l.head = e
	}

	l.tail = e
}

func (l *udpPacketList) PushBackList(m *udpPacketList) {
	if l.head == nil {
		l.head = m.head
		l.tail = m.tail
	} else if m.head != nil {
		l.tail.SetNext(m.head)
		m.head.SetPrev(l.tail)

		l.tail = m.tail
	}

	m.head = nil
	m.tail = nil
}

func (l *udpPacketList) InsertAfter(b, e *udpPacket) {
	a := b.Next()
	e.SetNext(a)
	e.SetPrev(b)
	b.SetNext(e)

	if a != nil {
		a.SetPrev(e)
	} else {
		l.tail = e
	}
}

func (l *udpPacketList) InsertBefore(a, e *udpPacket) {
	b := a.Prev()
	e.SetNext(a)
	e.SetPrev(b)
	a.SetPrev(e)

	if b != nil {
		b.SetNext(e)
	} else {
		l.head = e
	}
}

func (l *udpPacketList) Remove(e *udpPacket) {
	prev := e.Prev()
	next := e.Next()

	if prev != nil {
		prev.SetNext(next)
	} else {
		l.head = next
	}

	if next != nil {
		next.SetPrev(prev)
	} else {
		l.tail = prev
	}
}
