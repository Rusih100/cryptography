package polynomial

import (
	"math/big"
	"strconv"
)

// Константы для упрощения кода
// Не изменять в алгоритмах!
var (
	constNum1 = big.NewInt(1)
)

// Реализация многочлена и операций над ним

type Polynomial struct {
	coefficients []*big.Int
}

// Set - Задает массив коэфицентов, начиная со свободного члена многочлена.
func (c *Polynomial) Set(coefficients []*big.Int) *Polynomial {

	c.coefficients = []*big.Int{}

	lenCoefficients := len(coefficients)

	if lenCoefficients == 0 {
		c.coefficients = append(c.coefficients, big.NewInt(0))
		return c
	}

	for i := 0; i < lenCoefficients; i++ {
		c.coefficients = append(c.coefficients, big.NewInt(0))
	}

	for i := 0; i < lenCoefficients; i++ {
		c.coefficients[i].Set(coefficients[i])
	}

	return c
}

// Add - Складывает два многочлена a и b, и записывает в c
func (c *Polynomial) Add(a, b *Polynomial) *Polynomial {

	c.coefficients = []*big.Int{}

	aLen := len(a.coefficients)
	bLen := len(b.coefficients)
	maxLen := 0

	if aLen > bLen {
		maxLen = aLen
	} else {
		maxLen = bLen
	}

	for i := 0; i < maxLen; i++ {
		c.coefficients = append(c.coefficients, big.NewInt(0))
	}

	for i := 0; i < aLen; i++ {
		c.coefficients[i].Add(c.coefficients[i], a.coefficients[i])
	}

	for i := 0; i < bLen; i++ {
		c.coefficients[i].Add(c.coefficients[i], b.coefficients[i])
	}

	return c
}

// Sub - вычитает из многочлена a многочлен b, и записывает в c
func (c *Polynomial) Sub(a, b *Polynomial) *Polynomial {

	c.coefficients = []*big.Int{}

	aLen := len(a.coefficients)
	bLen := len(b.coefficients)
	maxLen := 0

	if aLen > bLen {
		maxLen = aLen
	} else {
		maxLen = bLen
	}

	for i := 0; i < maxLen; i++ {
		c.coefficients = append(c.coefficients, big.NewInt(0))
	}

	for i := 0; i < aLen; i++ {
		c.coefficients[i].Add(c.coefficients[i], a.coefficients[i])
	}

	for i := 0; i < bLen; i++ {
		c.coefficients[i].Sub(c.coefficients[i], b.coefficients[i])
	}

	return c
}

// Mul - Умножает два многочлена a и b, и записывает в c
func (c *Polynomial) Mul(a, b *Polynomial) *Polynomial {

	c.coefficients = []*big.Int{}

	aLen := len(a.coefficients)
	bLen := len(b.coefficients)
	maxLen := aLen + bLen - 1

	for i := 0; i < maxLen; i++ {
		c.coefficients = append(c.coefficients, big.NewInt(0))
	}

	for i := 0; i < aLen; i++ {
		for j := 0; j < bLen; j++ {
			c.coefficients[i+j].Add(c.coefficients[i+j], new(big.Int).Mul(a.coefficients[i], b.coefficients[j]))
		}
	}

	return c
}

// Представление полинома в виде строки
func (c *Polynomial) String() string {

	lenCoefficients := len(c.coefficients)
	result := ""

	temp := big.NewInt(0)

	// Случай длины 1
	if lenCoefficients == 1 && c.coefficients[0].Sign() == 0 {
		return "0 "
	}

	// Общий случай
	for i := lenCoefficients - 1; i >= 0; i-- {

		temp.Set(c.coefficients[i])

		// Печатаем число если оно не ноль
		if temp.Sign() != 0 {

			// Смотрим на знак
			// Добавляем плюс, если число положительное и не самая большая степень
			if temp.Sign() > 0 && i != lenCoefficients-1 {
				result = result + "+ "
			}
			// Добавляем минус, если число отрицательное
			if temp.Sign() < 0 {
				result = result + "-"
				if i != lenCoefficients-1 {
					result = result + " "
				}

			}

			// Смотрим на на коэфицент по модулю
			temp = temp.Abs(temp)

			// Печатаем если коэфицент не равен 1
			if temp.Cmp(constNum1) == 0 && i == 0 {
				result = result + temp.String()

			} else if temp.Cmp(constNum1) != 0 {
				result = result + temp.String()
			}

			// Смотрим на степень и ставим x
			if i != 0 {
				result = result + "x"
			}

			// Печатем степень
			if i != 0 && i != 1 {
				result = result + "^" + strconv.Itoa(i) + " "
			} else {
				result = result + " "
			}

		}

	}

	return result
}

// NewPolynomial - Создает полином и задает ему массив значений при создании
func NewPolynomial(coefficients []*big.Int) *Polynomial {
	return new(Polynomial).Set(coefficients)
}
