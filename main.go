package main

import (
	"cryptography/polynomial"
	"fmt"
	"math/big"
)

func main() {
	//console.Menu()

	arr := []*big.Int{big.NewInt(1), big.NewInt(1), big.NewInt(5)}

	x1 := new(polynomial.Polynomial).Set(arr)
	fmt.Println(*x1)

	fmt.Println(x1.String())

	arr[0].Set(big.NewInt(4))
	arr[1].Set(big.NewInt(4))
	fmt.Println(x1.String())
}
