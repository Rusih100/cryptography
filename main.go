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

	cipherRSA.GenerateKey(265)

	//cipherRSA.SaveKeys()

	message := "Если вы найдете человека с которым сможете себя вести также свободно, " +
		"как ведете себя наедине с собой, то цените его как воздух."

	messageBytes := []byte(message)

	cipherText := cipherRSA.Encrypt(messageBytes)

	res := cipherRSA.Decrypt(cipherText)
	fmt.Println(len(res))

}
