package bitmask_test

import (
	"fmt"

	"vimagination.zapto.org/bitmask"
)

func Example() {
	bm := bitmask.New(16)

	bm.Set(3, true)
	bm.Set(12, true)

	fmt.Println("Bit 3:", bm.Get(3))
	fmt.Println("Bit 5:", bm.Get(5))
	fmt.Println("Bit 12:", bm.Get(12))

	changed := bm.SetIfNot(12, false)
	fmt.Println("Bit 12 changed:", changed)
	fmt.Println("Bit 12:", bm.Get(12))

	// Output:
	// Bit 3: true
	// Bit 5: false
	// Bit 12: true
	// Bit 12 changed: true
	// Bit 12: false
}
