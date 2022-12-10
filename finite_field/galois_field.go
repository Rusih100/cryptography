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

// Add - Складывает два элемента в поле
func (g *GaloisField) Add(a, b *polynomial.Polynomial) *polynomial.Polynomial {

	maxValue := new(big.Int)
	maxValue = cryptography.Pow(g.p, g.n)

	// Проверка на принадлежность полю
	if a.Value(g.p).Sign() < 0 || a.Value(g.p).Cmp(maxValue) >= 0 {
		panic("The element a does not belong to the field")
	}

	if b.Value(g.p).Sign() < 0 || b.Value(g.p).Cmp(maxValue) >= 0 {
		panic("The element b does not belong to the field")
	}

	result := new(polynomial.Polynomial)
	result = result.Add(a, b)
	result = result.Mod(result, g.p)

	return result
}

// Mul - Умножает два элемента в поле
func (g *GaloisField) Mul(a, b *polynomial.Polynomial) *polynomial.Polynomial {

	maxValue := new(big.Int)
	maxValue = cryptography.Pow(g.p, g.n)

	// Проверка на принадлежность полю
	if a.Value(g.p).Sign() < 0 || a.Value(g.p).Cmp(maxValue) >= 0 {
		panic("The element a does not belong to the field")
	}

	if b.Value(g.p).Sign() < 0 || b.Value(g.p).Cmp(maxValue) >= 0 {
		panic("The element b does not belong to the field")
	}

	result := new(polynomial.Polynomial)
	result = result.Mul(a, b)
	result.Mod(result, g.p)

	_, result = result.QuoRem(result, g.mod)
	result.Mod(result, g.p)

	return result
}

// Строковое представление
func (g *GaloisField) String() string {
	return "GF(" + g.p.String() + "^" + g.n.String() + ")"
}

// NewGaloisField - Создает GaloisField и задает ему начальные значения
func NewGaloisField(p *big.Int, n *big.Int, poly *polynomial.Polynomial) *GaloisField {
	return new(GaloisField).Set(p, n, poly)
}
