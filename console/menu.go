package console

import (
	"fmt"
)

func Menu() {
	fmt.Println()

	runFlag := true

	menuString := "Выберете пункт меню:\n" +
		" 1. Расширенный алгоритм Евклида\n" +
		" 2. Алгоритм быстрого возведения в степень\n" +
		" 3. Алгоритм быстрого возведения в степень по модулю\n" +
		" 4. Вычисление символа Якоби\n" +
		" 5. Тесты простоты\n" +
		" 6. Генерация k-битного простого числа\n" +
		" 7. Решение сравнений\n" +
		" 8. Построение конечных полей\n" +
		" 9. Факторизация\n" +
		"10. Дискретное логарифмирование\n" +
		"\n" +
		"e - выход из программы\n"

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
				SimpleNumbersMenu()
				switchFlag = false
			case "6":
				GenerateSimpleNumberConsole()
				switchFlag = false
			case "7":
				ModuloComparisonMenu()
				switchFlag = false
			case "8":
				GaloisFieldMenu()
				switchFlag = false
			case "9":
				FactorizationMenu()
				switchFlag = false
			case "10":
				DiscreteLogarithmConsole()
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
