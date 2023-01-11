package main

import (
	"cryptography/ciphers"
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

func RSA() {
	cipherRSA := new(ciphers.RSA)

	//cipherRSA.GenerateKey(512)
	//cipherRSA.SaveKeys()

	cipherRSA.LoadKeys(
		"ciphers/RSA/PublicKey_191215.json",
		"ciphers/RSA/PrivateKey_191215.json",
	)

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

func Rabin() {
	cipherRabin := new(ciphers.Rabin)

	//cipherRabin.GenerateKey(512)
	//cipherRabin.SaveKeys()
	cipherRabin.LoadKeys(
		"ciphers/Rabin/PublicKey_161015.json",
		"ciphers/Rabin/PrivateKey_161015.json",
	)

	message := "Криптосистема Рабина!"

	messageBytes := []byte(message)

	cipherMessage := cipherRabin.Encrypt(messageBytes)

	fmt.Println("Зашифрованное сообщение:")

	fmt.Println(hex.EncodeToString(cipherMessage))

	fmt.Println("---")
	fmt.Println("Расшифрованные сообщения:")

	res := cipherRabin.Decrypt(cipherMessage)
	fmt.Println(string(res))

}

func main() {
	defer timer("main")()

	Rabin()
	//console.Menu()

}
