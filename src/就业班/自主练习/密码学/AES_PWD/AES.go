package main

import (
	"crypto/aes"
	"bytes"
	"crypto/cipher"
)

//AES加密
func EncryptAES(sourceData, key []byte) []byte {
	//1,创建一个返回使用AES密码块的接口
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	} //2,填充最后分组的数据
	sourceData = FillLastGroupData(sourceData, block.BlockSize())
	//3,创建密码分组为链接模式，底层使用AES加密算法的BlockModel
	blockModel := cipher.NewCBCEncrypter(block, key)
	//4.加密数据
	blockModel.CryptBlocks(sourceData, sourceData)
	//返回加密结果
	return sourceData
}

//AES解密
func DecryptAES(sourceData, key []byte) []byte {
	//1,创建并返回一个使用AES算法的cipher.Block接口
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	//2，创建一个密码分组为链接模式的底层使用AES解密的BlockModel接口
	blockModel := cipher.NewCBCDecrypter(block, key)
	//解密
	blockModel.CryptBlocks(sourceData, sourceData)
	//去掉填充数据
	sourceData = CutFillData(sourceData)
	return sourceData
}

//================================================
//填充最后一个分组的函数
//sourceData   源
//blockSize    每个分组数据长度
func FillLastGroupData(sourceData []byte, blockSize int) []byte {
	neeToFillDataLen := blockSize - len(sourceData)%blockSize
	//将neeToFillDataLen作为因子填入
	neeToFillDataLenText := bytes.Repeat([]byte{byte(neeToFillDataLen)}, neeToFillDataLen)
	//将最后一个添加的组拼接到源数据尾部
	NewData := append(sourceData, neeToFillDataLenText...)
	return NewData
}

//删除最后一个分组补位的数据
//sourceData   源
func CutFillData(sourceData []byte) []byte {
	//获取要切除切片的长度
	Cutlen := int(sourceData[len(sourceData)-1])
	//从源数据切除这个长度
	resSourceData := sourceData[:len(sourceData)-Cutlen]
	//返回最终的源数据
	return resSourceData
}
