# go-idpool

Basic id pool implementation. Unique ids can be pulled from and released back into the pool as needed.  
Useful for tagging things with a life cycle.

### Example
```go
package main

import (
	"fmt"

	idp "github.com/franklange/go-idpool"
)

func main() {
	p := idp.Idpool{}

	n0 := p.Next()
	fmt.Println(n0) // 0

	n1 := p.Next()
	fmt.Println(n1) // 1

	n2 := p.Next()
	fmt.Println(n2) // 2

	ok := p.Release(n1)
	fmt.Println(ok) // true, release 1 back into the id pool

	n3 := p.Next()
	fmt.Println(n3) // 1, lowest free id in the pool after previous release

	nok := p.Release(99)
	fmt.Println(nok) // false, can't release ids that haven't been given out

	p.SetStart(100) // set first id to be pulled (resets pool)

	n100 := p.Next()
	fmt.Println(n100) // 100
}

```