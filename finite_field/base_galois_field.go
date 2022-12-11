package finite_field

import (
	"cryptography/cryptography"
	"math/big"
	"os"
)

// Константы для упрощения кода
// Не изменять в алгоритмах!
var (
	constNum1 = big.NewInt(1)
	constNum2 = big.NewInt(2)
)

// Реализация базового конечного поля
// Операции выполняются по модулю p

type BaseGaloisField struct {
	p *big.Int
}

// Set - Задает начальное значение p
func (f *BaseGaloisField) Set(p *big.Int) *BaseGaloisField {

	f.p = big.NewInt(0)

	if p.Cmp(constNum2) < 0 {
		panic("p >= 2")
	}

	if !cryptography.MillerRabinTest(p) {
		panic("p is a prime number")
	}

	f.p.Set(p)

	return f
}

// Add - Складывает два элемента в поле
func (f *BaseGaloisField) Add(a, b *big.Int) *big.Int {

	// Проверки
	if a.Sign() < 0 || a.Cmp(f.p) >= 0 {
		panic("The element a does not belong to the field")
	}

	if b.Sign() < 0 || b.Cmp(f.p) >= 0 {
		panic("The element b does not belong to the field")
	}

	return new(big.Int).Mod(
		new(big.Int).Add(a, b),
		f.p,
	)

}

// Mul - Умножает два элемента в поле
func (f *BaseGaloisField) Mul(a, b *big.Int) *big.Int {

	// Проверки
	if a.Sign() < 0 || a.Cmp(f.p) >= 0 {
		panic("The element a does not belong to the field")
	}

	if b.Sign() < 0 || b.Cmp(f.p) >= 0 {
		panic("The element b does not belong to the field")
	}

	return new(big.Int).Mod(
		new(big.Int).Mul(a, b),
		f.p,
	)
}

// Строковое представление
func (f *BaseGaloisField) String() string {
	return "GF(" + f.p.String() + ")"
}

// CayleyTableAdd - Таблица Кэли для сложения
//
// Файл сохраняется в finite_field/cayley_table
func (f *BaseGaloisField) CayleyTableAdd() {

	name := f.p.String() + "_add"

	// Создание файла
	file, err := os.Create("finite_field/cayley_table/" + name + ".csv")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	// Первая строка
	result := "\t"

	temp := new(big.Int)

	for i := big.NewInt(0); i.Cmp(f.p) < 0; i.Add(i, constNum1) {
		result = result + i.String()

		if i.Cmp(new(big.Int).Sub(f.p, constNum1)) != 0 {
			result = result + "\t"
		}
	}
	result = result + "\n"

	_, err = file.WriteString(result)
	if err != nil {
		panic(err)
	}

	result = ""

	for i := big.NewInt(0); i.Cmp(f.p) < 0; i.Add(i, constNum1) {
		result = i.String() + "\t"

		for j := big.NewInt(0); j.Cmp(f.p) < 0; j.Add(j, constNum1) {
			temp = f.Add(i, j)
			result = result + temp.String()

			if j.Cmp(new(big.Int).Sub(f.p, constNum1)) != 0 {
				result = result + "\t"
			}
		}
		result = result + "\n"
		_, err = file.WriteString(result)
		if err != nil {
			panic(err)
		}
	}
}

// CayleyTableMul - Таблица Кэли для умножения
//
// Файл сохраняется в finite_field/cayley_table
func (f *BaseGaloisField) CayleyTableMul() {

	name := f.p.String() + "_mul"

	// Создание файла
	file, err := os.Create("finite_field/cayley_table/" + name + ".csv")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	// Первая строка
	result := "\t"

	temp := new(big.Int)

	for i := big.NewInt(0); i.Cmp(f.p) < 0; i.Add(i, constNum1) {
		result = result + i.String()

		if i.Cmp(new(big.Int).Sub(f.p, constNum1)) != 0 {
			result = result + "\t"
		}
	}
	result = result + "\n"

	_, err = file.WriteString(result)
	if err != nil {
		panic(err)
	}

	result = ""

	for i := big.NewInt(0); i.Cmp(f.p) < 0; i.Add(i, constNum1) {
		result = i.String() + "\t"

		for j := big.NewInt(0); j.Cmp(f.p) < 0; j.Add(j, constNum1) {
			temp = f.Mul(i, j)
			temp.Mod(temp, f.p)

			result = result + temp.String()

			if j.Cmp(new(big.Int).Sub(f.p, constNum1)) != 0 {
				result = result + "\t"
			}
		}
		result = result + "\n"
		_, err = file.WriteString(result)
		if err != nil {
			panic(err)
		}
	}
}

// NewFiniteField - Создает BaseGaloisField и задает ему начальное значение p
func NewFiniteField(p *big.Int) *BaseGaloisField {
	return new(BaseGaloisField).Set(p)
}
