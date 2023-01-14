package ciphers

import (
	"cryptography/crypto_math"
	"encoding/json"
	"math/big"
	"os"
	"time"
)

// Шифр Рабина

type Rabin struct {
	publicKey  *PublicKeyRabin
	privateKey *PrivateKeyRabin
}

// GenerateKey - Генерирует PublicKeyRabin и PrivateKeyRabin и задает их в структуру Rabin
func (rab *Rabin) GenerateKey(k int) *Rabin {

	const iter = 150

	// Проверка входных данных
	if k <= 1 {
		panic("k > 1")
	}

	p := new(big.Int)
	q := new(big.Int)
	n := new(big.Int)

	// Генерация p с проверкой
	for {
		p = crypto_math.SimpleNumber(k, iter)

		if new(big.Int).Mod(new(big.Int).Sub(p, constNum3), constNum4).Sign() == 0 {
			break
		}
	}

	// Генерация q с проверкой
	for {
		q = crypto_math.SimpleNumber(k, iter)

		if new(big.Int).Mod(new(big.Int).Sub(q, constNum3), constNum4).Sign() == 0 && p.Cmp(q) != 0 {
			break
		}
	}

	n = n.Mul(p, q)

	// Устанавливаем ключи

	rab.publicKey = NewPublicKeyRabin(n)
	rab.privateKey = NewPrivateKeyRabin(p, q)

	return rab
}

// SaveKeys - сохраняет ключи в JSON файлы
func (rab *Rabin) SaveKeys() {

	// Время записи
	date := time.Now().Format("150405")

	// Публичный ключ
	publicKeyJSON, err := json.MarshalIndent(rab.publicKey, "", "\t")
	if err != nil {
		panic(err)
	}

	publicKeyName := "ciphers/Rabin/PublicKey_" + date + ".json"
	publicKeyFile, err := os.Create(publicKeyName)

	defer publicKeyFile.Close()

	_, err = publicKeyFile.Write(publicKeyJSON)
	if err != nil {
		panic(err)
	}

	// Приватный ключ
	privateKeyJSON, err := json.MarshalIndent(rab.privateKey, "", "\t")
	if err != nil {
		panic(err)
	}

	privateKeyName := "ciphers/Rabin/PrivateKey_" + date + ".json"
	privateKeyFile, err := os.Create(privateKeyName)

	_, err = privateKeyFile.Write(privateKeyJSON)
	if err != nil {
		panic(err)
	}
}

// LoadKeys - загружает ключи из JSON файлов и задает их в структуру Rabin
func (rab *Rabin) LoadKeys(publicKeyPath string, privateKeyPath string) *Rabin {

	// Публичный ключ
	if publicKeyPath != "" {

		publicKey, err := os.ReadFile(publicKeyPath)
		if err != nil {
			panic(err)
		}
		err = json.Unmarshal(publicKey, &rab.publicKey)
	}

	// Приватный ключ
	if privateKeyPath != "" {

		privateKey, err := os.ReadFile(privateKeyPath)
		if err != nil {
			panic(err)
		}
		err = json.Unmarshal(privateKey, &rab.privateKey)
	}

	return rab
}

// Encrypt - Шифрует последовательность байт
func (rab *Rabin) Encrypt(message []byte) []byte {

	if rab.publicKey == nil {
		panic("No public key specified")
	}

	// Загружаем ключ
	n := new(big.Int)
	n.Set(rab.publicKey.N)

	// Размер блока для шифррования
	blockSize := len(n.Bytes()) - 1

	// Бьем сообщение на блоки
	messageBlocks := ToBlocks(message, blockSize)

	cipherBlocks := []*big.Int{}
	temp := new(big.Int)

	// Шифруем
	for i := 0; i < len(messageBlocks); i++ {
		temp.Set(messageBlocks[i])

		temp = crypto_math.PowMod(temp, constNum2, n)
		cipherBlocks = append(cipherBlocks, new(big.Int).Set(temp))
	}

	// Переводим блоки в байты
	result := ToCipherBytes(cipherBlocks, blockSize+1)

	return result
}

// Decrypt - Расшифровывает последовательность байт
func (rab *Rabin) Decrypt(ciphertext []byte) []byte {

	if rab.privateKey == nil {
		panic("No private key specified")
	}

	p := new(big.Int)
	q := new(big.Int)
	yp := new(big.Int)
	yq := new(big.Int)
	gcd := new(big.Int)

	p = rab.privateKey.Prime1
	q = rab.privateKey.Prime2

	// Загружаем ключ
	n := new(big.Int).Mul(p, q)

	gcd, yp, yq = crypto_math.AdvancedEuclidAlgorithm(p, q)

	if gcd.Cmp(constNum1) != 0 {
		panic("p and q are not mutually simple")
	}

	// Размер блока для расшифррования
	blockSize := len(n.Bytes())

	// Бьем сообщение на блоки
	cipherBlocks := ToCipherBlocks(ciphertext, blockSize)

	messageBlocks1 := []*big.Int{}
	messageBlocks2 := []*big.Int{}
	messageBlocks3 := []*big.Int{}
	messageBlocks4 := []*big.Int{}

	temp := new(big.Int)

	m1 := new(big.Int)
	m2 := new(big.Int)
	m3 := new(big.Int)
	m4 := new(big.Int)
	mp := new(big.Int)
	mq := new(big.Int)

	// Расшифровываем
	for i := 0; i < len(cipherBlocks); i++ {
		temp.Set(cipherBlocks[i])

		mp, _ = crypto_math.ModuloComparisonSecond(temp, p)
		mq, _ = crypto_math.ModuloComparisonSecond(temp, q)

		left := new(big.Int)
		right := new(big.Int)

		left = left.Mul(yp, p)
		left = left.Mod(left, n)
		left = left.Mul(left, mq)
		left = left.Mod(left, n)

		right = right.Mul(yq, q)
		right = right.Mod(right, n)
		right = right.Mul(right, mp)
		right = right.Mod(right, n)

		// m1
		m1 = m1.Add(left, right)
		m1 = m1.Mod(m1, n)

		// m2
		m2 = m2.Sub(n, m1)

		// m3
		m3 = m3.Sub(left, right)
		m3 = m3.Mod(m3, n)

		// m4
		m4 = m4.Sub(n, m3)

		messageBlocks1 = append(messageBlocks1, new(big.Int).Set(m1))
		messageBlocks2 = append(messageBlocks2, new(big.Int).Set(m2))
		messageBlocks3 = append(messageBlocks3, new(big.Int).Set(m3))
		messageBlocks4 = append(messageBlocks4, new(big.Int).Set(m4))
	}

	result := []byte{}

	it := len(messageBlocks1)

	// Собираем 4 массива блоков в 1 сообщение
	for b := 0; b < it; b++ {

		arr := messageBlocks1[b : b+1]

		msg, err := ToBytes(arr)

		if err == nil {
			result = append(result, msg...)
			continue
		}

		arr = messageBlocks2[b : b+1]

		msg, err = ToBytes(arr)

		if err == nil {
			result = append(result, msg...)
			continue
		}

		arr = messageBlocks3[b : b+1]

		msg, err = ToBytes(arr)

		if err == nil {
			result = append(result, msg...)
			continue
		}

		arr = messageBlocks4[b : b+1]

		msg, err = ToBytes(arr)

		if err == nil {
			result = append(result, msg...)
			continue
		}
	}

	return result
}
