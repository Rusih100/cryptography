package console

import (
	"bufio"
	"cryptography/cryptography"
	"fmt"
	"math/big"
	"os"
)

func EuclidAlgorithmTest() {

	runFlag := true

	for runFlag {
		switchFlag := true

		fmt.Println("Расширенный алгоритм Евклида")

		scanner := bufio.NewScanner(os.Stdin)

		fmt.Print("x = ")
		scanner.Scan()
		xString := scanner.Text()

		fmt.Print("y = ")
		scanner.Scan()
		yString := scanner.Text()

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
			scanner.Scan()
			command := scanner.Text()
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

func PowTest() {

	runFlag := true

	for runFlag {
		switchFlag := true

		fmt.Println("Алгоритм быстрого возведения в степень")

		scanner := bufio.NewScanner(os.Stdin)

		fmt.Print("a = ")
		scanner.Scan()
		aString := scanner.Text()

		fmt.Print("n = ")
		scanner.Scan()
		bString := scanner.Text()

		a := new(big.Int)
		a.SetString(aString, 10)

		n := new(big.Int)
		n.SetString(bString, 10)

		fmt.Println("\nРезультат:")
		fmt.Println(cryptography.Pow(a, n))

		for switchFlag {
			fmt.Print("\nr - повторить,\t b - назад\n")

			fmt.Print("> ")
			scanner.Scan()
			command := scanner.Text()
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

func PowModTest() {

	runFlag := true

	for runFlag {
		switchFlag := true

		fmt.Println("Алгоритм быстрого возведения в степень по модулю")

		scanner := bufio.NewScanner(os.Stdin)

		fmt.Print("a = ")
		scanner.Scan()
		aString := scanner.Text()

		fmt.Print("n = ")
		scanner.Scan()
		nString := scanner.Text()

		fmt.Print("mod = ")
		scanner.Scan()
		modString := scanner.Text()

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
			scanner.Scan()
			command := scanner.Text()
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

func JacobiTest() {

	runFlag := true

	for runFlag {
		switchFlag := true

		fmt.Println("Символ Якоби")

		scanner := bufio.NewScanner(os.Stdin)

		fmt.Print("a = ")
		scanner.Scan()
		aString := scanner.Text()

		fmt.Print("n = ")
		scanner.Scan()
		nString := scanner.Text()

		a := new(big.Int)
		a.SetString(aString, 10)

		n := new(big.Int)
		n.SetString(nString, 10)

		fmt.Println("\nРезультат:")
		fmt.Println(cryptography.Jacobi(a, n))

		for switchFlag {
			fmt.Print("\nr - повторить,\t b - назад\n")

			fmt.Print("> ")
			scanner.Scan()
			command := scanner.Text()
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

func FermatTestTest() {

	runFlag := true

	for runFlag {
		switchFlag := true

		fmt.Println("Тест Ферма")

		scanner := bufio.NewScanner(os.Stdin)

		fmt.Print("n = ")
		scanner.Scan()
		nString := scanner.Text()

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
			scanner.Scan()
			command := scanner.Text()
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

func SolovayStrassenTestTest() {

	runFlag := true

	for runFlag {
		switchFlag := true

		fmt.Println("Тест Соловэя-Штрассена")

		scanner := bufio.NewScanner(os.Stdin)

		fmt.Print("n = ")
		scanner.Scan()
		nString := scanner.Text()

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
			scanner.Scan()
			command := scanner.Text()
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

func MillerRabinTestTest() {

	runFlag := true

	for runFlag {
		switchFlag := true

		fmt.Println("Тест Миллера-Рабина")

		scanner := bufio.NewScanner(os.Stdin)

		fmt.Print("n = ")
		scanner.Scan()
		nString := scanner.Text()

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
			scanner.Scan()
			command := scanner.Text()
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

func SimpleNumberTest() {

	runFlag := true

	for runFlag {
		switchFlag := true

		fmt.Println("Генерация k-битного простого числа")

		var k, t uint
		fmt.Print("k = ")
		fmt.Scan(&k)

		fmt.Print("t = ")
		fmt.Scan(&t)

		// FIXME: бага с пропущенным вводом

		fmt.Println("\nРезультат:")
		fmt.Println(cryptography.SimpleNumber(k, t))

		for switchFlag {
			fmt.Print("\nr - повторить,\t b - назад\n")

			fmt.Print("> ")
			scanner := bufio.NewScanner(os.Stdin)

			scanner.Scan()
			command := scanner.Text()
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
