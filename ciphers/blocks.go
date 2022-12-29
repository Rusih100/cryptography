package ciphers

import (
	"math/big"
)

// Пакет для разбития набора байт на блоки

// ToBlocks - Разбивает набор байт на блоки, дополняняя недостающие октеты, возвращает массив big.Int
func ToBlocks(_byteArray []byte, blockSizeBits uint) []*big.Int {

	// Размер массива байт
	byteArraySize := len(_byteArray)

	// Копируем массив байт, чтобы не изменять его
	byteArray := make([]byte, byteArraySize)
	copy(byteArray, _byteArray)

	// Размер блока в байтах
	blockSize := int(blockSizeBits) / 8

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

	// Последний блок
	lastBlockBytes := []byte{}

	// Количество имеющихся октетов

	lastBlockBytes = byteArray[blockCount*blockSize:]
	octetsCount := len(lastBlockBytes)

	if octetsCount != 0 {
		// Случай не кратного блока, дополняем его по PKCS7

		value := (blockSize - octetsCount) % 256
		valueByte := byte(value) // Значение дополняемых октетов

		// Дополняем недостающие октеты
		for i := 0; i < blockSize-octetsCount; i++ {
			lastBlockBytes = append(lastBlockBytes, valueByte)
		}

	} else {
		// Случай кратного блока

		value := blockSize % 256
		valueByte := byte(value) // Значение дополняемых октетов

		// Дополняем недостающие октеты
		for i := 0; i < blockSize; i++ {
			lastBlockBytes = append(lastBlockBytes, valueByte)
		}
	}

	result = append(result, new(big.Int).SetBytes(lastBlockBytes))

	// Добаляем последний блок в массив

	return result
}

// ToBytes - Преобразует массив big.Int в массив байт
func ToBytes(blocksArray []*big.Int) []byte {
	return nil

	// TODO
}
