package main

import (
	"cryptography/factorization"
	"fmt"
	"math/big"
	"time"
)

func timer(name string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", name, time.Since(start))
	}
}

func main() {

	//console.Menu()

	defer timer("main")()

	num := new(big.Int)
	num.SetString("102432638347893748", 10)

	res := []*big.Int{}

	res = factorization.BruteForceFactorization(num)

	for i := 0; i < len(res); i++ {
		fmt.Println(res[i])
	}

}
