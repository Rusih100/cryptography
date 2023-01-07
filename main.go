package main

import (
	"cryptography/ciphers"
	"cryptography/console"
	"encoding/hex"
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

	console.Menu()

	cipherRSA := new(ciphers.RSA)

	cipherRSA.GenerateKey(512)
	//cipherRSA.SaveKeys()

	//cipherRSA.LoadKeys(
	//	"ciphers/RSA/PublicKey_210447.json",
	//	"ciphers/RSA/PrivateKey_210447.json",
	//)

	message := "Если вы найдете человека с которым сможете себя вести также свободно, " +
		"как ведете себя наедине с собой, то цените его как воздух."

	messageBytes := []byte(message)

	cipherMessage := cipherRSA.Encrypt(messageBytes)

	fmt.Println("Зашифрованное сообщение:")
	fmt.Println(hex.EncodeToString(cipherMessage))
	fmt.Println("---")
	fmt.Println("Расшифрованное сообщение:")
	result := cipherRSA.Decrypt(cipherMessage)
	fmt.Println(string(result))

}
