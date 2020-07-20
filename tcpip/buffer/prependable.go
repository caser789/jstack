package buffer

type Prependable struct {
	// Buf is the buffer backing the prependable buffer.
	buf View

	// usedIdx is the index where the used part of the buffer begins.
	usedIdx int
}

func NewPrependable(size int) Prependable {
	return Prependable{buf: NewView(size), usedIdx: size}
}

func (p *Prependable) Prepend(size int) []byte {
	if size > p.usedIdx {
		return nil
	}

	p.usedIdx -= size
	return p.buf[p.usedIdx:][:size:size]
}

func (p *Prependable) UsedLength() int {
	return len(p.buf) - p.usedIdx
}
