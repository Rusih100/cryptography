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

func ElGamal() {

	cipherElGamal := new(ciphers.ElGamal)

	//cipherElGamal.GenerateKey(512)
	//cipherElGamal.SaveKeys()

	cipherElGamal.LoadKeys(
		"ciphers/ElGamal/PublicKey_205240.json",
		"ciphers/ElGamal/PrivateKey_205240.json",
	)

	message := "Слова лишь мешают понимать друг друга"

	messageBytes := []byte(message)

	cipherMessage1, cipherMessage2 := cipherElGamal.Encrypt(messageBytes)

	fmt.Println("Зашифрованные сообщения:")
	fmt.Println("1.")
	fmt.Println(hex.EncodeToString(cipherMessage1))
	fmt.Println("2.")
	fmt.Println(hex.EncodeToString(cipherMessage2))
	fmt.Println("---")

	fmt.Println("Расшифрованное сообщение:")
	result := cipherElGamal.Decrypt(cipherMessage1, cipherMessage2)

	fmt.Println(string(result))

}

func main() {
	defer timer("main")()

	console.Menu()
}
