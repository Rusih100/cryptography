package finite_field

import (
	"cryptography/crypto_math"
	"github.com/Rusih100/polynomial"
	"math/big"
	"os"
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

	if !crypto_math.MillerRabinTest(p) {
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
	want = crypto_math.Pow(g.p, g.n)

	if want.Cmp(counter) != 0 {
		panic("The field cannot be created")
	}

	return g
}

// Add - Складывает два элемента в поле
func (g *GaloisField) Add(a, b *polynomial.Polynomial) *polynomial.Polynomial {

	maxValue := new(big.Int)
	maxValue = crypto_math.Pow(g.p, g.n)

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
	maxValue = crypto_math.Pow(g.p, g.n)

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

// CayleyTableAdd - Таблица Кэли для сложения
//
// Файл сохраняется в finite_field/cayley_table
func (g *GaloisField) CayleyTableAdd() {

	// Максимальное количество элементов
	maxValue := new(big.Int)
	maxValue = crypto_math.Pow(g.p, g.n)

	name := g.p.String() + "^" + g.n.String() + "_add"

	// Создание файла
	file, err := os.Create("finite_field/cayley_table/" + name + ".csv")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	// Создаем массив для итераций по многочленам
	var iArr []*big.Int
	for i := big.NewInt(0); i.Cmp(g.n) <= 0; i.Add(i, constNum1) {
		iArr = append(iArr, big.NewInt(0))
	}

	// Первая строка
	result := "\t"

	temp := new(polynomial.Polynomial)

	for O := big.NewInt(1); O.Cmp(maxValue) <= 0; O.Add(O, constNum1) {

		temp.Set(iArr)
		result = result + temp.String()

		if O.Cmp(maxValue) != 0 {
			result = result + "\t"
		}

		iArr[0] = iArr[0].Add(iArr[0], constNum1)

		for i := 0; i < len(iArr); i++ {

			if iArr[i].Cmp(g.p) == 0 {
				iArr[i].Mod(iArr[i], g.p)
				iArr[i+1].Add(iArr[i+1], constNum1)
			}
		}

	}
	result = result + "\n"

	_, err = file.WriteString(result)
	if err != nil {
		panic(err)
	}

	// Массив строк

	// Создаем массивы для итерации
	iArr = []*big.Int{}
	for i := big.NewInt(0); i.Cmp(g.n) <= 0; i.Add(i, constNum1) {
		iArr = append(iArr, big.NewInt(0))
	}

	a := new(polynomial.Polynomial)
	b := new(polynomial.Polynomial)

	for I := big.NewInt(1); I.Cmp(maxValue) <= 0; I.Add(I, constNum1) {

		jArr := []*big.Int{}
		for j := big.NewInt(0); j.Cmp(g.n) <= 0; j.Add(j, constNum1) {
			jArr = append(jArr, big.NewInt(0))
		}

		a.Set(iArr)
		result = a.String() + "\t"

		for J := big.NewInt(1); J.Cmp(maxValue) <= 0; J.Add(J, constNum1) {

			b.Set(jArr)

			temp = temp.Add(a, b)
			temp = temp.Mod(temp, g.p)

			result = result + temp.String()

			if J.Cmp(maxValue) != 0 {
				result = result + "\t"
			}

			jArr[0] = jArr[0].Add(jArr[0], constNum1)

			for j := 0; j < len(jArr); j++ {

				if jArr[j].Cmp(g.p) == 0 {
					jArr[j].Mod(jArr[j], g.p)
					jArr[j+1].Add(jArr[j+1], constNum1)
				}
			}

		}

		result = result + "\n"

		_, err = file.WriteString(result)
		if err != nil {
			panic(err)
		}

		// Увеличение полинома a

		if I.Cmp(new(big.Int).Sub(maxValue, constNum1)) != 0 {
			result = result + "\t"
		}

		iArr[0] = iArr[0].Add(iArr[0], constNum1)

		for i := 0; i < len(iArr); i++ {

			if iArr[i].Cmp(g.p) == 0 {
				iArr[i].Mod(iArr[i], g.p)
				iArr[i+1].Add(iArr[i+1], constNum1)
			}
		}

	}

}

// CayleyTableMul - Таблица Кэли для умножения
//
// Файл сохраняется в finite_field/cayley_table
func (g *GaloisField) CayleyTableMul() {

	// Максимальное количество элементов
	maxValue := new(big.Int)
	maxValue = crypto_math.Pow(g.p, g.n)

	name := g.p.String() + "^" + g.n.String() + "_mul"

	// Создание файла
	file, err := os.Create("finite_field/cayley_table/" + name + ".csv")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	// Создаем массив для итераций по многочленам
	var iArr []*big.Int
	for i := big.NewInt(0); i.Cmp(g.n) <= 0; i.Add(i, constNum1) {
		iArr = append(iArr, big.NewInt(0))
	}

	// Первая строка
	result := "\t"

	temp := new(polynomial.Polynomial)

	for O := big.NewInt(1); O.Cmp(maxValue) <= 0; O.Add(O, constNum1) {

		temp.Set(iArr)
		result = result + temp.String()

		if O.Cmp(maxValue) != 0 {
			result = result + "\t"
		}

		iArr[0] = iArr[0].Add(iArr[0], constNum1)

		for i := 0; i < len(iArr); i++ {

			if iArr[i].Cmp(g.p) == 0 {
				iArr[i].Mod(iArr[i], g.p)
				iArr[i+1].Add(iArr[i+1], constNum1)
			}
		}

	}
	result = result + "\n"

	_, err = file.WriteString(result)
	if err != nil {
		panic(err)
	}

	// Массив строк

	// Создаем массивы для итерации
	iArr = []*big.Int{}
	for i := big.NewInt(0); i.Cmp(g.n) <= 0; i.Add(i, constNum1) {
		iArr = append(iArr, big.NewInt(0))
	}

	a := new(polynomial.Polynomial)
	b := new(polynomial.Polynomial)

	for I := big.NewInt(1); I.Cmp(maxValue) <= 0; I.Add(I, constNum1) {

		jArr := []*big.Int{}
		for j := big.NewInt(0); j.Cmp(g.n) <= 0; j.Add(j, constNum1) {
			jArr = append(jArr, big.NewInt(0))
		}

		a.Set(iArr)
		result = a.String() + "\t"

		for J := big.NewInt(1); J.Cmp(maxValue) <= 0; J.Add(J, constNum1) {

			b.Set(jArr)

			temp = temp.Mul(a, b)
			temp = temp.Mod(temp, g.p)
			_, temp = temp.QuoRem(temp, g.mod)
			temp = temp.Mod(temp, g.p)

			result = result + temp.String()

			if J.Cmp(maxValue) != 0 {
				result = result + "\t"
			}

			jArr[0] = jArr[0].Add(jArr[0], constNum1)

			for j := 0; j < len(jArr); j++ {

				if jArr[j].Cmp(g.p) == 0 {
					jArr[j].Mod(jArr[j], g.p)
					jArr[j+1].Add(jArr[j+1], constNum1)
				}
			}

		}

		result = result + "\n"

		_, err = file.WriteString(result)
		if err != nil {
			panic(err)
		}

		// Увеличение полинома a

		if I.Cmp(new(big.Int).Sub(maxValue, constNum1)) != 0 {
			result = result + "\t"
		}

		iArr[0] = iArr[0].Add(iArr[0], constNum1)

		for i := 0; i < len(iArr); i++ {

			if iArr[i].Cmp(g.p) == 0 {
				iArr[i].Mod(iArr[i], g.p)
				iArr[i+1].Add(iArr[i+1], constNum1)
			}
		}

	}

}

// NewGaloisField - Создает GaloisField и задает ему начальные значения
func NewGaloisField(p *big.Int, n *big.Int, poly *polynomial.Polynomial) *GaloisField {
	return new(GaloisField).Set(p, n, poly)
}
