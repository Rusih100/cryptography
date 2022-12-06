package main

import (
	"cryptography/polynomial"
	"fmt"
	"math/big"
)

func main() {
	//console.Menu()

	arr1 := []*big.Int{
		big.NewInt(1),
		big.NewInt(1),
		big.NewInt(1),
	}

	x1 := new(polynomial.Polynomial).Set(arr1)

	fmt.Println(x1)

	//arr1 := []*big.Int{
	//	big.NewInt(0),
	//	big.NewInt(0),
	//	big.NewInt(2),
	//	big.NewInt(0),
	//}
	//
	//arr2 := []*big.Int{
	//	big.NewInt(0),
	//	big.NewInt(1),
	//	big.NewInt(1),
	//}
	//
	//x1 := new(polynomial.Polynomial).Set(arr1)
	//x2 := new(polynomial.Polynomial).Set(arr2)
	//
	//quo := new(polynomial.Polynomial)
	//rem := new(polynomial.Polynomial)
	//
	//quo, rem = quo.QuoRem(x1, x2)
	//
	//fmt.Println(x1)
	//fmt.Println(x2)
	//fmt.Println("-----------")
	//fmt.Println(quo)
	//fmt.Println(rem)
}
