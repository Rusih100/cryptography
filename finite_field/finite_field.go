package finite_field

import (
	"cryptography/cryptography"
	"math/big"
)

// Константы для упрощения кода
// Не изменять в алгоритмах!
var (
	constNum1 = big.NewInt(1)
	constNum2 = big.NewInt(2)
)

// Реализация базового конечного поля
// Операции выполняются по модулю p

type FiniteField struct {
	p *big.Int
}

// Set - Задает начальное значение p
func (f *FiniteField) Set(p *big.Int) *FiniteField {

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
func (f *FiniteField) Add(a, b *big.Int) *big.Int {

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
func (f *FiniteField) Mul(a, b *big.Int) *big.Int {

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
func (f *FiniteField) String() string {
	return "GF(" + f.p.String() + ")"
}

// CayleyTableAdd - Таблица Кэли для сложения
func (f *FiniteField) CayleyTableAdd() string {
	result := "+\t"

	temp := new(big.Int)
	for i := big.NewInt(0); i.Cmp(f.p) < 0; i.Add(i, constNum1) {
		result = result + i.String() + "\t"
	}
	result = result + "\n"

	for i := big.NewInt(0); i.Cmp(f.p) < 0; i.Add(i, constNum1) {
		result = result + i.String() + "\t"

		for j := big.NewInt(0); j.Cmp(f.p) < 0; j.Add(j, constNum1) {
			temp = f.Add(i, j)
			result = result + temp.String() + "\t"
		}
		result = result + "\n"
	}
	return result
}

// CayleyTableMul - Таблица Кэли для умножения
func (f *FiniteField) CayleyTableMul() string {
	result := "*\t"

	temp := new(big.Int)
	for i := big.NewInt(1); i.Cmp(f.p) < 0; i.Add(i, constNum1) {
		result = result + i.String() + "\t"
	}
	result = result + "\n"

	for i := big.NewInt(1); i.Cmp(f.p) < 0; i.Add(i, constNum1) {
		result = result + i.String() + "\t"

		for j := big.NewInt(1); j.Cmp(f.p) < 0; j.Add(j, constNum1) {
			temp = f.Mul(i, j)
			result = result + temp.String() + "\t"
		}
		result = result + "\n"
	}
	return result
}

// NewFiniteField - Создает FiniteField и задает ему начальное значение p
func NewFiniteField(p *big.Int) *FiniteField {
	return new(FiniteField).Set(p)
}
