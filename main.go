package main

import (
	"cryptography/ciphers"
	"fmt"
	"math/big"
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

	message := "Молчи. Смотри на звёзды и цени то, что ты живёшь."
	fmt.Println()
	fmt.Println(message)
	fmt.Println()

	messageBytes := []byte(message)
	fmt.Println("Длина в байтах - ", len(messageBytes))
	fmt.Println(messageBytes)
	fmt.Println("--------------")

	res := []*big.Int{}

	res = ciphers.ToBlocks(messageBytes, 512)
	fmt.Println(res)

	for i := 0; i < len(res); i++ {
		fmt.Println(len(res[i].Bytes()), res[i].Bytes())

	}

}
