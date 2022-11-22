package console

import (
	"bufio"
	"cryptography/cryptography"
	"fmt"
	"math/big"
	"os"
)

func EuclidAlgorithmTest() {
	// Простые числа для теста
	// 19636142283486986146701619
	// 141709103505838236078001

	fmt.Println("Тест алгоритма Евклида")

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

	fmt.Println("\nResult:")
	fmt.Println("m =", m)
	fmt.Println("a =", a)
	fmt.Println("b =", b)

	checkM := new(big.Int).Add(
		new(big.Int).Mul(x, a),
		new(big.Int).Mul(y, b),
	)

	fmt.Println("\nCheck result:")
	fmt.Println("m = x * a + y * b =", checkM)
}

func PowTest() {
	fmt.Println("Тест алгоритма быстрого возведения в степень")

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

	fmt.Println(cryptography.Pow(a, n))
}

func PowModTest() {
	fmt.Println("Тест алгоритма быстрого возведения в степень по модулю")

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

	res1 := new(big.Int)
	res1 = cryptography.PowMod(a, n, mod)

	res2 := new(big.Int)
	res2 = cryptography.Pow(a, n)

	fmt.Println("Cтепень по модулю:")
	fmt.Println(res1)

	fmt.Println("Обычная степень:")
	fmt.Println(res2)
}

func JacobiTest() {
	// 219 и 283 -> 1
	// 5 и 19    -> 1
	// 13 и 39   -> 0
	fmt.Println("Тест Символ Якоби")
	for {
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

		fmt.Println("Результат:\n", cryptography.Jacobi(a, n))
		fmt.Println("Результат встроенной функции:\n", big.Jacobi(a, n))
		fmt.Println()
	}
}

func FermatTestTest() {
	fmt.Println("Тест Ферма")
	for {
		scanner := bufio.NewScanner(os.Stdin)

		fmt.Print("n = ")
		scanner.Scan()
		nString := scanner.Text()

		n := new(big.Int)
		n.SetString(nString, 10)

		if cryptography.FermatTest(n) {
			fmt.Println("Число n, вероятно, простое")
		} else {
			fmt.Println("Число n составное")
		}
		fmt.Println()
	}

}

func SolovayStrassenTestTest() {
	fmt.Println("Тест Соловэя-Штрассена")
	for {
		scanner := bufio.NewScanner(os.Stdin)

		fmt.Print("n = ")
		scanner.Scan()
		nString := scanner.Text()

		n := new(big.Int)
		n.SetString(nString, 10)

		if cryptography.SolovayStrassenTest(n) {
			fmt.Println("Число n, вероятно, простое")
		} else {
			fmt.Println("Число n составное")
		}
		fmt.Println()
	}

}

func MillerRabinTestTest() {
	fmt.Println("Тест Миллера-Рабина")
	for {
		scanner := bufio.NewScanner(os.Stdin)

		fmt.Print("n = ")
		scanner.Scan()
		nString := scanner.Text()

		n := new(big.Int)
		n.SetString(nString, 10)

		if cryptography.MillerRabinTest(n) {
			fmt.Println("Число n, вероятно, простое")
		} else {
			fmt.Println("Число n составное")
		}
		fmt.Println()
	}

}

func SimpleNumberTest() {
	fmt.Println("Тест генерации простого числа")

	for {
		var k, t uint
		fmt.Print("k = ")
		fmt.Scan(&k)

		fmt.Print("t = ")
		fmt.Scan(&t)

		fmt.Println()
		fmt.Println(cryptography.SimpleNumber(k, t))
	}
}
