package desctr

import (
	"crypto/des"
	"crypto/cipher"
)

func DesEnCrypt(sourceData, key []byte) []byte {
	block, err := des.NewCipher(key)
	if err != nil {
		panic(err)
	}
	iv := []byte("12345678")
	stream := cipher.NewCTR(block, iv)
	stream.XORKeyStream(sourceData, sourceData)
	return sourceData
}

func DesDeCrypt(sourceData, key []byte) []byte {
	block, err := des.NewCipher(key)
	if err != nil {
		panic(err)
	}
	iv := []byte("12345678")
	stream := cipher.NewCTR(block, iv)
	stream.XORKeyStream(sourceData, sourceData)
	return sourceData
}
