package main

import (
	"fmt"

	idp "github.com/franklange/go-idpool"
)

func main() {
	p := idp.Idpool{}

	n0 := p.Next()
	fmt.Println(n0)

	n1 := p.Next()
	fmt.Println(n1)

	n2 := p.Next()
	fmt.Println(n2)

	ok := p.Release(n1)
	fmt.Println(ok)

	n3 := p.Next()
	fmt.Println(n3)

	nok := p.Release(99)
	fmt.Println(nok)

	p.SetStart(100) // reset

	n100 := p.Next()
	fmt.Println(n100)
}
