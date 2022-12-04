package main

import (
	"cryptography/cryptography"
	"fmt"
	"math/big"
)

func main() {
	//console.Menu()

	bArray := []*big.Int{
		big.NewInt(56),
		big.NewInt(23),
		big.NewInt(4),
		big.NewInt(35),
	}

	mArray := []*big.Int{
		big.NewInt(113),
		big.NewInt(51),
		big.NewInt(19),
		big.NewInt(43),
	}

	fmt.Println(cryptography.ModuloComparisonSystem(bArray, mArray))
}
