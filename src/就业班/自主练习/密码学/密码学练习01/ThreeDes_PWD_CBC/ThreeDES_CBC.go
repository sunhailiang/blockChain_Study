package threeDES_CBC

import (
	"crypto/des"
	"就业班/自主练习/密码学/密码学练习01/public"
	"crypto/cipher"
)

func ThreeDESEnCrypt(sourceData, key []byte) []byte {
	//使用3des算法创建加密块
	block, err := des.NewTripleDESCipher(key)
	if err != nil {
		panic(err)
	}
	//填充数据
	sourceData = public.FillLastGroup(sourceData, block.BlockSize())
	//cbc分组
	iv := []byte("1234567887654332112345678")
	blockModel := cipher.NewCBCEncrypter(block, iv)
	//加密
	blockModel.CryptBlocks(sourceData, sourceData)
	return sourceData
}

func ThreeDESDeCrypt(sourceData, key []byte) []byte {
	//创建基于3des的加密块
	block, err := des.NewTripleDESCipher(key)
	if err != nil {
		panic(err)
	}
	//创建基于cbc链接模式的解密块
	iv := []byte("1234567887654332112345678")
	blockModel := cipher.NewCBCDecrypter(block, iv)
	//解密
	blockModel.CryptBlocks(sourceData, sourceData)
	//剔除补充数据
	sourceData = public.CutFillData(sourceData)
	return sourceData
}
