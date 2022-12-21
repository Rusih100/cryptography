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
	num.SetString("1000000000000000000000002", 10)
	fmt.Println("Число =", num)

	fmt.Println("Бит =", num.BitLen())
	fmt.Println("----")

	res := []*big.Int{}

	res = factorization.Factorization(num)

	fmt.Println("---- Результат -----")
	for i := 0; i < len(res); i++ {
		fmt.Print(res[i], " ")
	}
	fmt.Println()

}
