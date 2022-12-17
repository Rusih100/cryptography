package main

import (
	"cryptography/factorization"
	"fmt"
	"math/big"
)

func main() {

	//console.Menu()

	num := big.NewInt(3)

	for {
		factor := new(big.Int)

		factor = factorization.RoOnePollardFactor(num)
		fmt.Println(factor)

	}

}
