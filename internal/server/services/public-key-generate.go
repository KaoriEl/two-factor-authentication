package services

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base32"
	"fmt"
	_struct "main/internal/struct"
	"strconv"
	"strings"
	"time"
)

func GivePublicKey(v _struct.Default, c chan string) {
	input := v.SecretKey
	fmt.Println(input)
	// декодируем ключ из первого аргумента

	key, _ := convertSecret(input)

	// генерируем одноразовый пароль, используя время с 30-секундными интервалами
	epochSeconds := time.Now().Unix()

	pwd := oneTimePassword(key, toBytes(epochSeconds/30))

	c <- strconv.Itoa(int(pwd))

}

func convertSecret(secret string) ([]byte, error) {
	inputNoSpaces := strings.Replace(secret, " ", "", -1)
	decodeKey, err := base32.StdEncoding.DecodeString(checkSecret(strings.ToUpper(inputNoSpaces)))
	if err != nil {
		return nil, err
	}
	return decodeKey, nil
}

func checkSecret(secret string) string {
	length := len(secret)
	if length%8 == 0 {
		return secret
	}
	n := length/8*8 + 8 - length
	return secret + strings.Repeat("=", n)
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
	fmt.Println(toUint32(hashParts))
	number := toUint32(hashParts)

	// размер до 6 цифр
	// один миллион - это первое число из 7 цифр, поэтому остаток
	// деления всегда будет возвращать < 7 цифр
	fmt.Println(number)
	pwd := number % 1000000

	return pwd
}
