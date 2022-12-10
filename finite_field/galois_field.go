package finite_field

import (
	"cryptography/cryptography"
	"cryptography/polynomial"
	"math/big"
)

// Реализация расширения базового конечного поля

type GaloisField struct {
	p   *big.Int
	n   *big.Int
	mod *polynomial.Polynomial
}

// Set - Задает начальные значения и проверяет может ли поле быть построено
func (g *GaloisField) Set(p *big.Int, n *big.Int, poly *polynomial.Polynomial) *GaloisField {

	// Проверка входных данных

	// Копируем и проверяем p
	g.p = big.NewInt(0)

	if p.Cmp(constNum2) < 0 {
		panic("p >= 2")
	}

	if !cryptography.MillerRabinTest(p) {
		panic("p is a prime number")
	}

	g.p.Set(p)

	// Копируем и проверяем n

	if n.Cmp(constNum1) <= 0 {
		panic("n > 1")
	}

	g.n = big.NewInt(0)
	g.n.Set(n)

	// Копируем Полином
	g.mod = new(polynomial.Polynomial)
	g.mod.SetPolynomial(poly)

	// Проверка на количество элементов
	offsetArr := []*big.Int{
		big.NewInt(0),
		big.NewInt(1),
	}
	offset := polynomial.NewPolynomial(offsetArr)

	// Счетчик элементов
	counter := big.NewInt(2)

	// Начальный полином
	xArr := []*big.Int{
		big.NewInt(0),
		big.NewInt(1),
	}

	x := polynomial.NewPolynomial(xArr)

	for x.Value(g.p).Cmp(constNum1) != 0 {
		x = x.Mul(offset, x)
		_, x = new(polynomial.Polynomial).QuoRem(x, g.mod)
		x.Mod(x, g.p)

		counter.Add(counter, constNum1)
	}

	want := new(big.Int)
	want = cryptography.Pow(g.p, g.n)

	if want.Cmp(counter) != 0 {
		panic("The field cannot be created")
	}

	return g
}

// Строковое представление
func (g *GaloisField) String() string {
	return "GF(" + g.p.String() + "^" + g.n.String() + ")"
}
