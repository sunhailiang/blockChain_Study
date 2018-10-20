package des_cbc

import (
	"crypto/des"
	"就业班/自主练习/密码学/密码学练习01/public"
	"crypto/cipher"
)

//desSUANF
func DesEnCrypt(sourceData, key []byte) []byte {
	//创建基于des的加密对象
	block, err := des.NewCipher(key)
	if err != nil {
		panic(err)
	}
	//填充数据
	sourceData = public.FillLastGroup(sourceData, block.BlockSize())
	//选择分组模式
	iv := []byte("12345678")
	blockModel := cipher.NewCBCEncrypter(block, iv)
	//加密
	blockModel.CryptBlocks(sourceData, sourceData)
	return sourceData

}
func DesDeCrypt(sourceData, key []byte) []byte {
	//基于des算法的加密对象
	block, err := des.NewCipher(key)
	if err != nil {
		panic(err)
	}
	//创建基于cbc分组模式的链对象模型
	iv := []byte("12345678")
	blockModel := cipher.NewCBCDecrypter(block, iv)
	//解密
	blockModel.CryptBlocks(sourceData, sourceData)
	return sourceData
}
