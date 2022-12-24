package main

import (
	"cryptography/discrete_logarithm"
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
	defer timer("main")()

	//console.Menu()

	a := big.NewInt(14)
	b := big.NewInt(64)
	m := big.NewInt(107)

	fmt.Println(discrete_logarithm.DiscreteLogarithm(a, b, m))
}
