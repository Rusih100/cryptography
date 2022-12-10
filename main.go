package main

import (
	"cryptography/finite_field"
	"cryptography/polynomial"
	"fmt"
	"math/big"
)

func main() {
	//console.Menu()

	modArr := []*big.Int{
		big.NewInt(11),
		big.NewInt(4),
		big.NewInt(0),
		big.NewInt(0),
		big.NewInt(0),
		big.NewInt(1),
	}

	mod := polynomial.NewPolynomial(modArr)

	p := big.NewInt(13)
	n := big.NewInt(5)

	GF := new(finite_field.GaloisField)

	GF.Set(p, n, mod)
	fmt.Println(GF.String())

}
