package ciphers

import (
	"crypto/rand"
	"cryptography/crypto_math"
	"encoding/json"
	"math/big"
	"os"
	"time"
)

// Шифр Эль-Гамаля

type ElGamal struct {
	publicKey  *PublicKeyElGamal
	privateKey *PrivateKeyElGamal
}

// GenerateKey - Генерирует PublicKeyRSA и PrivateKeyRSA и задает их в структуру RSA
func (elg *ElGamal) GenerateKey(k int) *ElGamal {

	const iter = 150

	// Проверка входных данных
	if k <= 1 {
		panic("k > 1")
	}

	p := new(big.Int)
	a := new(big.Int)

	alpha := new(big.Int)
	beta := new(big.Int)

	// Генерация

	p = crypto_math.SimpleNumber(k, iter)

	// Генерируем случайное число а

	a, err := rand.Int(
		rand.Reader,
		new(big.Int).Sub(p, constNum3),
	)
	if err != nil {
		panic(err)
	}
	a = a.Add(a, constNum2)
	//

	p2 := new(big.Int).Div(new(big.Int).Sub(p, constNum1), constNum2)

	// Генерируем альфу
	for {
		alpha, err = rand.Int(
			rand.Reader,
			new(big.Int).Sub(p, constNum4),
		)
		if err != nil {
			panic(err)
		}
		alpha = alpha.Add(alpha, constNum3)

		conditionOne := crypto_math.PowMod(alpha, p2, p)
		if conditionOne.Cmp(constNum1) != 0 {

			conditionTwo := crypto_math.PowMod(
				alpha,
				new(big.Int).Div(new(big.Int).Sub(p, constNum1), p2),
				p,
			)

			if conditionTwo.Cmp(constNum1) != 0 {
				break
			}
		}
	}

	// Вычисляем бета
	beta = crypto_math.PowMod(alpha, a, p)

	// Устанавливаем ключи

	elg.publicKey = NewPublicKeyElGamal(p, alpha, beta)
	elg.privateKey = NewPrivateKeyElGamal(p, a)

	return elg
}

// SaveKeys - сохраняет ключи в JSON файлы
func (elg *ElGamal) SaveKeys() {

	// Время записи
	date := time.Now().Format("150405")

	// Публичный ключ
	publicKeyJSON, err := json.MarshalIndent(elg.publicKey, "", "\t")
	if err != nil {
		panic(err)
	}

	publicKeyName := "ciphers/ElGamal/PublicKey_" + date + ".json"
	publicKeyFile, err := os.Create(publicKeyName)

	defer publicKeyFile.Close()

	_, err = publicKeyFile.Write(publicKeyJSON)
	if err != nil {
		panic(err)
	}

	// Приватный ключ
	privateKeyJSON, err := json.MarshalIndent(elg.privateKey, "", "\t")
	if err != nil {
		panic(err)
	}

	privateKeyName := "ciphers/ElGamal/PrivateKey_" + date + ".json"
	privateKeyFile, err := os.Create(privateKeyName)

	_, err = privateKeyFile.Write(privateKeyJSON)
	if err != nil {
		panic(err)
	}
}

// LoadKeys - загружает ключи из JSON файлов и задает их в структуру ElGamal
func (elg *ElGamal) LoadKeys(publicKeyPath string, privateKeyPath string) *ElGamal {

	// Публичный ключ
	if publicKeyPath != "" {

		publicKey, err := os.ReadFile(publicKeyPath)
		if err != nil {
			panic(err)
		}
		err = json.Unmarshal(publicKey, &elg.publicKey)
	}

	// Приватный ключ
	if privateKeyPath != "" {

		privateKey, err := os.ReadFile(privateKeyPath)
		if err != nil {
			panic(err)
		}
		err = json.Unmarshal(privateKey, &elg.privateKey)
	}

	return elg
}

// Encrypt - Шифрует последовательность байт
func (elg *ElGamal) Encrypt(message []byte) ([]byte, []byte) {

	if elg.publicKey == nil {
		panic("No public key specified")
	}

	// Загружаем ключ
	p := new(big.Int)
	alpha := new(big.Int)
	beta := new(big.Int)

	p.Set(elg.publicKey.P)
	alpha.Set(elg.publicKey.Alpha)
	beta.Set(elg.publicKey.Beta)

	// Размер блока для шифррования
	blockSize := len(p.Bytes()) - 1

	// Бьем сообщение на блоки
	messageBlocks := ToBlocks(message, blockSize)

	cipherBlocks1 := []*big.Int{}
	cipherBlocks2 := []*big.Int{}

	temp := new(big.Int)

	cipher1 := new(big.Int)
	cipher2 := new(big.Int)

	// Шифруем
	for i := 0; i < len(messageBlocks); i++ {
		temp.Set(messageBlocks[i])

		// Случайное r

		r, err := rand.Int(
			rand.Reader,
			new(big.Int).Sub(p, constNum3),
		)
		if err != nil {
			panic(err)
		}
		r = r.Add(r, constNum2)
		//

		// Первое сообщение
		cipher1 = crypto_math.PowMod(alpha, r, p)
		cipherBlocks1 = append(cipherBlocks1, new(big.Int).Set(cipher1))

		// Второе сообщение
		cipher2 = crypto_math.PowMod(beta, r, p)
		cipher2 = cipher2.Mul(cipher2, temp)
		cipher2 = cipher2.Mod(cipher2, p)
		cipherBlocks2 = append(cipherBlocks2, new(big.Int).Set(cipher2))
	}

	// Переводим блоки в байты
	res1 := ToCipherBytes(cipherBlocks1, blockSize+1)
	res2 := ToCipherBytes(cipherBlocks2, blockSize+1)

	return res1, res2
}

// Decrypt - Расшифровывает последовательность байт
func (elg *ElGamal) Decrypt(ciphertext1 []byte, ciphertext2 []byte) []byte {

	if elg.privateKey == nil {
		panic("No private key specified")
	}

	p := new(big.Int)
	a := new(big.Int)
	p.Set(elg.privateKey.P)
	a.Set(elg.privateKey.A)

	// Размер блока для расшифррования
	blockSize := len(p.Bytes())

	// Бьем сообщение на блоки
	cipherBlocks1 := ToCipherBlocks(ciphertext1, blockSize)
	cipherBlocks2 := ToCipherBlocks(ciphertext2, blockSize)

	messageBlocks := []*big.Int{}

	temp := new(big.Int)

	cipher1 := new(big.Int)
	cipher2 := new(big.Int)

	// Расшифровываем
	for i := 0; i < len(cipherBlocks1); i++ {
		cipher1.Set(cipherBlocks1[i])
		cipher2.Set(cipherBlocks2[i])

		temp = crypto_math.PowMod(cipher1, a, p)
		temp = crypto_math.InverseElement(temp, p)
		temp = temp.Mul(temp, cipher2)
		temp = temp.Mod(temp, p)

		messageBlocks = append(messageBlocks, new(big.Int).Set(temp))
	}

	// Переводим блоки в байты
	result, _ := ToBytes(messageBlocks)

	return result
}
