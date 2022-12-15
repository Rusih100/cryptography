package main

import (
	"cryptography/factorization"
	"fmt"
	"math/big"
)

func main() {

	//console.Menu()

	num := big.NewInt(2873)

	factor := new(big.Int)

	factor = factorization.RoOnePollardFactorization(num)
	fmt.Println(factor)

}
