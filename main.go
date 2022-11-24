package main

import (
	"cryptography/cryptography"
	"fmt"
	"math/big"
)

func main() {
	//console.Menu()

	a := big.NewInt(17)
	b := big.NewInt(0)
	mod := big.NewInt(15)

	ls := cryptography.ModuloComparisonFirst(a, b, mod)

	for e := ls.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value, " ")
	}
	fmt.Println()
}
