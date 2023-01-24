package console

import (
	"fmt"
	cryptoMath "github.com/Rusih100/crypto-math"
	"math/big"
)

func ModuloComparisonMenu() {

	runFlag := true

	menuString := "Решение сравнений, выберете пункт меню:\n" +
		"1. Решение сравнения первой степени\n" +
		"2. Решение сравнения второй степени\n" +
		"3. Решение системы сравнений\n" +
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
				ModuloComparisonFirstConsole()
				switchFlag = false
			case "2":
				ModuloComparisonSecondConsole()
				switchFlag = false
			case "3":
				ModuloComparisonSystemConsole()
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

func ModuloComparisonFirstConsole() {

	runFlag := true

	for runFlag {
		switchFlag := true

		fmt.Println("Решение сравнения первой степени")

		fmt.Print("a = ")
		var aString string
		_, _ = fmt.Scan(&aString)

		fmt.Print("b = ")
		var bString string
		_, _ = fmt.Scan(&bString)

		fmt.Print("mod = ")
		var modString string
		_, _ = fmt.Scan(&modString)

		a := new(big.Int)
		a.SetString(aString, 10)

		b := new(big.Int)
		b.SetString(bString, 10)

		mod := new(big.Int)
		mod.SetString(modString, 10)

		countSolutions := new(big.Int)
		x1 := new(big.Int)
		offset := new(big.Int)

		countSolutions, x1, offset = cryptoMath.ModuloComparisonFirst(a, b, mod)

		fmt.Println("\nРезультат:")

		if countSolutions.Sign() == 0 {
			fmt.Println("Решений:", countSolutions)
			fmt.Print("[]")

		} else if countSolutions.Cmp(big.NewInt(1)) == 0 {
			fmt.Println("Решений:", countSolutions)
			fmt.Println("[")
			fmt.Println("  ", x1)
			fmt.Print("]")

		} else if countSolutions.Cmp(big.NewInt(15)) >= 0 {
			fmt.Println("Решений:", countSolutions)
			x := new(big.Int).Set(x1)

			fmt.Println("[")
			for i := big.NewInt(0); i.Cmp(big.NewInt(3)) == -1; i = i.Add(i, big.NewInt(1)) {
				fmt.Println("  ", x)
				x = x.Add(x, offset)
			}
			fmt.Println("   ...")
			fmt.Println("   x =", x1, "+ k *", offset)
			fmt.Print("]")

		} else {
			fmt.Println("Решений:", countSolutions)
			x := new(big.Int).Set(x1)

			fmt.Println("[")
			for i := big.NewInt(0); i.Cmp(countSolutions) == -1; i = i.Add(i, big.NewInt(1)) {
				fmt.Println("  ", x)
				x = x.Add(x, offset)
			}
			fmt.Print("]")
		}

		fmt.Println()

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

func ModuloComparisonSecondConsole() {
	runFlag := true

	for runFlag {
		switchFlag := true

		fmt.Println("Решение сравнения второй степени")

		fmt.Print("a = ")
		var aString string
		_, _ = fmt.Scan(&aString)

		fmt.Print("p = ")
		var pString string
		_, _ = fmt.Scan(&pString)

		a := new(big.Int)
		a.SetString(aString, 10)

		p := new(big.Int)
		p.SetString(pString, 10)

		x1 := new(big.Int)
		x2 := new(big.Int)

		x1, x2 = cryptoMath.ModuloComparisonSecond(a, p)

		fmt.Println("\nРезультат:")

		if x1 == nil && x2 == nil {
			fmt.Println("Решений нет")

		} else {
			fmt.Println("x1 =", x1)
			fmt.Println("x2 =", x2)
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

func ModuloComparisonSystemConsole() {
	runFlag := true

	for runFlag {
		switchFlag := true

		fmt.Println("Решение системы сравнений")

		fmt.Print("Размер системы: \nn = ")
		var n int
		_, _ = fmt.Scan(&n)

		var bArray []*big.Int
		var mArray []*big.Int

		var tempString string
		temp := new(big.Int)

		x := new(big.Int)

		fmt.Print("Массив коэфицентов b:\n")
		fmt.Print("[\n")
		for i := 0; i < n; i++ {
			fmt.Print("  ")
			_, _ = fmt.Scan(&tempString)
			temp.SetString(tempString, 10)

			bArray = append(bArray, new(big.Int).Set(temp))
		}
		fmt.Print("]\n")

		fmt.Print("Массив модулей n:\n")
		fmt.Print("[\n")
		for i := 0; i < n; i++ {
			fmt.Print("  ")
			_, _ = fmt.Scan(&tempString)
			temp.SetString(tempString, 10)

			mArray = append(mArray, new(big.Int).Set(temp))
		}
		fmt.Print("]\n")

		x = cryptoMath.ModuloComparisonSystem(bArray, mArray)

		fmt.Println("\nРезультат:")

		if x == nil {
			fmt.Println("Решений нет")
		} else {
			fmt.Println(x)
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
