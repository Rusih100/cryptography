package ciphers

import (
	"fmt"
	"math/big"
)

// Пакет для разбития набора байт на блоки

// ToBlocks - Разбивает набор незашифрованных байт на блоки, дополняняя недостающие октеты, возвращает массив незашифрованных big.Int
//
// Примечание: blockSize - ожидаемый размер блока учетом падинга
func ToBlocks(_byteArray []byte, blockSize int) []*big.Int {

	// Размер массива байт
	byteArraySize := len(_byteArray)

	// Копируем массив байт, чтобы не изменять его
	byteArray := make([]byte, byteArraySize)
	copy(byteArray, _byteArray)

	// Количество цельных блоков
	blockCount := byteArraySize / (blockSize - 1)

	result := []*big.Int{}
	tempBytes := []byte{}

	// Добавляем цельные блоки
	for i := 0; i < blockCount; i++ {

		// Срез байтов
		tempBytes = byteArray[i*(blockSize-1) : (i+1)*(blockSize-1)]

		paddingBytes := make([]byte, len(tempBytes))
		copy(paddingBytes, tempBytes)

		// Добавляем единичный падинг
		paddingBytes = append(paddingBytes, byte(1))

		// Добавляем в массив число
		result = append(result, new(big.Int).SetBytes(paddingBytes))
	}

	// Последний блок
	lastBlockBytes := []byte{}

	// Количество имеющихся октетов

	lastBlockBytes = byteArray[blockCount*(blockSize-1):]
	octetsCount := len(lastBlockBytes)

	if octetsCount != 0 {
		// Случай не кратного блока, дополняем его по PKCS7

		value := (blockSize - octetsCount) % 256
		valueByte := byte(value) // Значение дополняемых октетов

		// Дополняем недостающие октеты
		for i := 0; i < blockSize-octetsCount; i++ {
			lastBlockBytes = append(lastBlockBytes, valueByte)
		}

		// Добовляем последний блок в массив
		result = append(result, new(big.Int).SetBytes(lastBlockBytes))
	}

	return result
}

// ToCipherBlocks - Разбивает набор зашифрованных байт на блоки, возвращает массив зашифрованных big.Int
func ToCipherBlocks(_byteArray []byte, blockSize int) []*big.Int {

	// Размер массива байт
	byteArraySize := len(_byteArray)

	// Копируем массив байт, чтобы не изменять его
	byteArray := make([]byte, byteArraySize)
	copy(byteArray, _byteArray)

	// Количество цельных блоков
	blockCount := byteArraySize / blockSize

	result := []*big.Int{}
	tempBytes := []byte{}

	// Добавляем цельные блоки
	for i := 0; i < blockCount; i++ {

		// Срез байтов
		tempBytes = byteArray[i*blockSize : (i+1)*blockSize]

		// Добавляем в массив число
		result = append(result, new(big.Int).SetBytes(tempBytes))
	}

	return result
}

// ToBytes - Преобразует массив незашифрованных big.Int в массив незашифрованных байт
func ToBytes(blocksArray []*big.Int) ([]byte, error) {

	byteArray := []byte{}

	// Преобразуем числа в байты
	for i := 0; i < len(blocksArray); i++ {

		temp := blocksArray[i].Bytes()

		// Убираем падинг

		lastByte := temp[len(temp)-1]
		lastByteValue := int(lastByte)

		for k := 0; k < lastByteValue; k++ {
			if int(temp[len(temp)-1]) == lastByteValue {
				// Удаляем последний элемент
				temp = append(temp[:len(temp)-1], temp[len(temp):]...)

			} else {
				return nil, fmt.Errorf("error when removing padding")
			}
		}

		byteArray = append(byteArray, temp...)

	}

	return byteArray, nil
}

// ToCipherBytes - Преобразует массив зашифрованных big.Int в массив зашифрованных байт
func ToCipherBytes(blocksArray []*big.Int, blockSize int) []byte {

	byteArray := []byte{}

	null := byte(0)

	// Преобразуем числа в байты
	for i := 0; i < len(blocksArray); i++ {

		temp := blocksArray[i].Bytes()

		nullCount := blockSize - len(temp)

		for j := 0; j < nullCount; j++ {
			byteArray = append(byteArray, null)
		}

		byteArray = append(byteArray, temp...)

		temp = nil
	}

	return byteArray
}
