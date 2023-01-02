package main

import (
	"cryptography/ciphers"
	"fmt"
	"time"
)

func timer(name string) func() {
	start := time.Now()
	return func() {
		fmt.Print("\n\nВремя выполнения: ")
		fmt.Printf("%s - %v", name, time.Since(start))
	}
}

func main() {
	defer timer("main")()

	//console.Menu()

	cipherRSA := new(ciphers.RSA)

	//cipherRSA.GenerateKey(512)

	//cipherRSA.SaveKeys()

	cipherRSA.LoadKeys(
		"ciphers/RSA/PublicKey_160849.json",
		"ciphers/RSA/PrivateKey_ 160849.json",
	)

	//n := big.NewInt(131232)
	//d := big.NewInt(3434343434)
	//
	//k := ciphers.PublicKeyRSA{d, n}
	//fmt.Println(k)

}
