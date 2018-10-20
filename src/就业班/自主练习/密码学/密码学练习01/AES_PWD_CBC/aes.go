package aes

import (
	"crypto/aes"
	"就业班/自主练习/密码学/密码学练习01/public"
	"crypto/cipher"
)

//aes key的长度8
func AesEnCrypt(sourceData, key []byte) []byte {
	//创建aes密码块
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	//填充分组数据
	sourceData = public.FillLastGroup(sourceData, block.BlockSize())
	//使用aes算法
	iv := "1234567812345678"
	//创建分组链接模式
	blockModel := cipher.NewCBCEncrypter(block, []byte(iv))
	//加密
	blockModel.CryptBlocks(sourceData, sourceData)
	return sourceData
}
func AesDeCrypt(sourceData, key []byte) []byte {
	//创建aes密码块
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	//创建aes解密接口
	iv := []byte("1234567812345678")
	//创建分组是链接模式的使用aes算法的解密接口
	blockModel := cipher.NewCBCDecrypter(block, iv)
	//解密
	blockModel.CryptBlocks(sourceData, sourceData)
	//删除填充数据
	sourceData=public.CutFillData(sourceData)
	//返回数据
	return sourceData
}
