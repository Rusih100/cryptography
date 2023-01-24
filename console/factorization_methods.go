package console

import (
	"crypto/rand"
	"fmt"
	cryptoMath "github.com/Rusih100/crypto-math"
	"github.com/Rusih100/polynomial"
	"math/big"
)

func FactorizationMenu() {

	runFlag := true

	menuString := "Факторизация:\n" +
		"1. Поиск делителя Ро - методом Полларда\n" +
		"2. Поиск делителя (Ро - 1) - методом Полларда\n" +
		"3. Факторизация числа\n" +
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
			case "3":
				FactorizationConsole()
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

		factor := new(big.Int)

		for i := 0; i < 100 && factor != nil; i++ {

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

			factor = cryptoMath.RoPollardFactor(n, c, poly)
		}

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

		for i := 0; i < 100 && factor != nil; i++ {
			factor = cryptoMath.RoOnePollardFactor(n)
		}

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

func FactorizationConsole() {

	runFlag := true

	for runFlag {
		switchFlag := true

		fmt.Println("Факторизация числа")

		fmt.Print("n = ")
		var nString string
		_, _ = fmt.Scan(&nString)

		n := new(big.Int)
		n.SetString(nString, 10)

		factors := []*big.Int{}
		factors = cryptoMath.Factorization(n)

		fmt.Println("Результат:")
		fmt.Print("[\n")
		for i := 0; i < len(factors); i++ {
			fmt.Println("  ", factors[i])
		}
		fmt.Println("]")

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
