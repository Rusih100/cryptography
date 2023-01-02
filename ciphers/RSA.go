package ciphers

import (
	"crypto/rand"
	"cryptography/crypto_math"
	"encoding/json"
	"math/big"
	"os"
	"time"
)

// Константы для упрощения кода
// Не изменять в алгоритмах!
var (
	constNum1 = big.NewInt(1)
)

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

// RSA

type RSA struct {
	publicKey  *PublicKeyRSA
	privateKey *PrivateKeyRSA
}

// GenerateKey - Генерирует PublicKeyRSA и PrivateKeyRSA и задает их в структуру RSA
func (rsa *RSA) GenerateKey(k int) *RSA {

	// Проверка входных данных
	if k <= 1 {
		panic("k > 1")
	}

	p := new(big.Int)
	q := new(big.Int)

	// Генерация случайных p и q
	p = crypto_math.SimpleNumber(k, 100)
	q = crypto_math.SimpleNumber(k, 100)

	// Случай одинаковых p и q
	for p.Cmp(q) == 0 {
		q = crypto_math.SimpleNumber(k, 100)
	}

	n := new(big.Int).Mul(p, q)

	phi := new(big.Int).Mul(
		new(big.Int).Sub(p, constNum1),
		new(big.Int).Sub(q, constNum1),
	)

	e := big.NewInt(0)

	for {
		// Выбираем открытую экспоненту
		exp, err := rand.Int(
			rand.Reader,
			new(big.Int).Sub(phi, constNum1),
		)
		if err != nil {
			panic(err)
		}
		exp = exp.Add(exp, constNum1)
		//

		gcd := new(big.Int)
		gcd = crypto_math.EuclidAlgorithm(exp, phi)

		if gcd.Cmp(constNum1) == 1 {
			e.Set(exp)
			break
		}
	}

	d := new(big.Int)

	d = crypto_math.InverseElement(e, phi)

	// Устанавливаем ключи

	rsa.publicKey = NewPublicKeyRSA(n, e)
	rsa.privateKey = NewPrivateKeyRSA(d, p, q)

	return rsa
}

// SaveKeys - сохраняет ключи в JSON файлы
func (rsa *RSA) SaveKeys() {

	// Время записи
	date := time.Now().Format("150405")

	// Публичный ключ
	publicKeyJSON, err := json.MarshalIndent(rsa.publicKey, "", "\t")
	if err != nil {
		panic(err)
	}

	publicKeyName := "ciphers/RSA/PublicKey_" + date + ".json"
	publicKeyFile, err := os.Create(publicKeyName)

	defer publicKeyFile.Close()

	_, err = publicKeyFile.Write(publicKeyJSON)
	if err != nil {
		panic(err)
	}

	// Приватный ключ
	privateKeyJSON, err := json.MarshalIndent(rsa.privateKey, "", "\t")
	if err != nil {
		panic(err)
	}

	privateKeyName := "ciphers/RSA/PrivateKey_ " + date + ".json"
	privateKeyFile, err := os.Create(privateKeyName)

	_, err = privateKeyFile.Write(privateKeyJSON)
	if err != nil {
		panic(err)
	}
}
