# bitmask

[![CI](https://github.com/MJKWoolnough/bitmask/actions/workflows/go-checks.yml/badge.svg)](https://github.com/MJKWoolnough/bitmask/actions)
[![Go Reference](https://pkg.go.dev/badge/vimagination.zapto.org/bitmask.svg)](https://pkg.go.dev/vimagination.zapto.org/bitmask)
[![Go Report Card](https://goreportcard.com/badge/vimagination.zapto.org/bitmask)](https://goreportcard.com/report/vimagination.zapto.org/bitmask)

--
    import "vimagination.zapto.org/bitmask"

Package bitmask implements a simple bitmask type.

## Highlights

 - Create artibrary size bit fields.
 - Easily get and set bits.
 - Simple method to set bits and determine whether the bit changed.

## Usage

```go
package main

import (
	"fmt"

	"vimagination.zapto.org/bitmask"
)

func main() {
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
```

## Documentation

Full API docs can be found at:

https://pkg.go.dev/vimagination.zapto.org/bitmask
