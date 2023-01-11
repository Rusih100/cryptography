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
