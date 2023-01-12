package ciphers

import "math/big"

// RSA

// Публичный ключ

type PublicKeyRSA struct {
	PublicExponent *big.Int
	N              *big.Int
}

func NewPublicKeyRSA(e *big.Int, n *big.Int) *PublicKeyRSA {

	key := PublicKeyRSA{
		PublicExponent: e,
		N:              n,
	}

	return &key
}

// Приватный ключ

type PrivateKeyRSA struct {
	PrivateExponent *big.Int
	Prime1          *big.Int
	Prime2          *big.Int
}

func NewPrivateKeyRSA(d *big.Int, p *big.Int, q *big.Int) *PrivateKeyRSA {

	key := PrivateKeyRSA{
		PrivateExponent: d,
		Prime1:          p,
		Prime2:          q,
	}

	return &key
}

// Шифр Рабина

// Публичный ключ

type PublicKeyRabin struct {
	N *big.Int
}

func NewPublicKeyRabin(n *big.Int) *PublicKeyRabin {

	key := PublicKeyRabin{
		N: n,
	}

	return &key
}

// Приватный ключ

type PrivateKeyRabin struct {
	Prime1 *big.Int
	Prime2 *big.Int
}

func NewPrivateKeyRabin(p *big.Int, q *big.Int) *PrivateKeyRabin {

	key := PrivateKeyRabin{
		Prime1: p,
		Prime2: q,
	}

	return &key
}

// Шифр Эль-Гамаля

// Публичный ключ

type PublicKeyElGamal struct {
	p     *big.Int
	alpha *big.Int
	beta  *big.Int
}

func NewPublicKeyElGamal(p *big.Int, alpha *big.Int, beta *big.Int) *PublicKeyElGamal {

	key := PublicKeyElGamal{
		p:     p,
		alpha: alpha,
		beta:  beta,
	}

	return &key
}

// Приватный ключ

type PrivateKeyElGamal struct {
	a *big.Int
}

func NewPrivateKeyElGamal(a *big.Int) *PrivateKeyElGamal {

	key := PrivateKeyElGamal{
		a: a,
	}

	return &key
}
