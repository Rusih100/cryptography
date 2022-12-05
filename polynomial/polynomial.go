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

// TODO: Переписать с учетом минуса и знаков

// Представление полинома в виде строки
func (c *Polynomial) String() string {

	lenCoefficients := len(c.coefficients)
	result := ""

	if lenCoefficients == 1 && c.coefficients[0].Sign() == 0 {
		return c.coefficients[0].String()
	}

	for i := lenCoefficients - 1; i >= 0; i-- {

		if c.coefficients[i].Cmp(constNum1) == 0 {

			if i == 0 {
				result = result + c.coefficients[i].String()
			} else if i == 1 {
				result = result + "x"
			} else {
				result = result + "x^" + strconv.Itoa(i)
			}

			if i != 0 && c.coefficients[i].Cmp(constNum1) != 0 {
				result = result + " + "
			}

		} else if c.coefficients[i].Sign() != 0 {

			if i == 0 {
				result = result + c.coefficients[i].String()
			} else if i == 1 {
				result = result + c.coefficients[i].String() + "x"
			} else {
				result = result + c.coefficients[i].String() + "x^" + strconv.Itoa(i)
			}

			if i != 0 && c.coefficients[i].Cmp(constNum1) != 0 {
				result = result + " + "
			}

		}
	}
	return result
}

// NewPolynomial - Создает полином и задает ему массив значений при создании
func NewPolynomial(coefficients []*big.Int) *Polynomial {
	return new(Polynomial).Set(coefficients)
}
