package factorization

import (
	"crypto/rand"
	"cryptography/crypto_math"
	"cryptography/polynomial"
	bigFloat "github.com/ALTree/bigfloat"
	"math"
	"math/big"
)

// Константы для упрощения кода
// Не изменять в алгоритмах!
var (
	constNum1 = big.NewInt(1)
	constNum2 = big.NewInt(2)
	constNum3 = big.NewInt(3)
)

// RoPollardFactorization - Ро-метод Полларда.
//
// Вход: Число n, начальное значение c, функция fx (полином)
//
// Выход: Нетривиальный делитель числа p числа n
func RoPollardFactorization(_n *big.Int, _c *big.Int, _fx *polynomial.Polynomial) *big.Int {

	// Копируем значения чтобы не менять значения по указателю
	n := new(big.Int)
	n.Set(_n)

	c := new(big.Int)
	c.Set(_c)

	fx := new(polynomial.Polynomial)
	fx.SetPolynomial(_fx)

	// Проверка входных данных
	if n.Sign() <= 0 {
		panic("n > 0")
	}

	if c.Sign() <= 0 || c.Cmp(n) >= 0 {
		panic("0 < c < n")
	}

	a := new(big.Int)
	b := new(big.Int)

	a.Set(c)
	b.Set(c)

	for {
		a = fx.Value(a)
		a = a.Mod(a, n)

		b = fx.Value(b)
		b = b.Mod(b, n)
		b = fx.Value(b)
		b = b.Mod(b, n)

		d := new(big.Int)
		d = crypto_math.EuclidAlgorithm(new(big.Int).Sub(a, b), n)

		if d.Cmp(constNum1) > 0 && d.Cmp(n) < 0 {
			return d
		}

		if d.Cmp(n) == 0 {
			return nil
		}
	}
}

// RoOnePollardFactorization - (Po - 1)-метод Полларда.
//
// Вход: Число n
//
// Выход: Нетривиальный делитель числа p числа n
func RoOnePollardFactorization(_n *big.Int) *big.Int {

	// Копируем значения чтобы не менять значения по указателю
	n := new(big.Int)
	n.Set(_n)

	// Проверка входных данных
	if n.Sign() <= 0 {
		panic("n > 0")
	}

	// Генерируем случайное число 2 ≤ a < n - 1:
	// 0 ≤ a < n - 3	| +2
	a, err := rand.Int(
		rand.Reader,
		new(big.Int).Sub(n, constNum3),
	)
	if err != nil {
		panic(err)
	}
	a = a.Add(a, constNum2)
	//

	d := new(big.Int)
	d = crypto_math.EuclidAlgorithm(a, n)

	if d.Cmp(constNum2) >= 0 {
		return d
	}

	// Для логарифма
	lnN := new(big.Float).SetInt(n)
	lnN = bigFloat.Log(lnN)

	lnN64, _ := lnN.Float64()

	var lnI float64

	for i := 0; i < len(simpleNumbersBase); i++ {
		lnI = math.Log(
			float64(simpleNumbersBase[i]),
		)

		l := int64(lnN64 / lnI)

		power := new(big.Int)
		power = crypto_math.Pow(
			big.NewInt(simpleNumbersBase[i]),
			big.NewInt(l),
		)

		a = crypto_math.PowMod(a, power, n)
	}

	d = crypto_math.EuclidAlgorithm(new(big.Int).Sub(a, constNum1), n)

	if d.Cmp(constNum1) == 0 || d.Cmp(n) == 0 {
		return nil
	}

	return d
}
