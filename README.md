# bitmask
--
    import "github.com/MJKWoolnough/bitmask"

Package bitmask implements a simple bitmask type

## Usage

#### type BitMask

```go
type BitMask []byte
```

BitMask is the main type. It is completely safe to cast from an existing byte
slice or to create a BitMask with make.

#### func  New

```go
func New(bits int) BitMask
```
New creates a new BitMask with enough storage to hold at least the required
number of bits.

#### func (BitMask) Get

```go
func (b BitMask) Get(bit int) bool
```
Get will retrieve the bool stored at given bit position.

#### func (BitMask) Set

```go
func (b BitMask) Set(bit int, d bool)
```
Set will set the given bool at the given position.
