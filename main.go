package main

import (
	"cryptography/polynomial"
	"fmt"
	"math/big"
)

func main() {
	//console.Menu()

	arr1 := []*big.Int{big.NewInt(-1), big.NewInt(1)}
	arr2 := []*big.Int{big.NewInt(1), big.NewInt(1)}

	x1 := new(polynomial.Polynomial).Set(arr1)
	x2 := new(polynomial.Polynomial).Set(arr2)

	x3 := new(polynomial.Polynomial).Mul(x1, x2)

	fmt.Println(x1)
	fmt.Println(x2)
	fmt.Println(x3)

}
