package main

import (
	"cryptography/finite_field"
	"fmt"
	"log"
	"math/big"
	"os"
)

func main() {
	//console.Menu()
	p := big.NewInt(29)
	Z3 := finite_field.NewFiniteField(p)

	fmt.Println(Z3)

	if err := os.WriteFile("finite_field/cayley_table/add.txt", []byte(Z3.CayleyTableAdd()), 0666); err != nil {
		log.Fatal(err)
	}

	if err := os.WriteFile("finite_field/cayley_table/mul.txt", []byte(Z3.CayleyTableMul()), 0666); err != nil {
		log.Fatal(err)
	}

}
