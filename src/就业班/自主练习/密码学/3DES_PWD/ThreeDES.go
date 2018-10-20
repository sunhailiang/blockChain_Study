package main

import (
	"bytes"
	"crypto/des"
	"crypto/cipher"
)

//CBC分组模式
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

//3DES对称加密
func Encrypt3DES(sourceData, key []byte) []byte {
	//1,创建并返回一个使用3DES算法的cipher.Blockj接口
	block, err := des.NewTripleDESCipher(key)
	if err != nil {
		panic(err)
	}
	//2,最后一个明文分组数据填充
	sourceData = FillLastGroupData(sourceData, block.BlockSize())
	//3,创建一个密码分组为链接模式的，底层使用3DES加密的BlockModelj接口
	blockModel := cipher.NewCBCEncrypter(block, key[:block.BlockSize()]) //此处的向量采用使用key的方式
	//4,加密连续数据块
	blockModel.CryptBlocks(sourceData, sourceData)
	return sourceData
}

//3DES解密
func Decrypt3DES(sourceData, key []byte) []byte {
	//1,创建并返回一个使用3DES算法的cipher.Block接口
	block, err := des.NewTripleDESCipher(key)
	if err != nil {
		panic(err)
	}
	//2，创建一个密码分组为链接模式，底层使用3DES解密的blocModel接口
	blockModel := cipher.NewCBCDecrypter(block, key[:block.BlockSize()])
	//3,数据块解密
	blockModel.CryptBlocks(sourceData, sourceData)
	//4，去掉填充数据
	res := CutFillData(sourceData)
	return res
}
