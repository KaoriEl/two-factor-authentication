package services

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base32"
	"fmt"
	_struct "main/internal/struct"
	"os"
	"strconv"
	"strings"
	"time"
)

func GivePublicKey(v _struct.Default, c chan string) {
	fmt.Println("3")
	input := v.SecretKey

	// декодируем ключ из первого аргумента
	inputNoSpaces := strings.Replace(input, " ", "", -1)
	inputNoSpacesUpper := strings.ToUpper(inputNoSpaces)
	key, err := base32.StdEncoding.WithPadding(base32.NoPadding).DecodeString(inputNoSpacesUpper)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	// генерируем одноразовый пароль, используя время с 30-секундными интервалами
	epochSeconds := time.Now().Unix()
	pwd := oneTimePassword(key, toBytes(epochSeconds/30))

	c <- strconv.Itoa(int(pwd))

}

func toBytes(value int64) []byte {
	var result []byte
	mask := int64(0xFF)
	shifts := [8]uint16{56, 48, 40, 32, 24, 16, 8, 0}
	for _, shift := range shifts {
		result = append(result, byte((value>>shift)&mask))
	}
	return result
}

func toUint32(bytes []byte) uint32 {
	return (uint32(bytes[0]) << 24) + (uint32(bytes[1]) << 16) +
		(uint32(bytes[2]) << 8) + uint32(bytes[3])
}

func oneTimePassword(key []byte, value []byte) uint32 {
	// подписываем значение с помощью HMAC-SHA1
	hmacSha1 := hmac.New(sha1.New, key)
	hmacSha1.Write(value)
	hash := hmacSha1.Sum(nil)

	// Мы собираемся использовать подмножество сгенерированного хеша.
	// Использование последнего полубайта для выбора индекса, с которого следует начать.
	// Это число всегда подходит, так как максимальное десятичное число равно 15, хэш будет
	// максимальный индекс 19 (20 байт SHA1) и нам нужно 4 байта.
	offset := hash[len(hash)-1] & 0x0F
	// получить 32-битный (4-байтовый) кусок из хэша, начиная со смещения
	hashParts := hash[offset : offset+4]

	// игнорировать старший бит согласно RFC 4226
	hashParts[0] = hashParts[0] & 0x7F

	number := toUint32(hashParts)

	// размер до 6 цифр
	// один миллион - это первое число из 7 цифр, поэтому остаток
	// деления всегда будет возвращать < 7 цифр
	pwd := number % 100000

	return pwd
}
