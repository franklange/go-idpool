package idpool

import "slices"

type Idpool struct {
	id       uint
	released []uint
}

func (p *Idpool) SetStart(s uint) {
	p.id = s
	p.released = []uint{}
}

func (p *Idpool) Next() uint {

	if len(p.released) > 0 {
		n := p.released[0]
		p.released = p.released[1:]
		return n
	}

	n := p.id
	p.id++

	return n
}

func (p *Idpool) Release(n uint) bool {
	if n > p.id {
		return false
	}

	p.released = append(p.released, n)
	slices.Sort(p.released)

	return true
}
