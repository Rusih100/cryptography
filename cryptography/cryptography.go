package cryptography

import (
	"container/list"
	"crypto/rand"
	"math/big"
	rnd "math/rand"
	"time"
)

// 1. - OK (Сдано)

// EuclidAlgorithm - обобщенный (расширенный) алгоритм Евклида.
//
// Вход: положительные числа x и y отличные от нуля.
//
// Выход: m, a, b - модуль и его линейное представление.
func EuclidAlgorithm(_x *big.Int, _y *big.Int) (m, a, b *big.Int) {

	// Копируем значения, чтобы не менять значения по указателю
	x := new(big.Int)
	y := new(big.Int)

	x.Set(_x)
	y.Set(_y)

	flagSwap := false

	if x.Sign() <= 0 || y.Sign() <= 0 {
		panic("x and y must be positive numbers other than zero")
	}

	if x.Cmp(y) == -1 {
		*x, *y = *y, *x
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
		*a, *b = *b, *a
	}

	return m, a, b
}

// 2. - OK (Сдано)

// Pow - Алгоритм быстрого возведения в степень.
//
// Вход: a - основание (число), n - положительная степень (число).
//
// Выход: result - число a**n.
func Pow(_a *big.Int, _n *big.Int) (result *big.Int) {

	// Копируем значения, чтобы не менять значения по указателю
	a := new(big.Int)
	n := new(big.Int)

	a.Set(_a)
	n.Set(_n)

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

	if n.Sign() == -1 {
		panic("n must be greater than or equal to zero")
	}
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

	if n.Cmp(big.NewInt(3)) == -1 {
		panic("n must be greater than or equal to 3")
	}

	if a.Sign() == -1 || a.Cmp(n) >= 0 {
		panic("a: 0 <= a < n")
	}

	// 1. Проверка взаимной простоты
	gcd := new(big.Int)
	gcd, _, _ = EuclidAlgorithm(a, n)

	if gcd.Cmp(big.NewInt(1)) != 0 {
		return 0
	}

	// 2. Инициализация
	var result int64 = 1

	for {
		// 3. Избавление от четности
		k := big.NewInt(0)
		for a.Bit(0) == 0 {
			k = k.Add(k, big.NewInt(1))
			a = a.Div(a, big.NewInt(2))
		}

		if k.Bit(0) == 1 &&
			(new(big.Int).Mod(n, big.NewInt(8)).Cmp(big.NewInt(3)) == 0 || new(big.Int).Mod(n, big.NewInt(8)).Cmp(big.NewInt(5)) == 0) {
			result = -result
		}

		// 4. Квадратичный закон взамности
		if new(big.Int).Mod(a, big.NewInt(4)).Cmp(big.NewInt(3)) == 0 &&
			new(big.Int).Mod(n, big.NewInt(4)).Cmp(big.NewInt(3)) == 0 {
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

	if n.Bit(0) == 0 && n.Cmp(big.NewInt(2)) != 0 {
		return false
	}

	if n.Cmp(big.NewInt(1)) == 1 && n.Cmp(big.NewInt(5)) == -1 {
		return true
	}

	if n.Cmp(big.NewInt(1)) <= 0 {
		panic("n > 1")
	}

	// Генерируем случайное число 2 ≤ a < n - 1:
	// 0 ≤ a < n - 3	| +2
	a, err := rand.Int(
		rand.Reader,
		new(big.Int).Sub(n, big.NewInt(3)),
	)

	if err != nil {
		panic(err)
	}

	a = a.Add(a, big.NewInt(2))

	//

	r := new(big.Int)
	r = PowMod(
		a,
		new(big.Int).Sub(n, big.NewInt(1)),
		n,
	)
	if r.Bit(0) == 1 {
		return true
	} else {
		return false
	}
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

	if n.Bit(0) == 0 && n.Cmp(big.NewInt(2)) != 0 {
		return false
	}

	if n.Cmp(big.NewInt(1)) == 1 && n.Cmp(big.NewInt(5)) == -1 {
		return true
	}

	if n.Cmp(big.NewInt(1)) <= 0 {
		panic("n > 1")
	}

	// Генерируем случайное число 2 ≤ a < n - 1:
	// 0 ≤ a < n - 3	| +2
	a, err := rand.Int(
		rand.Reader,
		new(big.Int).Sub(n, big.NewInt(3)),
	)

	if err != nil {
		panic(err)
	}

	a = a.Add(a, big.NewInt(2))

	//

	r := new(big.Int)

	r = PowMod(
		a,
		new(big.Int).Div(
			new(big.Int).Sub(n, big.NewInt(1)),
			big.NewInt(2)),
		n,
	)

	if r.Cmp(big.NewInt(1)) != 0 && r.Cmp(new(big.Int).Sub(n, big.NewInt(1))) != 0 {
		return false
	}

	s := Jacobi(a, n)

	if new(big.Int).Mod(new(big.Int).Sub(r, big.NewInt(s)), n).BitLen() == 0 {
		return true
	} else {
		return false
	}
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

	if n.Bit(0) == 0 && n.Cmp(big.NewInt(2)) != 0 {
		return false
	}

	if n.Cmp(big.NewInt(1)) == 1 && n.Cmp(big.NewInt(5)) == -1 {
		return true
	}

	if n.Cmp(big.NewInt(1)) <= 0 {
		panic("n > 1")
	}

	// n - 1
	r := new(big.Int).Sub(n, big.NewInt(1))

	s := big.NewInt(0)
	for r.Bit(0) == 0 {
		s = s.Add(s, big.NewInt(1))
		r = r.Div(r, big.NewInt(2))
	}

	// Генерируем случайное число 2 ≤ a < n - 1:
	// 0 ≤ a < n - 3	| +2
	a, err := rand.Int(
		rand.Reader,
		new(big.Int).Sub(n, big.NewInt(3)),
	)

	if err != nil {
		panic(err)
	}

	a = a.Add(a, big.NewInt(2))

	//

	y := new(big.Int)
	y = PowMod(a, r, n)

	if y.Cmp(big.NewInt(1)) != 0 && y.Cmp(new(big.Int).Sub(n, big.NewInt(1))) != 0 {
		j := big.NewInt(1)

		if new(big.Int).Sub(s, big.NewInt(1)).Cmp(j) >= 0 && y.Cmp(new(big.Int).Sub(n, big.NewInt(1))) != 0 {
			y = PowMod(y, big.NewInt(2), n)
			if y.Cmp(big.NewInt(1)) == 0 {
				return false
			}
			j = j.Add(j, big.NewInt(1))
		}

		if y.Cmp(new(big.Int).Sub(n, big.NewInt(1))) != 0 {
			return false
		}
	}
	return true
}

// 8. - OK

// RandNumber - Генерация k-битного случайного нечетного числа.
//
// Вход: Разрядность k генерируемого числа.
//
// Выход: Случайное k-битное нечетное число.
func RandNumber(k uint) (result *big.Int) {

	// Проверка входных данных
	if k <= 1 {
		panic("k > 1")
	}

	randNumber := big.NewInt(1)

	// Получаем k битное число из 1: 1000...001
	randNumber = randNumber.Lsh(randNumber, k-1)
	randNumber = randNumber.Add(randNumber, big.NewInt(1))

	// Случайные числа
	rnd.Seed(time.Now().UnixNano())

	shift := big.NewInt(2) // Для битового сдвига

	// Побитовая догенерация случайного числа с помощью OR
	for i := 1; i < randNumber.BitLen()-1; i++ {
		bit := rnd.Int31n(2)
		if bit == 1 {
			randNumber = randNumber.Or(randNumber, shift)
		}
		shift = shift.Lsh(shift, 1)
	}
	return randNumber
}

// SimpleNumber - Генерация k-битного простого числа.
//
// Вход: Разрядность k искомого простого числа, параметр t >= 1.
//
// Выход: Число, простое с вероятностью 1 - 1 / (4**t).
func SimpleNumber(k uint, t uint) (result *big.Int) {

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

	for i := uint(0); i < t; i++ {
		if !MillerRabinTest(randNumber) {
			randNumber = RandNumber(k)
			i = 0
		}
	}
	return randNumber
}

// 9. - OK

// InverseElement - Нахождение обратного элемента по модулю через расширенный алгоритм Евклида
//
// Вход: a > 0, mod > 0
//
// Выход: Обратный элемент к a по модулю mod
func InverseElement(_a *big.Int, _mod *big.Int) (result *big.Int) {

	// Копируем значения, чтобы не менять значения по указателю
	a := new(big.Int)
	mod := new(big.Int)

	a.Set(_a)
	mod.Set(_mod)

	// Проверка входных данных
	if a.Sign() <= 0 {
		panic("a > 0")
	}

	if mod.Sign() <= 0 {
		panic("mod > 0")
	}

	_, _, result = EuclidAlgorithm(mod, a)

	return result
}

// ModuloComparisonFirst - Решение сравнения первой степени.
//
// Вход: Сравнение вида ax = b (mod): числа a, b и mod. (a > 0, mod > 0)
//
// Выход: Список, содержащий все решения данного сравнение, если оно разрешимо,
// иначе возвращается пустой список
//
// Примечание: Количество решений не может превышать размерности int64
func ModuloComparisonFirst(_a *big.Int, _b *big.Int, _mod *big.Int) (result list.List) {

	// Копируем значения, чтобы не менять значения по указателю
	a := new(big.Int)
	b := new(big.Int)
	mod := new(big.Int)

	a.Set(_a)
	b.Set(_b)
	mod.Set(_mod)

	// Проверка входных данных
	if a.Sign() <= 0 {
		panic("a > 0")
	}

	if mod.Sign() <= 0 {
		panic("mod > 0")
	}

	// Проверяем разрешимость сравнения
	gcd := new(big.Int)
	gcd, _, _ = EuclidAlgorithm(a, mod)

	// Если неразрешимо возвращаем пустой массив
	if new(big.Int).Mod(b, gcd).Sign() != 0 {
		return result
	}

	// Единственное решение
	if gcd.Cmp(big.NewInt(1)) == 0 {

		x := new(big.Int)
		// Записываем в x обратный к а элемент, далее умножаем на b
		x = InverseElement(a, mod)
		x = x.Mod(new(big.Int).Mul(x, b), mod)

		result.PushBack(new(big.Int).Set(x))
		return result
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

	for i := big.NewInt(0); i.Cmp(x) == -1; i = i.Add(i, big.NewInt(1)) {
		result.PushBack(new(big.Int).Set(x))
		x = x.Add(x, mod1)
	}
	return result
}
