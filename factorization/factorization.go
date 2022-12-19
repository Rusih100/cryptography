package factorization

import (
	"crypto/rand"
	"cryptography/crypto_math"
	"cryptography/polynomial"
	bigFloat "github.com/ALTree/bigfloat"
	"math"
	"math/big"
	"sort"
)

// Константы для упрощения кода
// Не изменять в алгоритмах!
var (
	constNum1 = big.NewInt(1)
	constNum2 = big.NewInt(2)
	constNum3 = big.NewInt(3)
)

// RoPollardFactor - Ро-метод Полларда.
//
// Вход: Число n, начальное значение c, функция fx (полином)
//
// Выход: Нетривиальный делитель числа p числа n
func RoPollardFactor(_n *big.Int, _c *big.Int, _fx *polynomial.Polynomial) *big.Int {

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

// RoOnePollardFactor - (Po - 1)-метод Полларда.
//
// Вход: Число n
//
// Выход: Нетривиальный делитель числа p числа n
func RoOnePollardFactor(_n *big.Int) *big.Int {

	// Копируем значения чтобы не менять значения по указателю
	n := new(big.Int)
	n.Set(_n)

	// Проверка входных данных
	if n.Sign() <= 0 {
		panic("n > 0")
	}

	if n.Cmp(constNum3) <= 0 {
		return nil
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

// PollardFactorization - Факторизация числа, методами RoPollardFactor и RoOnePollardFactor.
//
// Вход: Число n, steps - количество повторений каждого Ro-метода
//
// Выход: Массив различных делителей числа n
func PollardFactorization(_n *big.Int, steps int) (result []*big.Int) {

	// Копируем значения, что бы не ихменять их по указателю
	n := new(big.Int)
	n.Set(_n)

	result = append(result, new(big.Int).Set(n))

	factor := new(big.Int)

	polyArr := []*big.Int{
		big.NewInt(1),
		big.NewInt(0),
		big.NewInt(1),
	}

	poly := polynomial.NewPolynomial(polyArr)

	for s := 0; s < steps; s++ {

		for i := 0; i < len(result); i++ {

			if result[i].Cmp(constNum3) <= 0 {
				continue
			}

			// Генерируем случайное число c
			c, err := rand.Int(
				rand.Reader,
				new(big.Int).Sub(result[i], constNum2),
			)

			if err != nil {
				panic(err)
			}
			c = c.Add(c, constNum1)
			//

			factor = RoPollardFactor(result[i], c, poly)

			if factor != nil {
				result[i].Set(new(big.Int).Div(result[i], factor))

				result = append(result, new(big.Int).Set(factor))
			}

			// Проверка повторяющихся делителей

			for k := 0; k < 0; k++ {
				for new(big.Int).Mod(result[k], factor).Sign() == 0 {
					result[k].Div(result[k], factor)
				}
			}
		}
	}

	for s := 0; s < steps; s++ {

		for i := 0; i < len(result); i++ {

			if result[i].Cmp(constNum3) <= 0 {
				continue
			}

			factor = RoOnePollardFactor(result[i])

			if factor != nil {
				result[i].Set(new(big.Int).Div(result[i], factor))

				result = append(result, new(big.Int).Set(factor))
			}

			// Проверка повторяющихся делителей

			for k := 0; k < 0; k++ {
				for new(big.Int).Mod(result[k], factor).Sign() == 0 {
					result[k].Div(result[k], factor)
				}
			}
		}
	}

	// Сортируем массив
	sort.Slice(result, func(i, j int) bool {
		return result[i].Cmp(result[j]) < 0
	})

	// TODO: Убрать повторы

	return result
}

// BruteForceFactorization - Факторизация числа перебором.
//
// Вход: Число n
//
// Выход: Массив различных делителей числа n
func BruteForceFactorization(_n *big.Int) (result []*big.Int) {

	// Копируем значения, что бы не ихменять их по указателю
	n := new(big.Int)
	n.Set(_n)

	sqrtFloatN := new(big.Float)
	sqrtFloatN.SetInt(n)

	sqrtFloatN = bigFloat.Sqrt(sqrtFloatN)

	sqrtIntN := new(big.Int)
	sqrtIntN, _ = sqrtFloatN.Int(sqrtIntN)
	sqrtIntN = sqrtIntN.Add(sqrtIntN, constNum1)

	// Цикл до корня из n
	for factor := big.NewInt(2); factor.Cmp(sqrtIntN) <= 0; factor = factor.Add(factor, constNum1) {

		if new(big.Int).Mod(n, factor).Sign() == 0 {
			result = append(result, new(big.Int).Set(factor))
		}

		for new(big.Int).Mod(n, factor).Sign() == 0 {
			n.Div(n, factor)
		}
	}

	return result
}
