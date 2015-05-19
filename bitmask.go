package bitmask

type BitMask []byte

func New(bits int) BitMask {
	numCells := bits >> 3
	if bits&7 != 0 {
		numCells++
	}
	return make(BitMask, numCells)
}

func (b BitMask) Get(bit int) bool {
	pos := bit >> 3
	if pos > len(b) {
		return false
	}
	return b[bit>>3]&byte(1<<byte(bit&7)) != 0
}

func (b BitMask) Set(bit int, d bool) {
	pos := bit >> 3
	if pos > len(b) {
		return
	}
	shift := byte(1 << byte(bit&7))
	if d {
		b[pos] |= shift
	} else {
		b[pos] &^= shift
	}
}
