package finite_field

import (
	"cryptography/cryptography"
	"cryptography/polynomial"
	"math/big"
)

// Реализация расширения базового конечного поля

type GaloisField struct {
	p    *big.Int
	poly *polynomial.Polynomial
}

func (g *GaloisField) Set(p *big.Int, polynomial *polynomial.Polynomial) *GaloisField {

	// Проверка входных данных

	g.p = big.NewInt(0)

	if p.Cmp(constNum2) < 0 {
		panic("p >= 2")
	}

	if !cryptography.MillerRabinTest(p) {
		panic("p is a prime number")
	}

	g.p.Set(p)

	return g
}
