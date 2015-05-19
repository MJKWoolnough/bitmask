// Package bitmask implements a simple bitmask type
package bitmask

// BitMask is the main type. It is completely safe to cast from an existing byte
// slice or to create a BitMask with make.
type BitMask []byte

// New creates a new BitMask with enough storage to hold at least the required
// number of bits.
func New(bits int) BitMask {
	numCells := bits >> 3
	if bits&7 != 0 {
		numCells++
	}
	return make(BitMask, numCells)
}

// Get will retrieve the bool stored at given bit position.
func (b BitMask) Get(bit int) bool {
	pos := bit >> 3
	if pos > len(b) {
		return false
	}
	return b[bit>>3]&byte(1<<byte(bit&7)) != 0
}

// Set will set the given bool at the given position.
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
