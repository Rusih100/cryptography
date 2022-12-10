package main

import (
	"cryptography/finite_field"
	"cryptography/polynomial"
	"fmt"
	"math/big"
)

func main() {
	//console.Menu()

	p := big.NewInt(2)
	n := big.NewInt(3)

	polyArr := []*big.Int{
		big.NewInt(1),
		big.NewInt(1),
		big.NewInt(0),
		big.NewInt(1),
	}

	poly := polynomial.NewPolynomial(polyArr)

	GF := finite_field.NewGaloisField(p, n, poly)
	fmt.Println(GF)

	GF.CayleyTableAdd()

}
