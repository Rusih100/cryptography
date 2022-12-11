package main

import (
	"cryptography/finite_field"
	"cryptography/polynomial"
	"fmt"
	"math/big"
)

func main() {
	//console.Menu()

	p := big.NewInt(11)
	n := big.NewInt(2)

	polyArr := []*big.Int{
		big.NewInt(2),
		big.NewInt(7),
		big.NewInt(1),
	}

	poly := polynomial.NewPolynomial(polyArr)

	GF := finite_field.NewGaloisField(p, n, poly)
	fmt.Println(GF)

	GF.CayleyTableAdd()
	GF.CayleyTableMul()

}
