package cryptography

import (
	"math/big"
	"testing"
)

// Тесты только для детерменированных алгоритмов

// Алгоритм Евклида

func TestEuclidAlgorithm1(t *testing.T) {
	x := big.NewInt(13)
	y := big.NewInt(17)

	want := big.NewInt(1)

	//

	gcd := new(big.Int)
	a := new(big.Int)
	b := new(big.Int)

	gcd, a, b = EuclidAlgorithm(x, y)

	subTest := new(big.Int).Add(
		new(big.Int).Mul(x, a),
		new(big.Int).Mul(y, b),
	)

	if want.Cmp(gcd) != 0 || subTest.Cmp(gcd) != 0 {
		t.Fatal()
	}
}

func TestEuclidAlgorithm2(t *testing.T) {
	x := big.NewInt(17)
	y := big.NewInt(13)

	want := big.NewInt(1)

	//

	gcd := new(big.Int)
	a := new(big.Int)
	b := new(big.Int)

	gcd, a, b = EuclidAlgorithm(x, y)

	subTest := new(big.Int).Add(
		new(big.Int).Mul(x, a),
		new(big.Int).Mul(y, b),
	)

	if want.Cmp(gcd) != 0 || subTest.Cmp(gcd) != 0 {
		t.Fatal()
	}
}

func TestEuclidAlgorithm3(t *testing.T) {
	x := big.NewInt(81)
	y := big.NewInt(27)

	want := big.NewInt(27)

	//

	gcd := new(big.Int)
	a := new(big.Int)
	b := new(big.Int)

	gcd, a, b = EuclidAlgorithm(x, y)

	subTest := new(big.Int).Add(
		new(big.Int).Mul(x, a),
		new(big.Int).Mul(y, b),
	)

	if want.Cmp(gcd) != 0 || subTest.Cmp(gcd) != 0 {
		t.Fatal()
	}
}

func TestEuclidAlgorithm4(t *testing.T) {
	x := big.NewInt(1)
	y := big.NewInt(1)

	want := big.NewInt(1)

	//

	gcd := new(big.Int)
	a := new(big.Int)
	b := new(big.Int)

	gcd, a, b = EuclidAlgorithm(x, y)

	subTest := new(big.Int).Add(
		new(big.Int).Mul(x, a),
		new(big.Int).Mul(y, b),
	)

	if want.Cmp(gcd) != 0 || subTest.Cmp(gcd) != 0 {
		t.Fatal()
	}
}

// Возведение в степень

func TestPow1(t *testing.T) {
	a := big.NewInt(2)
	n := big.NewInt(6)

	want := big.NewInt(64)

	//

	res := new(big.Int)
	res = Pow(a, n)

	if want.Cmp(res) != 0 {
		t.Fatal()
	}
}

func TestPow2(t *testing.T) {
	a := big.NewInt(1)
	n := big.NewInt(1024)

	want := big.NewInt(1)

	//

	res := new(big.Int)
	res = Pow(a, n)

	if want.Cmp(res) != 0 {
		t.Fatal()
	}
}

func TestPow3(t *testing.T) {
	a := big.NewInt(1)
	n := big.NewInt(0)

	want := big.NewInt(1)

	//

	res := new(big.Int)
	res = Pow(a, n)

	if want.Cmp(res) != 0 {
		t.Fatal()
	}
}

func TestPow4(t *testing.T) {
	a := big.NewInt(256)
	n := big.NewInt(0)

	want := big.NewInt(1)

	//

	res := new(big.Int)
	res = Pow(a, n)

	if want.Cmp(res) != 0 {
		t.Fatal()
	}
}

func TestPow5(t *testing.T) {
	a := big.NewInt(3)
	n := big.NewInt(4)

	want := big.NewInt(81)

	//

	res := new(big.Int)
	res = Pow(a, n)

	if want.Cmp(res) != 0 {
		t.Fatal()
	}
}

func TestPow6(t *testing.T) {
	a := big.NewInt(27)
	n := big.NewInt(2)

	want := big.NewInt(729)

	//

	res := new(big.Int)
	res = Pow(a, n)

	if want.Cmp(res) != 0 {
		t.Fatal()
	}
}

// Возведение в степень по модулю

func TestPowMod1(t *testing.T) {
	a := big.NewInt(3)
	n := big.NewInt(3)
	mod := big.NewInt(10)

	want := big.NewInt(7)

	//

	res := new(big.Int)
	res = PowMod(a, n, mod)

	if want.Cmp(res) != 0 {
		t.Fatal()
	}
}

func TestPowMod2(t *testing.T) {
	a := big.NewInt(2)
	n := big.NewInt(64)
	mod := big.NewInt(2)

	want := big.NewInt(0)

	//

	res := new(big.Int)
	res = PowMod(a, n, mod)

	if want.Cmp(res) != 0 {
		t.Fatal()
	}
}

func TestPowMod3(t *testing.T) {
	a := big.NewInt(1023)
	n := big.NewInt(0)
	mod := big.NewInt(15)

	want := big.NewInt(1)

	//

	res := new(big.Int)
	res = PowMod(a, n, mod)

	if want.Cmp(res) != 0 {
		t.Fatal()
	}
}

func TestPowMod4(t *testing.T) {
	a := big.NewInt(1)
	n := big.NewInt(41)
	mod := big.NewInt(1)

	want := big.NewInt(0)

	//

	res := new(big.Int)
	res = PowMod(a, n, mod)

	if want.Cmp(res) != 0 {
		t.Fatal()
	}
}

// Символ Якоби

func TestJacobi1(t *testing.T) {
	a := big.NewInt(0)
	n := big.NewInt(13)

	want := big.Jacobi(a, n)

	//

	res := Jacobi(a, n)

	if want != int(res) {
		t.Fatal()
	}
}

func TestJacobi2(t *testing.T) {
	a := big.NewInt(1)
	n := big.NewInt(13)

	want := big.Jacobi(a, n)

	//

	res := Jacobi(a, n)

	if want != int(res) {
		t.Fatal()
	}
}

func TestJacobi3(t *testing.T) {
	a := big.NewInt(2)
	n := big.NewInt(13)

	want := big.Jacobi(a, n)

	//

	res := Jacobi(a, n)

	if want != int(res) {
		t.Fatal()
	}
}

func TestJacobi4(t *testing.T) {
	a := big.NewInt(3)
	n := big.NewInt(13)

	want := big.Jacobi(a, n)

	//

	res := Jacobi(a, n)

	if want != int(res) {
		t.Fatal()
	}
}

func TestJacobi5(t *testing.T) {
	a := big.NewInt(4)
	n := big.NewInt(13)

	want := big.Jacobi(a, n)

	//

	res := Jacobi(a, n)

	if want != int(res) {
		t.Fatal()
	}
}

func TestJacobi6(t *testing.T) {
	a := big.NewInt(5)
	n := big.NewInt(13)

	want := big.Jacobi(a, n)

	//

	res := Jacobi(a, n)

	if want != int(res) {
		t.Fatal()
	}
}

func TestJacobi7(t *testing.T) {
	a := big.NewInt(6)
	n := big.NewInt(13)

	want := big.Jacobi(a, n)

	//

	res := Jacobi(a, n)

	if want != int(res) {
		t.Fatal()
	}
}

func TestJacobi8(t *testing.T) {
	a := big.NewInt(7)
	n := big.NewInt(13)

	want := big.Jacobi(a, n)

	//

	res := Jacobi(a, n)

	if want != int(res) {
		t.Fatal()
	}
}

func TestJacobi9(t *testing.T) {
	a := big.NewInt(8)
	n := big.NewInt(13)

	want := big.Jacobi(a, n)

	//

	res := Jacobi(a, n)

	if want != int(res) {
		t.Fatal()
	}
}

func TestJacobi10(t *testing.T) {
	a := big.NewInt(9)
	n := big.NewInt(13)

	want := big.Jacobi(a, n)

	//

	res := Jacobi(a, n)

	if want != int(res) {
		t.Fatal()
	}
}

func TestJacobi11(t *testing.T) {
	a := big.NewInt(10)
	n := big.NewInt(13)

	want := big.Jacobi(a, n)

	//

	res := Jacobi(a, n)

	if want != int(res) {
		t.Fatal()
	}
}

func TestJacobi12(t *testing.T) {
	a := big.NewInt(11)
	n := big.NewInt(13)

	want := big.Jacobi(a, n)

	//

	res := Jacobi(a, n)

	if want != int(res) {
		t.Fatal()
	}
}

func TestJacobi13(t *testing.T) {
	a := big.NewInt(12)
	n := big.NewInt(13)

	want := big.Jacobi(a, n)

	//

	res := Jacobi(a, n)

	if want != int(res) {
		t.Fatal()
	}
}

// Решение сравнения первой степени

func TestModuloComparisonFirst1(t *testing.T) {
	a := big.NewInt(3)
	b := big.NewInt(1)
	mod := big.NewInt(13)

	wantCount := big.NewInt(1)
	wantX1 := big.NewInt(9)

	//

	count := new(big.Int)
	x1 := new(big.Int)
	offset := new(big.Int)

	count, x1, offset = ModuloComparisonFirst(a, b, mod)

	if wantCount.Cmp(count) != 0 || wantX1.Cmp(x1) != 0 || offset != nil {
		t.Fatal()
	}
}

func TestModuloComparisonFirst2(t *testing.T) {
	a := big.NewInt(8)
	b := big.NewInt(1)
	mod := big.NewInt(5)

	wantCount := big.NewInt(1)
	wantX1 := big.NewInt(2)

	//

	count := new(big.Int)
	x1 := new(big.Int)
	offset := new(big.Int)

	count, x1, offset = ModuloComparisonFirst(a, b, mod)

	if wantCount.Cmp(count) != 0 || wantX1.Cmp(x1) != 0 || offset != nil {
		t.Fatal()
	}
}

func TestModuloComparisonFirst3(t *testing.T) {
	a := big.NewInt(5)
	b := big.NewInt(2)
	mod := big.NewInt(8)

	wantCount := big.NewInt(1)
	wantX1 := big.NewInt(2)

	//

	count := new(big.Int)
	x1 := new(big.Int)
	offset := new(big.Int)

	count, x1, offset = ModuloComparisonFirst(a, b, mod)

	if wantCount.Cmp(count) != 0 || wantX1.Cmp(x1) != 0 || offset != nil {
		t.Fatal()
	}
}

func TestModuloComparisonFirst4(t *testing.T) {
	a := big.NewInt(6)
	b := big.NewInt(1)
	mod := big.NewInt(7)

	wantCount := big.NewInt(1)
	wantX1 := big.NewInt(6)

	//

	count := new(big.Int)
	x1 := new(big.Int)
	offset := new(big.Int)

	count, x1, offset = ModuloComparisonFirst(a, b, mod)

	if wantCount.Cmp(count) != 0 || wantX1.Cmp(x1) != 0 || offset != nil {
		t.Fatal()
	}
}

func TestModuloComparisonFirst5(t *testing.T) {
	a := big.NewInt(3)
	b := big.NewInt(3)
	mod := big.NewInt(9)

	wantCount := big.NewInt(3)
	wantX1 := big.NewInt(1)
	wantOffset := big.NewInt(3)

	//

	count := new(big.Int)
	x1 := new(big.Int)
	offset := new(big.Int)

	count, x1, offset = ModuloComparisonFirst(a, b, mod)

	if wantCount.Cmp(count) != 0 || wantX1.Cmp(x1) != 0 || wantOffset.Cmp(offset) != 0 {
		t.Fatal()
	}
}

func TestModuloComparisonFirst6(t *testing.T) {
	a := big.NewInt(7)
	b := big.NewInt(0)
	mod := big.NewInt(5)

	wantCount := big.NewInt(1)
	wantX1 := big.NewInt(0)

	//

	count := new(big.Int)
	x1 := new(big.Int)
	offset := new(big.Int)

	count, x1, offset = ModuloComparisonFirst(a, b, mod)

	if wantCount.Cmp(count) != 0 || wantX1.Cmp(x1) != 0 || offset != nil {
		t.Fatal()
	}
}

func TestModuloComparisonFirst7(t *testing.T) {
	a := big.NewInt(49)
	b := big.NewInt(14)
	mod := big.NewInt(77)

	wantCount := big.NewInt(7)
	wantX1 := big.NewInt(5)
	wantOffset := big.NewInt(11)

	//

	count := new(big.Int)
	x1 := new(big.Int)
	offset := new(big.Int)

	count, x1, offset = ModuloComparisonFirst(a, b, mod)

	if wantCount.Cmp(count) != 0 || wantX1.Cmp(x1) != 0 || wantOffset.Cmp(offset) != 0 {
		t.Fatal()
	}
}

func TestModuloComparisonFirst8(t *testing.T) {
	a := big.NewInt(4)
	b := big.NewInt(3)
	mod := big.NewInt(2)

	wantCount := big.NewInt(0)

	//

	count := new(big.Int)
	x1 := new(big.Int)
	offset := new(big.Int)

	count, x1, offset = ModuloComparisonFirst(a, b, mod)

	if wantCount.Cmp(count) != 0 || x1 != nil || offset != nil {
		t.Fatal()
	}
}
