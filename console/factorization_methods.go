package console

import (
	"crypto/rand"
	"cryptography/factorization"
	"cryptography/polynomial"
	"fmt"
	"math/big"
)

func FactorizationMenu() {

	runFlag := true

	menuString := "Факторизация:\n" +
		"1. Поиск делителя Ро - методом Полларда\n" +
		"2. Поиск делителя (Ро - 1) - методом Полларда\n" +
		"\n" +
		"b - Назад\n"

	for runFlag {
		switchFlag := true

		fmt.Print(menuString)

		for switchFlag {
			fmt.Print("> ")

			var command string
			_, _ = fmt.Scan(&command)

			fmt.Println()

			switch command {
			case "1":
				RoPollardFactorConsole()
				switchFlag = false
			case "2":
				RoOnePollardFactorConsole()
				switchFlag = false

			case "b":
				runFlag = false
				switchFlag = false
			default:
				fmt.Print("Не распознал команду, повторите ввод\n")
			}
		}

	}
}

func RoPollardFactorConsole() {

	runFlag := true

	for runFlag {
		switchFlag := true

		fmt.Println("Поиск делителя Ро - методом Полларда")

		fmt.Print("n = ")
		var nString string
		_, _ = fmt.Scan(&nString)

		n := new(big.Int)
		n.SetString(nString, 10)

		polyArr := []*big.Int{
			big.NewInt(1),
			big.NewInt(0),
			big.NewInt(1),
		}

		poly := polynomial.NewPolynomial(polyArr)

		// Генерируем случайное число c
		c, err := rand.Int(
			rand.Reader,
			new(big.Int).Sub(n, big.NewInt(2)),
		)

		if err != nil {
			panic(err)
		}
		c = c.Add(c, big.NewInt(1))
		//

		factor := new(big.Int)
		factor = factorization.RoPollardFactor(n, c, poly)

		fmt.Println("Результат:")
		if factor != nil {
			fmt.Println(factor)
		} else {
			fmt.Println("Делитель не найден")
		}

		for switchFlag {
			fmt.Print("\nr - повторить,\t b - назад\n")

			fmt.Print("> ")

			var command string
			_, _ = fmt.Scan(&command)

			fmt.Println()

			switch command {
			case "r":
				switchFlag = false
			case "b":
				runFlag = false
				switchFlag = false
			default:
				fmt.Print("Не распознал команду, повторите ввод\n\n")
			}
		}
	}
}

func RoOnePollardFactorConsole() {

	runFlag := true

	for runFlag {
		switchFlag := true

		fmt.Println("Поиск делителя (Ро - 1) - методом Полларда")

		fmt.Print("n = ")
		var nString string
		_, _ = fmt.Scan(&nString)

		n := new(big.Int)
		n.SetString(nString, 10)

		factor := new(big.Int)
		factor = factorization.RoOnePollardFactor(n)

		fmt.Println("Результат:")
		if factor != nil {
			fmt.Println(factor)
		} else {
			fmt.Println("Делитель не найден")
		}

		for switchFlag {
			fmt.Print("\nr - повторить,\t b - назад\n")

			fmt.Print("> ")

			var command string
			_, _ = fmt.Scan(&command)

			fmt.Println()

			switch command {
			case "r":
				switchFlag = false
			case "b":
				runFlag = false
				switchFlag = false
			default:
				fmt.Print("Не распознал команду, повторите ввод\n\n")
			}
		}
	}
}
