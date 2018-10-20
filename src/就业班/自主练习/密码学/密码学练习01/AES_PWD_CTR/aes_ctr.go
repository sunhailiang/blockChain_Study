package aesctr

import (
	"crypto/aes"
	"crypto/cipher"
)

func AesEnCrypt(sourceData, key []byte) []byte {
	//创建采用aes算法协议,创建加密对象
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	//创建ctr分组模式
	iv := []byte("12345678abcdefgh")
	strem := cipher.NewCTR(block, iv)
	//加密
	strem.XORKeyStream(sourceData, sourceData)
	return sourceData
}
func AesDeCrypt(sourceData, key []byte) []byte {
	//采用aes算法协议，创建加密对象
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	//创建ctr分组模式
	iv := []byte("12345678abcdefgh")
	stream := cipher.NewCTR(block, iv)
	//解密
	stream.XORKeyStream(sourceData, sourceData)
	return sourceData

}
