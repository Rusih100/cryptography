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

	// Проверка входных данных
	if lenCoefficients == 0 {
		panic("An empty array was passed")
	}

	for i := 0; i < lenCoefficients; i++ {
		c.coefficients = append(c.coefficients, big.NewInt(0))
	}

	for i := 0; i < lenCoefficients; i++ {
		c.coefficients[i].Set(coefficients[i])
	}

	return c
}

// Представление полинома в виде строки
func (c *Polynomial) String() string {

	lenCoefficients := len(c.coefficients)
	result := ""

	for i := lenCoefficients - 1; i >= 0; i-- {

		if c.coefficients[i].Cmp(constNum1) == 0 {

			if i == 0 {
				result = result + c.coefficients[i].String()
			} else if i == 1 {
				result = result + "x"
			} else {
				result = result + "x^" + strconv.Itoa(i)
			}

			if i != 0 {
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

			if i != 0 {
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
