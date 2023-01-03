package ciphers

import (
	"math/big"
)

// Пакет для разбития набора байт на блоки

const blockSizeBits uint = 2048

// ToBlocks - Разбивает набор незашифрованных байт на блоки, дополняняя недостающие октеты, возвращает массив незашифрованных big.Int
func ToBlocks(_byteArray []byte) []*big.Int {

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

// ToBytes - Преобразует массив незашифрованных big.Int в массив незашифрованных байт
func ToBytes(blocksArray []*big.Int) []byte {

	byteArray := []byte{}

	// Преобразуем числа в байты
	for i := 0; i < len(blocksArray); i++ {

		temp := blocksArray[i].Bytes()
		byteArray = append(byteArray, temp...)
	}

	// Убираем падинг

	// Получаем последний байт
	lastByte := byteArray[len(byteArray)-1]
	lastByteValue := int(lastByte)

	if lastByteValue < len(byteArray) {
		byteArray = byteArray[:len(byteArray)-lastByteValue]
	}

	return byteArray
}

// ToCipherBytes - Преобразует массив зашифрованных big.Int в массив зашифрованных байт
func ToCipherBytes(blocksArray []*big.Int) []byte {

	byteArray := []byte{}

	// Преобразуем числа в байты
	for i := 0; i < len(blocksArray); i++ {

		temp := blocksArray[i].Bytes()

		// Количество дополняемых незначащих нулей слева

		count := (int(blockSizeBits) / 8) - len(temp)

		for j := 0; j < count; j++ {
			byteArray = append(byteArray, byte(0))
		}

		byteArray = append(byteArray, temp...)
	}

	return byteArray
}

// ToCipherBlocks - Разбивает набор зашифрованных байт на блоки, возвращает массив зашифрованных big.Int
func ToCipherBlocks(_byteArray []byte) []*big.Int {

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

	return result
}
