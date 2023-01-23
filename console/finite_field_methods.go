package console

import (
	"cryptography/finite_field"
	"fmt"
	"github.com/Rusih100/polynomial"
	"math/big"
)

func GaloisFieldMenu() {

	runFlag := true

	menuString := "Построение конечных полей, выберете пункт меню:\n" +
		"1. Построение базового поля\n" +
		"2. Построение расширения базового поля\n" +
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
				BaseGaloisFieldConsole()
				switchFlag = false
			case "2":
				GaloisFieldConsole()
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

func BaseGaloisFieldConsole() {

	fmt.Println("Построение расширения базового поля")

	fmt.Print("p = ")
	var pString string
	_, _ = fmt.Scan(&pString)

	p := new(big.Int)
	p.SetString(pString, 10)

	fmt.Println()

	GF := finite_field.NewBaseGaloisField(p)
	GF.CayleyTableAdd()
	GF.CayleyTableMul()

	fmt.Println("Поле" + GF.String() + "построено")
	fmt.Print("Сохранено в finite_field/cayley_table" + "\n\n")
}

func GaloisFieldConsole() {

	fmt.Println("Построение базового поля")

	fmt.Print("p = ")
	var pString string
	_, _ = fmt.Scan(&pString)

	p := new(big.Int)
	p.SetString(pString, 10)

	fmt.Print("n = ")
	var nString string
	_, _ = fmt.Scan(&nString)

	n := new(big.Int)
	n.SetString(nString, 10)

	var bArray []*big.Int

	var tempString string
	temp := new(big.Int)

	fmt.Print("Массив коэфицентов многочлена начиная с x^0:\n")
	fmt.Print("[\n")
	for i := big.NewInt(0); i.Cmp(n) <= 0; i.Add(i, big.NewInt(1)) {
		fmt.Print("  ")
		_, _ = fmt.Scan(&tempString)
		temp.SetString(tempString, 10)

		bArray = append(bArray, new(big.Int).Set(temp))
	}

	fmt.Print("]\n")
	poly := polynomial.NewPolynomial(bArray)
	fmt.Println(poly)

	fmt.Println()

	GF := finite_field.NewGaloisField(p, n, poly)
	GF.CayleyTableAdd()
	GF.CayleyTableMul()

	fmt.Println("Поле" + GF.String() + "построено")
	fmt.Print("Сохранено в finite_field/cayley_table" + "\n\n")
}
