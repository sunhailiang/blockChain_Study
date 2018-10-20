package main

import (
	"bytes"
	"crypto/des"
	"crypto/cipher"
)
//======================对称加密DES==================

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

//DES对称加密key:就是密钥
func encryptDES(sourceData, key []byte) []byte {
	//  1，创建并返回一个使用DES算法的cipher.Block接口
	//  key就是初始化的一个向量
	block, err := des.NewCipher(key)
	if err != nil {
		panic(err)
	}
	//2,填充最后一个明文分组
	sourceData=FillLastGroupData(sourceData, block.BlockSize())
	//3，创建一个密码分组为链接模式，底层使用DES加密的BlocMode接口
	iv := []byte("aaaabbbb") //正式环境中发送和接收者需要协商iv使用key作为iv能避免这个操作
	blockModel := cipher.NewCBCEncrypter(block, iv)
	//4,连续加密数据块
	dst := make([]byte, len(sourceData))
	blockModel.CryptBlocks(dst, sourceData)
	return dst //返回加密结果

}

//DES解密
func decryptDES(sourceData, key []byte) []byte {
	//1,创建并返回一个使用DES算法的cipher.Block
	block, err := des.NewCipher(key)
	if err != nil {
		panic(err)
	}
	//2,创建一个密码分组为链接模式的，底层使用DES解密的BlockModel接口
	iv := []byte("aaaabbbb")
	blockModel := cipher.NewCBCDecrypter(block, iv)
	//3,数据块解密
	blockModel.CryptBlocks(sourceData, sourceData)
	//4,剔除最后一组填充数据
	resSourceData := CutFillData(sourceData)
	return resSourceData
}
