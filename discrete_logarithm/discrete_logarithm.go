package discrete_logarithm

import (
	"cryptography/crypto_math"
	"cryptography/factorization"
	"cryptography/polynomial"
	"math/big"
)

// Константы для упрощения кода
// Не изменять в алгоритмах!
var (
	constNum1 = big.NewInt(1)
	constNum2 = big.NewInt(2)
)

// NumberOrder - Нахождения порядка числа a по простому модулю mod.
//
// Вход: Числа a и mod (Простое)
//
// Выход: Порядок l числа a по модулю mod
func NumberOrder(_a *big.Int, _mod *big.Int) *big.Int {

	// Копируем значения, чтобы не менять по указателю
	a := new(big.Int)
	mod := new(big.Int)

	a.Set(_a)
	mod.Set(_mod)

	// Проверка модуля
	if mod.Cmp(constNum1) <= 0 {
		panic("mod > 1")
	}

	if !crypto_math.MillerRabinTest(mod) {
		panic("mod is a prime number")
	}

	a = a.Mod(a, mod)

	if a.Sign() == 0 {
		panic("a != 0")
	}

	// Нахождение функции Эйлера
	phi := new(big.Int).Sub(mod, constNum1)

	// Нахождение всех делителей phi
	factors := []*big.Int{}

	factors = factorization.Factorization(phi)

	// Поиск порядка
	result := big.NewInt(0)

	temp := new(big.Int)

	for i := 0; i < len(factors); i++ {

		temp = crypto_math.Pow(a, factors[i])
		temp = temp.Sub(temp, constNum1)
		temp = temp.Mod(temp, mod)

		if temp.Sign() == 0 {
			result.Set(factors[i])
			break
		}
	}

	if result.Sign() == 0 {
		return nil
	}

	return result
}

// DiscreteLogarithm - Дискретное логарифмирование.
//
// Вход: a порядка r по модулю p, b
//
// Выход: x
func DiscreteLogarithm(_a *big.Int, _b *big.Int, _p *big.Int) *big.Int {

	// Копируем значения, чтобы не менять по указателю
	a := new(big.Int)
	b := new(big.Int)
	p := new(big.Int)

	a.Set(_a)
	b.Set(_b)
	p.Set(_p)

	// Проверка входных данных
	if !crypto_math.MillerRabinTest(p) {
		panic("p is a prime number")
	}

	b = b.Mod(b, p)
	if b.Sign() == 0 {
		panic("b != 0")
	}

	// Нахождение порядка r числа a
	r := new(big.Int)
	r = NumberOrder(a, p)

	if r == nil {
		return nil
	}

	// p / 2
	p2 := new(big.Int).Div(p, constNum2)

	// Полиномы для ветвящегося отображения
	x1Arr := []*big.Int{
		big.NewInt(1),
	}

	x2Arr := []*big.Int{
		big.NewInt(0),
		big.NewInt(1),
	}

	x1 := polynomial.NewPolynomial(x1Arr)
	x2 := polynomial.NewPolynomial(x2Arr)

	// Ветвящееся отображение
	fx := func(x *big.Int, logX *polynomial.Polynomial) (*big.Int, *polynomial.Polynomial) {

		y := big.NewInt(0)
		logYArr := []*big.Int{
			big.NewInt(0),
		}
		logY := polynomial.NewPolynomial(logYArr)

		if x.Cmp(p2) < 0 {
			y = y.Mod(new(big.Int).Mul(a, x), p)
			logY = logY.Add(logX, x1)
			return y, logY

		} else {
			y = y.Mod(new(big.Int).Mul(b, x), p)
			logY = logY.Add(logX, x2)
			return y, logY
		}
	}

	// 1. Случайны U и V (Полагаем равными 2)
	u := big.NewInt(2)
	v := big.NewInt(2)

	// Переменные
	c := new(big.Int)
	c = c.Mul(
		crypto_math.PowMod(a, u, p),
		crypto_math.PowMod(b, v, p),
	)
	c = c.Mod(c, p)

	d := new(big.Int)
	d.Set(c)

	// Логарифмы
	logArr := []*big.Int{
		new(big.Int).Set(u),
		new(big.Int).Set(v),
	}

	logC := polynomial.NewPolynomial(logArr)
	logD := polynomial.NewPolynomial(logArr)

	for {
		c, logC = fx(c, logC)

		d, logD = fx(d, logD)
		d, logD = fx(d, logD)

		if c.Cmp(d) == 0 {
			break
		}
	}

	logC = logC.Sub(logD, logC)
	logC = logC.Mod(logC, r)

	x := new(big.Int)

	item0 := new(big.Int)
	item1 := new(big.Int)

	item0 = logC.Get(0)
	item1 = logC.Get(1)

	item0 = item0.Neg(item0)
	item0 = item0.Mod(item0, r)

	_, x, _ = crypto_math.ModuloComparisonFirst(item1, item0, r)

	return x
}
