package main

import (
	"cryptography/polynomial"
	"fmt"
	"math/big"
)

func main() {
	//console.Menu()

	modArr := []*big.Int{
		big.NewInt(1),
		big.NewInt(1),
		big.NewInt(0),
		big.NewInt(0),
		big.NewInt(1),
	}

	offsetArr := []*big.Int{
		big.NewInt(0),
		big.NewInt(1),
	}

	mod := polynomial.NewPolynomial(modArr)
	offset := polynomial.NewPolynomial(offsetArr)

	fmt.Println(mod)
	fmt.Println(offset)
	fmt.Println("-----------")

	xArr := []*big.Int{
		big.NewInt(1),
	}

	x := polynomial.NewPolynomial(xArr)

	for i := 0; i < 16; i++ {
		fmt.Println(x)
		x = x.Mul(offset, x)
		_, x = new(polynomial.Polynomial).QuoRem(x, mod)
		x.Mod(x, big.NewInt(2))

	}

}
