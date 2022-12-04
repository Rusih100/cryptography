package console

import (
	"cryptography/cryptography"
	"fmt"
	"math/big"
)

func EuclidAlgorithmConsole() {

	runFlag := true

	for runFlag {
		switchFlag := true

		fmt.Println("Расширенный алгоритм Евклида")

		fmt.Print("x = ")
		var xString string
		_, _ = fmt.Scan(&xString)

		fmt.Print("y = ")
		var yString string
		_, _ = fmt.Scan(&yString)

		x := new(big.Int)
		x.SetString(xString, 10)

		y := new(big.Int)
		y.SetString(yString, 10)

		m := new(big.Int)
		a := new(big.Int)
		b := new(big.Int)

		m, a, b = cryptography.EuclidAlgorithm(x, y)

		fmt.Println("\nРезультат:")
		fmt.Println("m =", m)
		fmt.Println("a =", a)
		fmt.Println("b =", b)

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

func PowConsole() {

	runFlag := true

	for runFlag {
		switchFlag := true

		fmt.Println("Алгоритм быстрого возведения в степень")

		fmt.Print("a = ")
		var aString string
		_, _ = fmt.Scan(&aString)

		fmt.Print("n = ")
		var nString string
		_, _ = fmt.Scan(&nString)

		a := new(big.Int)
		a.SetString(aString, 10)

		n := new(big.Int)
		n.SetString(nString, 10)

		fmt.Println("\nРезультат:")
		fmt.Println(cryptography.Pow(a, n))

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

func PowModConsole() {

	runFlag := true

	for runFlag {
		switchFlag := true

		fmt.Println("Алгоритм быстрого возведения в степень по модулю")

		fmt.Print("a = ")
		var aString string
		_, _ = fmt.Scan(&aString)

		fmt.Print("n = ")
		var nString string
		_, _ = fmt.Scan(&nString)

		fmt.Print("mod = ")
		var modString string
		_, _ = fmt.Scan(&modString)

		a := new(big.Int)
		a.SetString(aString, 10)

		n := new(big.Int)
		n.SetString(nString, 10)

		mod := new(big.Int)
		mod.SetString(modString, 10)

		res := new(big.Int)
		res = cryptography.PowMod(a, n, mod)

		fmt.Println("\nРезультат:")
		fmt.Println(res)

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

func JacobiConsole() {

	runFlag := true

	for runFlag {
		switchFlag := true

		fmt.Println("Символ Якоби")

		fmt.Print("a = ")
		var aString string
		_, _ = fmt.Scan(&aString)

		fmt.Print("n = ")
		var nString string
		_, _ = fmt.Scan(&nString)

		a := new(big.Int)
		a.SetString(aString, 10)

		n := new(big.Int)
		n.SetString(nString, 10)

		fmt.Println("\nРезультат:")
		fmt.Println(cryptography.Jacobi(a, n))

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
		if cryptography.FermatTest(n) {
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
		if cryptography.SolovayStrassenTest(n) {
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
		if cryptography.MillerRabinTest(n) {
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

func SimpleNumberConsole() {

	runFlag := true

	for runFlag {
		switchFlag := true

		fmt.Println("Генерация k-битного простого числа")

		var k, t int
		fmt.Print("k = ")
		_, _ = fmt.Scan(&k)

		fmt.Print("t = ")
		_, _ = fmt.Scan(&t)

		fmt.Println("\nРезультат:")
		fmt.Println(cryptography.SimpleNumber(k, t))

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

		countSolutions, x1, offset = cryptography.ModuloComparisonFirst(a, b, mod)

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

		x1, x2 = cryptography.ModuloComparisonSecond(a, p)

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

		x = cryptography.ModuloComparisonSystem(bArray, mArray)

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
