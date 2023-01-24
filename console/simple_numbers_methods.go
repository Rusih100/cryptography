package console

import (
	"fmt"
	cryptoMath "github.com/Rusih100/crypto-math"
	"math/big"
)

func SimpleNumbersMenu() {

	runFlag := true

	menuString := "Тесты простоты, выберете пункт меню:\n" +
		"1. Тест Ферма\n" +
		"2. Тест Соловэя-Штрассена\n" +
		"3. Тест Миллера-Рабина\n" +
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
				FermatTestConsole()
				switchFlag = false
			case "2":
				SolovayStrassenTestConsole()
				switchFlag = false
			case "3":
				MillerRabinTestConsole()
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

func FermatTestConsole() {

	runFlag := true

	for runFlag {
		switchFlag := true

		fmt.Println("Тест Ферма")

		fmt.Print("n = ")
		var nString string
		_, _ = fmt.Scan(&nString)

		n := new(big.Int)
		n.SetString(nString, 10)

		fmt.Println()
		if cryptoMath.FermatTest(n) {
			fmt.Println("Число n, вероятно, простое")
		} else {
			fmt.Println("Число n составное")
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

func SolovayStrassenTestConsole() {

	runFlag := true

	for runFlag {
		switchFlag := true

		fmt.Println("Тест Соловэя-Штрассена")

		fmt.Print("n = ")
		var nString string
		_, _ = fmt.Scan(&nString)

		n := new(big.Int)
		n.SetString(nString, 10)

		fmt.Println()
		if cryptoMath.SolovayStrassenTest(n) {
			fmt.Println("Число n, вероятно, простое")
		} else {
			fmt.Println("Число n составное")
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

func MillerRabinTestConsole() {

	runFlag := true

	for runFlag {
		switchFlag := true

		fmt.Println("Тест Миллера-Рабина")

		fmt.Print("n = ")
		var nString string
		_, _ = fmt.Scan(&nString)

		n := new(big.Int)
		n.SetString(nString, 10)

		fmt.Println()
		if cryptoMath.MillerRabinTest(n) {
			fmt.Println("Число n, вероятно, простое")
		} else {
			fmt.Println("Число n составное")
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
