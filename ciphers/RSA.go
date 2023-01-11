package ciphers

import (
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
	constNum3 = big.NewInt(3)
	constNum4 = big.NewInt(4)
)

// RSA

type RSA struct {
	publicKey  *PublicKeyRSA
	privateKey *PrivateKeyRSA
}

// GenerateKey - Генерирует PublicKeyRSA и PrivateKeyRSA и задает их в структуру RSA
func (rsa *RSA) GenerateKey(k int) *RSA {

	const iter = 150

	// Проверка входных данных
	if k <= 1 {
		panic("k > 1")
	}

	p := new(big.Int)
	q := new(big.Int)
	n := new(big.Int)
	phi := new(big.Int)

	e := new(big.Int)
	d := new(big.Int)

	// Генерация

	p = crypto_math.SimpleNumber(k, iter)
	q = crypto_math.SimpleNumber(k, iter)

	// Если p и q равны генерируем новое q
	for p.Cmp(q) == 0 {
		q = crypto_math.SimpleNumber(k, iter)
	}

	n = n.Mul(p, q)
	phi = phi.Mul(
		new(big.Int).Sub(p, constNum1),
		new(big.Int).Sub(q, constNum1),
	)

	// Выбираем e

	// Выбираем открытую экспоненту
	e = crypto_math.SimpleNumber(k/8, iter)

	for {
		d = crypto_math.InverseElement(e, phi)

		gcd := new(big.Int)
		gcd = crypto_math.EuclidAlgorithm(e, phi)

		if gcd.Cmp(constNum1) == 0 && new(big.Int).Mod(new(big.Int).Mul(e, d), phi).Cmp(constNum1) == 0 {
			break
		}
		e = e.Add(e, constNum1)
	}

	// Устанавливаем ключи

	rsa.publicKey = NewPublicKeyRSA(e, n)
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

	privateKeyName := "ciphers/RSA/PrivateKey_" + date + ".json"
	privateKeyFile, err := os.Create(privateKeyName)

	_, err = privateKeyFile.Write(privateKeyJSON)
	if err != nil {
		panic(err)
	}
}

// LoadKeys - загружает ключи из JSON файлов и задает их в структуру RSA
func (rsa *RSA) LoadKeys(publicKeyPath string, privateKeyPath string) *RSA {

	// Публичный ключ
	if publicKeyPath != "" {

		publicKey, err := os.ReadFile(publicKeyPath)
		if err != nil {
			panic(err)
		}
		err = json.Unmarshal(publicKey, &rsa.publicKey)
	}

	// Приватный ключ
	if privateKeyPath != "" {

		privateKey, err := os.ReadFile(privateKeyPath)
		if err != nil {
			panic(err)
		}
		err = json.Unmarshal(privateKey, &rsa.privateKey)
	}

	return rsa
}

// Encrypt - Шифрует последовательность байт
func (rsa *RSA) Encrypt(message []byte) []byte {

	if rsa.publicKey == nil {
		panic("No public key specified")
	}

	// Загружаем ключ
	e := new(big.Int)
	n := new(big.Int)

	e.Set(rsa.publicKey.PublicExponent)
	n.Set(rsa.publicKey.N)

	// Размер блока для шифррования
	blockSize := len(n.Bytes()) - 1

	// Бьем сообщение на блоки
	messageBlocks := ToBlocks(message, blockSize)

	cipherBlocks := []*big.Int{}
	temp := new(big.Int)

	// Шифруем
	for i := 0; i < len(messageBlocks); i++ {
		temp.Set(messageBlocks[i])

		temp = crypto_math.PowMod(temp, e, n)
		cipherBlocks = append(cipherBlocks, new(big.Int).Set(temp))
	}

	// Переводим блоки в байты
	result := ToCipherBytes(cipherBlocks, blockSize+1)

	return result
}

// Decrypt - Расшифровывает последовательность байт
func (rsa *RSA) Decrypt(ciphertext []byte) []byte {

	if rsa.privateKey == nil {
		panic("No private key specified")
	}

	// Загружаем ключ
	n := new(big.Int).Mul(
		rsa.privateKey.Prime1,
		rsa.privateKey.Prime2,
	)

	d := new(big.Int)
	d.Set(rsa.privateKey.PrivateExponent)

	// Размер блока для шифррования
	blockSize := len(n.Bytes())

	// Бьем сообщение на блоки
	cipherBlocks := ToCipherBlocks(ciphertext, blockSize)

	messageBlocks := []*big.Int{}
	temp := new(big.Int)

	// Расшифровываем
	for i := 0; i < len(cipherBlocks); i++ {
		temp.Set(cipherBlocks[i])

		temp = crypto_math.PowMod(temp, d, n)
		messageBlocks = append(messageBlocks, new(big.Int).Set(temp))
	}

	// Переводим блоки в байты
	result := ToBytes(messageBlocks)

	return result
}
