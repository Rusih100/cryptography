package console

import (
	"bufio"
	"fmt"
	"os"
)

func Menu() {
	fmt.Println()

	runFlag := true

	menuString := "Выберете пункт меню:\n" +
		"1. Расширенный алгоритм Евклида\n" +
		"2. Алгоритм быстрого возведения в степень\n" +
		"3. Алгоритм быстрого возведения в степень по модулю\n" +
		"4. Вычисление символа Якоби\n" +
		"5. Тест Ферма\n" +
		"6. Тест Соловэя-Штрассена\n" +
		"7. Тест Миллера-Рабина\n" +
		"8. Генерация k-битного простого числа\n" +
		"\n" +
		"e - выход из программы\n"

	for runFlag {
		switchFlag := true

		fmt.Print(menuString)

		scanner := bufio.NewScanner(os.Stdin)

		for switchFlag {
			fmt.Print("> ")
			scanner.Scan()
			command := scanner.Text()
			fmt.Println()

			switch command {
			case "1":
				EuclidAlgorithmConsole()
				switchFlag = false
			case "2":
				PowConsole()
				switchFlag = false
			case "3":
				PowModConsole()
				switchFlag = false
			case "4":
				JacobiConsole()
				switchFlag = false
			case "5":
				FermatTestConsole()
				switchFlag = false
			case "6":
				SolovayStrassenTestConsole()
				switchFlag = false
			case "7":
				MillerRabinTestConsole()
				switchFlag = false
			case "8":
				SimpleNumberConsole()
				switchFlag = false

			case "e":
				runFlag = false
				switchFlag = false
			default:
				fmt.Print("Не распознал команду, повторите ввод\n")
			}
		}

	}
}
