package crypto_math

import (
	"crypto/rand"
	"math/big"
	rnd "math/rand"
	"time"
)

// Константы для упрощения кода
// Не изменять в алгоритмах!
var (
	constNum1 = big.NewInt(1)
	constNum2 = big.NewInt(2)
	constNum3 = big.NewInt(3)
	constNum4 = big.NewInt(4)
	constNum5 = big.NewInt(5)
	constNum8 = big.NewInt(8)
)

// 1. - OK (Сдано)

// AdvancedEuclidAlgorithm - обобщенный (расширенный) алгоритм Евклида.
//
// Вход: натуральные числа x и y отличные от нуля.
//
// Выход: m, a, b - наибольший общий делитель и его линейное представление.
func AdvancedEuclidAlgorithm(_x *big.Int, _y *big.Int) (m, a, b *big.Int) {

	// Копируем значения, чтобы не менять значения по указателю
	x := new(big.Int)
	y := new(big.Int)

	x.Set(_x)
	y.Set(_y)

	flagSwap := false

	// x <= 0 или y <= 0
	if x.Sign() <= 0 || y.Sign() <= 0 {
		panic("x and y must be positive numbers other than zero")
	}

	// x < y
	if x.Cmp(y) == -1 {
		x, y = y, x
		flagSwap = true
	}

	a2 := big.NewInt(1)
	a1 := big.NewInt(0)
	b2 := big.NewInt(0)
	b1 := big.NewInt(1)

	for y.BitLen() > 0 {

		// q = x / y
		q := new(big.Int).Div(x, y)

		// r = x - q * y
		r := new(big.Int).Sub(
			x,
			new(big.Int).Mul(q, y),
		)

		// a = a2 - q * a1
		a = new(big.Int).Sub(
			a2,
			new(big.Int).Mul(q, a1),
		)

		// b = b2 - q * b1
		b = new(big.Int).Sub(
			b2,
			new(big.Int).Mul(q, b1),
		)

		x = y
		y = r
		a2 = a1
		a1 = a
		b2 = b1
		b1 = b

		m = x
		a = a2
		b = b2
	}
	if flagSwap {
		a, b = b, a
	}

	return m, a, b
}

// EuclidAlgorithm - алгоритм Евклида для целых чисел.
//
// Вход: целые числа x и y.
//
// Выход: m - наибольший общий делитель.
func EuclidAlgorithm(_x *big.Int, _y *big.Int) *big.Int {

	// Копируем значения, чтобы не менять значения по указателю
	x := new(big.Int)
	y := new(big.Int)

	x.Set(_x)
	y.Set(_y)

	// x < y
	if x.Cmp(y) == -1 {
		x, y = y, x
	}

	// Если числа отрицательные
	if x.Sign() < 0 {
		x = x.Neg(x)
	}

	if y.Sign() < 0 {
		y = y.Neg(y)
	}

	r := big.NewInt(0)

	for y.Sign() != 0 {
		r = r.Mod(x, y)
		x.Set(y)
		y.Set(r)
	}
	return x
}

// 2. - OK (Сдано)

// Pow - Алгоритм быстрого возведения в степень.
//
// Вход: a - основание (число), n - положительная степень (число).
//
// Выход: result - число a^n.
func Pow(_a *big.Int, _n *big.Int) (result *big.Int) {

	// Копируем значения, чтобы не менять значения по указателю
	a := new(big.Int)
	n := new(big.Int)

	a.Set(_a)
	n.Set(_n)

	// n < 0
	if n.Sign() == -1 {
		panic("n must be greater than or equal to zero")
	}

	result = big.NewInt(1)

	for i := 0; i < n.BitLen(); i++ {
		if n.Bit(i) == 1 {
			result = result.Mul(result, a)
		}
		a = a.Mul(a, a)
	}
	return result
}

// 3. - OK (Сдано)

// PowMod - Алгоритм быстрого возведения в степень по модулю.
//
// Вход: a - основание (число), n - положительная степень (число),
// mod - модуль (положительное число отличное от нуля).
//
// Выход: result - число a^n по модулю mod.
func PowMod(_a *big.Int, _n *big.Int, _mod *big.Int) (result *big.Int) {

	// Копируем значения, чтобы не менять значения по указателю
	a := new(big.Int)
	n := new(big.Int)
	mod := new(big.Int)

	a.Set(_a)
	n.Set(_n)
	mod.Set(_mod)

	// n < 0
	if n.Sign() < 0 {
		panic("n must be greater than or equal to zero")
	}

	// mod <= 0
	if mod.Sign() <= 0 {
		panic("mod must be a positive number other than zero")
	}

	result = big.NewInt(1)

	for i := 0; i < n.BitLen(); i++ {
		if n.Bit(i) == 1 {
			result = result.Mod(
				result.Mul(result, a),
				mod,
			)
		}
		a = a.Mod(
			a.Mul(a, a),
			mod,
		)
	}
	return result
}

// 4. - OK (Сдано)

// Jacobi - Алгоритм вычисления символа Якоби (Алгоритм взят с Википедии).
//
// Вход: a (a: 0 <= a < n) , n - натуральное нечетное больше 1 (n >= 3).
//
// Выход: Символ Якоби - 0, 1 или -1.
func Jacobi(_a *big.Int, _n *big.Int) int64 {

	// Копируем значения, чтобы не менять значения по указателю
	a := new(big.Int)
	n := new(big.Int)

	a.Set(_a)
	n.Set(_n)

	// Проверка входных данных
	if n.Bit(0) == 0 {
		panic("n must be odd")
	}

	// n < 3
	if n.Cmp(constNum3) < 0 {
		panic("n must be greater than or equal to 3")
	}

	// a < 0 или a >= n
	if a.Sign() < 0 || a.Cmp(n) >= 0 {
		panic("a: 0 <= a < n")
	}

	// a == 0
	if a.Sign() == 0 {
		return 0
	}

	// 1. Проверка взаимной простоты
	gcd := new(big.Int)
	gcd = EuclidAlgorithm(a, n)

	// gcd != 1
	if gcd.Cmp(constNum1) != 0 {
		return 0
	}

	// 2. Инициализация
	var result int64 = 1

	for {
		// 3. Избавление от четности
		k := big.NewInt(0)
		for a.Bit(0) == 0 {
			k = k.Add(k, constNum1)
			a = a.Rsh(a, 1)
		}

		// k - нечетное и (n (mod 8) = 3 или n (mod 8) = 5)
		if k.Bit(0) == 1 &&
			(new(big.Int).Mod(n, constNum8).Cmp(constNum3) == 0 ||
				new(big.Int).Mod(n, constNum8).Cmp(constNum5) == 0) {
			result = -result
		}

		// 4. Квадратичный закон взамности

		// a (mod 4) = 3 и n (mod 4) = 3
		if new(big.Int).Mod(a, constNum4).Cmp(constNum3) == 0 &&
			new(big.Int).Mod(n, constNum4).Cmp(constNum3) == 0 {
			result = -result
		}
		c := new(big.Int)
		c.Set(a)
		a = a.Mod(n, c)
		n.Set(c)

		// 5. Выход из алгоритма?
		if a.BitLen() == 0 {
			return result
		}
	}
}

// 5. - OK (Сдано)

// FermatTest - Тест Ферма.
//
// Вход: n - целое число, n > 1.
//
// Выход: true - "Число n, вероятно, простое" или false - "Число n составное".
func FermatTest(_n *big.Int) bool {

	// Копируем значения, чтобы не менять значения по указателю
	n := new(big.Int)
	n.Set(_n)

	if n.Bit(0) == 0 && n.Cmp(constNum2) != 0 {
		return false
	}

	// n > 0 и n < 5
	if n.Cmp(constNum1) > 0 && n.Cmp(constNum5) < 0 {
		return true
	}

	// n <= 1
	if n.Cmp(constNum1) <= 0 {
		panic("n > 1")
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

	r := new(big.Int)
	r = PowMod(
		a,
		new(big.Int).Sub(n, constNum1),
		n,
	)
	if r.Cmp(constNum1) == 0 {
		return true
	}
	return false
}

// 6. - OK (Сдано)

// SolovayStrassenTest - Тест Соловэя-Штрассена.
//
// Вход: n - целое число, n > 1.
//
// Выход: true - "Число n, вероятно, простое" или false - "Число n составное".
func SolovayStrassenTest(_n *big.Int) bool {

	// Копируем значения, чтобы не менять значения по указателю
	n := new(big.Int)
	n.Set(_n)

	if n.Bit(0) == 0 && n.Cmp(constNum2) != 0 {
		return false
	}

	// n > 0 и n < 5
	if n.Cmp(constNum1) > 0 && n.Cmp(constNum5) < 0 {
		return true
	}

	// n <= 1
	if n.Cmp(constNum1) <= 0 {
		panic("n > 1")
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

	r := new(big.Int)

	r = PowMod(
		a,
		new(big.Int).Div(
			new(big.Int).Sub(n, constNum1),
			constNum2),
		n,
	)

	// r != 1 и r != n - 1
	if r.Cmp(constNum1) != 0 && r.Cmp(new(big.Int).Sub(n, constNum1)) != 0 {
		return false
	}

	s := Jacobi(a, n)

	// (r - s) (mod n) == 0
	if new(big.Int).Mod(new(big.Int).Sub(r, big.NewInt(s)), n).Sign() == 0 {
		return true
	}
	return false
}

// 7. - OK (Сдано)

// MillerRabinTest - Тест Миллера-Рабина.
//
// Вход: n - целое число, n > 1.
//
// Выход: true - "Число n, вероятно, простое" или false - "Число n составное".
func MillerRabinTest(_n *big.Int) bool {

	// Копируем значения, чтобы не менять значения по указателю
	n := new(big.Int)
	n.Set(_n)

	if n.Bit(0) == 0 && n.Cmp(constNum2) != 0 {
		return false
	}

	// n > 0 и n < 5
	if n.Cmp(constNum1) > 0 && n.Cmp(constNum5) < 0 {
		return true
	}

	// n <= 1
	if n.Cmp(constNum1) <= 0 {
		panic("n > 1")
	}

	// n - 1
	r := new(big.Int).Sub(n, constNum1)

	s := big.NewInt(0)
	for r.Bit(0) == 0 {
		s = s.Add(s, constNum1)
		r = r.Rsh(r, 1)
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

	y := new(big.Int)
	y = PowMod(a, r, n)

	// y != 1 и y != n - 1
	if y.Cmp(constNum1) != 0 && y.Cmp(new(big.Int).Sub(n, constNum1)) != 0 {
		j := big.NewInt(1)

		// s - 1 >= j и y != n -1
		for new(big.Int).Sub(s, constNum1).Cmp(j) >= 0 && y.Cmp(new(big.Int).Sub(n, constNum1)) != 0 {
			y = PowMod(y, constNum2, n)
			// y == 1
			if y.Cmp(constNum1) == 0 {
				return false
			}
			j = j.Add(j, constNum1)
		}

		// y != n - 1
		if y.Cmp(new(big.Int).Sub(n, constNum1)) != 0 {
			return false
		}
	}
	return true
}

// 8. - OK (Сдано)

// RandNumber - Генерация k-битного случайного нечетного числа.
//
// Вход: Разрядность k генерируемого числа.
//
// Выход: Случайное k-битное нечетное число.
func RandNumber(k int) (result *big.Int) {

	// Проверка входных данных
	if k <= 1 {
		panic("k > 1")
	}

	randNumber := new(big.Int)

	// Получаем k битное число из 1: 1000...001
	randNumber = randNumber.SetBit(randNumber, k-1, 1)
	randNumber = randNumber.SetBit(randNumber, 0, 1)

	// Случайные числа
	rnd.Seed(time.Now().UnixNano())

	// Побитовая догенерация случайного числа с помощью OR
	for i := 1; i < randNumber.BitLen()-1; i++ {
		bit := rnd.Int31n(2)
		if bit == 1 {
			randNumber = randNumber.SetBit(randNumber, i, uint(bit))
		}
	}
	return randNumber
}

// SimpleNumber - Генерация k-битного простого числа.
//
// Вход: Разрядность k искомого простого числа, параметр t >= 1.
//
// Выход: Число, простое с вероятностью 1 - 1 / (4**t).
func SimpleNumber(k int, t int) (result *big.Int) {

	// Проверка входных данных
	if k <= 1 {
		panic("k > 1")
	}

	if t < 1 {
		panic("t >= 1")
	}

	// Генерируем случайное нечетное k-битное число
	randNumber := new(big.Int)
	randNumber = RandNumber(k)

	for i := 0; i < t; i++ {
		if !MillerRabinTest(randNumber) {
			randNumber = RandNumber(k)
			i = 0
		}
	}
	return randNumber
}

// 9. - OK (Сдано)

// InverseElement - Нахождение обратного элемента по модулю через расширенный алгоритм Евклида.
//
// Вход: a > 0, mod > 0.
//
// Выход: Обратный элемент к a по модулю mod.
func InverseElement(_a *big.Int, _mod *big.Int) (result *big.Int) {

	// Копируем значения, чтобы не менять значения по указателю
	a := new(big.Int)
	mod := new(big.Int)

	a.Set(_a)
	mod.Set(_mod)

	// Проверка входных данных
	// a <= 0
	if a.Sign() <= 0 {
		panic("a > 0")
	}

	// mod <= 0
	if mod.Sign() <= 0 {
		panic("mod > 0")
	}

	_, _, result = AdvancedEuclidAlgorithm(mod, a)

	return result
}

// ModuloComparisonFirst - Решение сравнения первой степени.
//
// Вход: Сравнение вида ax = b (по модулю mod): числа a, b и mod. (a > 0, mod > 0).
//
// Выход: Количество решений, первое решение, сдвиг для получения следущего решения.
func ModuloComparisonFirst(_a *big.Int, _b *big.Int, _mod *big.Int) (countSolutions, x1, offset *big.Int) {

	// Копируем значения, чтобы не менять значения по указателю
	a := new(big.Int)
	b := new(big.Int)
	mod := new(big.Int)

	a.Set(_a)
	b.Set(_b)
	mod.Set(_mod)

	// Проверка входных данных
	// a == 0
	if a.Sign() == 0 {
		panic("a != 0")
	}

	// mod <= 0
	if mod.Sign() <= 0 {
		panic("mod > 0")
	}

	// Переход к положительным числам
	// a (mod) + mod
	a = new(big.Int).Add(new(big.Int).Mod(a, mod), mod)
	b = new(big.Int).Add(new(big.Int).Mod(b, mod), mod)

	// Проверяем разрешимость сравнения
	gcd := new(big.Int)
	gcd = EuclidAlgorithm(a, mod)

	// Если неразрешимо
	// b (mod gcd) != 0
	if new(big.Int).Mod(b, gcd).Sign() != 0 {
		return big.NewInt(0), nil, nil
	}

	// Единственное решение
	// gcd == 1
	if gcd.Cmp(constNum1) == 0 {

		x := new(big.Int)

		// Записываем в x обратный к а элемент, далее умножаем на b
		x = InverseElement(a, mod)
		x = x.Mod(new(big.Int).Mul(x, b), mod)

		return big.NewInt(1), x, nil
	}

	// Множество решений

	// Переход к новому сравнению
	a1 := new(big.Int).Div(a, gcd)
	b1 := new(big.Int).Div(b, gcd)
	mod1 := new(big.Int).Div(mod, gcd)

	x := new(big.Int)
	// Записываем в x обратный к а элемент, далее умножаем на b
	x = InverseElement(a1, mod1)
	x = x.Mod(new(big.Int).Mul(x, b1), mod1)

	return gcd, x, mod1
}

// 10. - OK (Сдано)

// ModuloComparisonSecond - Решение сравнения второй степени.
//
// Вход: Сравнение вида x^2 = a (по модулю p): числа a и p. (p - простое и p > 2).
//
// Выход: Решение сравнения второй степени.
func ModuloComparisonSecond(_a *big.Int, _p *big.Int) (xPos, xNeg *big.Int) {

	// Копируем значения, чтобы не менять значения по указателю
	a := new(big.Int)
	p := new(big.Int)

	a.Set(_a)
	p.Set(_p)

	// Проверка входных данных
	// p <= 2
	if p.Cmp(constNum2) <= 0 {
		panic("p > 2")
	}

	if !MillerRabinTest(p) {
		panic("p is a prime number")
	}

	// Переход к положительным числам
	a = a.Mod(a, p)

	// a == 0
	if a.Sign() == 0 {
		panic("a is not divisible by p")
	}

	// Проверяем квадратичный вычет a
	if Jacobi(a, p) != 1 {
		return nil, nil
	}

	// Перебором ищем квадратичный невычет N
	N := big.NewInt(1)

	// Пока N < p; N++
	for ; N.Cmp(p) < 0; N = N.Add(N, constNum1) {
		if Jacobi(N, p) == -1 {
			break
		}
	}

	// 1. Представление p в виде p = 2^k * h + 1
	h := new(big.Int).Sub(p, constNum1)

	k := big.NewInt(0)
	for h.Bit(0) == 0 {
		k = k.Add(k, constNum1)
		h = h.Rsh(h, 1)
	}

	// 2.
	a1 := new(big.Int)
	a1 = PowMod(
		a,
		new(big.Int).Div(new(big.Int).Add(h, constNum1), constNum2),
		p,
	)

	a2 := new(big.Int)
	a2 = InverseElement(a, p)

	N1 := new(big.Int)
	N1 = PowMod(N, h, p)

	N2 := big.NewInt(1)

	j := big.NewInt(0)

	// 3.
	// i = 0; i <= k - 2; i++
	for i := big.NewInt(0); i.Cmp(new(big.Int).Sub(k, constNum2)) <= 0; i.Add(i, constNum1) {

		// 3.1
		b := new(big.Int).Mod(
			new(big.Int).Mul(a1, N2),
			p,
		)

		// 3.2
		bPow2 := new(big.Int) // Квадрат b
		bPow2 = PowMod(b, constNum2, p)

		c := new(big.Int).Mod(
			new(big.Int).Mul(a2, bPow2),
			p,
		)

		// 3.3
		dPower := new(big.Int)
		dPower = Pow(
			constNum2,
			new(big.Int).Sub(new(big.Int).Sub(k, constNum2), i),
		)

		d := new(big.Int)
		d = PowMod(c, dPower, p)

		// d == 1
		if d.Cmp(constNum1) == 0 {
			j = big.NewInt(0)
		}

		// d == -1
		if d.Cmp(new(big.Int).Add(p, big.NewInt(-1))) == 0 {
			j = big.NewInt(1)
		}

		// 3.4
		N1Power := new(big.Int)
		N1Power = Pow(constNum2, i)
		N1Power = N1Power.Mul(N1Power, j)

		temp := new(big.Int)
		temp = PowMod(N1, N1Power, p)

		N2 = new(big.Int).Mod(
			new(big.Int).Mul(N2, temp),
			p,
		)
	}

	xPos = new(big.Int).Mod(
		new(big.Int).Mul(a1, N2),
		p,
	)

	xNeg = new(big.Int).Mod(
		new(big.Int).Mul(a1, N2),
		p,
	)

	xNeg.Neg(xNeg)
	xNeg = xNeg.Mod(xNeg, p)

	return xPos, xNeg
}

// 11. - OK (Сдано)

// ModuloComparisonSystem - Решение системы сравнений.
//
// Вход: Массив коэфицентов bArray и массив модулей mArray.
//
// Выход: Решение системы сравнений, если все модули взаимопросты.
func ModuloComparisonSystem(bArray []*big.Int, mArray []*big.Int) (result *big.Int) {

	// Длины массивов
	bArrayLen := len(bArray)
	mArrayLen := len(mArray)

	// Проверка входных данных
	if bArrayLen == 0 {
		panic("bArray: An empty array was passed")
	}

	if mArrayLen == 0 {
		panic("mArray: An empty array was passed")
	}

	if bArrayLen != mArrayLen {
		panic("Arrays of various lengths were transmitted")
	}

	// Проверка взаимопростоты модулей
	testGCD := new(big.Int)
	x := new(big.Int)
	y := new(big.Int)

	for i := 0; i < mArrayLen; i++ {
		for j := i + 1; j < mArrayLen; j++ {
			x = mArray[i]
			y = mArray[j]

			testGCD = EuclidAlgorithm(x, y)

			if testGCD.Cmp(constNum1) != 0 {
				return nil
			}
		}
	}

	// Ищем произведение модулей
	M := big.NewInt(1)

	for i := 0; i < mArrayLen; i++ {
		x = mArray[i]
		M = M.Mul(M, x)
	}

	// Ищем решение
	result = big.NewInt(0)
	Mj := new(big.Int)
	Nj := new(big.Int)

	for j := 0; j < mArrayLen; j++ {
		bj := bArray[j]
		mj := mArray[j]

		Mj = Mj.Div(M, mj)
		Nj = InverseElement(Mj, mj)

		sub := new(big.Int).Mul(bj, Mj)
		sub = sub.Mod(sub, M)
		sub = sub.Mul(sub, Nj)
		sub = sub.Mod(sub, M)

		result.Add(result, sub)
		result.Mod(result, M)
	}

	return result
}
