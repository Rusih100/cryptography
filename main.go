package main

import (
	"cryptography/cryptography"
	"fmt"
	"math/big"
)

func main() {
	//console.Menu()

	a := big.NewInt(12)
	b := big.NewInt(9)
	mod := big.NewInt(15)

	fmt.Println(cryptography.ModuloComparisonFirst(a, b, mod))
}
