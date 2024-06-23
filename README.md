# coll

`coll` is a Go library providing utility functions for working with collections, such as reversing slices and converting
slices to sets.

## Installation

To install the package, use `go get`:

```sh
go get github.com/GiGurra/coll
```

## Usage

### Reversed

The `Reversed` function returns a new slice with the elements in reverse order.

```go
package main

import (
	"fmt"
	"github.com/GiGurra/coll"
)

func main() {
	items := []string{"a", "b", "c"}
	reversed := coll.Reversed(items)
	fmt.Println(reversed) // Output: [c b a]
}
```

### ToSet

The `ToSet` function converts a slice to a set represented as a map.

```go
package main

import (
	"fmt"
	"github.com/GiGurra/coll"
)

func main() {
	items := []string{"a", "b", "c", "a"}
	set := coll.ToSet(items)
	fmt.Println(set) // Output: map[a:true b:true c:true]
}
```

### ToSetF

The `ToSetF` function converts a slice to a set using a key function to determine the set keys.

```go
package main

import (
	"fmt"
	"github.com/GiGurra/coll"
)

type FancyStruct struct {
	A string
	B int
}

func main() {
	items := []FancyStruct{
		{"a", 1},
		{"b", 2},
		{"c", 3},
	}
	set := coll.ToSetF(items, func(item FancyStruct) string {
		return item.A
	})
	fmt.Println(set) // Output: map[a:true b:true c:true]
}
```

## Testing

To run the tests, use the `go test` command:

```sh
go test ./...
```

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

---

*This README was AI-generated and may require further editing for accuracy and completeness.*
