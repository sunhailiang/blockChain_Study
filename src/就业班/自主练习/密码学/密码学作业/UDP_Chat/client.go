package main

import (
	"net"
	"fmt"
	"crypto/aes"
	"bytes"
	"crypto/cipher"
	"os"
	"io"
	"encoding/pem"
	"crypto/x509"
	"crypto/rsa"
	"crypto/rand"
)

type msg struct {
	publicKey string
	content   string
}

func main() {
	conn, err := net.Dial("udp", "127.0.0.1:3456")
	if err != nil {
		fmt.Println("Dial:", err)
		return
	}
	var msg msg
	buf := make([]byte, 4096)
	file, err := os.Open("E:/go/src/clientkey/privateKey.pem")
	for {
		fmt.Scanf("%v", &msg.content)
		if msg.content == "esc" {
			break
		}
		if err != nil {
			fmt.Println("clientPrivateKey os.Open:", err)
			return
		}
		for {
			_, err := file.Read(buf)
			if err == io.EOF {
				break
			}
			msg.publicKey += string(buf)
		}
		//数据加密
		key := []byte("ABCDEFGH12345678")
		publicKeyFilePath := "E:/go/src/clientkey/publicKey.pem"
		//对称加密
		AesRes := CliAesEnCrypt([]byte(msg.content), key)
		////非对称加密
		resSendData, err := CliRsaPrivateDecrypt(AesRes, publicKeyFilePath)
		if err != nil {
			fmt.Println("CliRsaPrivateDecrypt:", err)
			return
		}
		msg.content = string(resSendData)
		_, err = conn.Write([]byte("{key:" + msg.publicKey + ",content:" + msg.content + "}"))
		if err != nil {
			fmt.Println("conn.write", err)
			return
		}
		conn.Read(buf)
	}
	defer file.Close()
}

//公钥解密
func CliRsaPublicEncrypt(sourceData []byte, ClientpublcKey string) {

}

//公钥加密
func CliRsaPrivateDecrypt(sourceData []byte, publicKeyFilePath string) ([]byte, error) {
	resMsg := []byte("")
	//读取私钥
	file, err := os.Open(publicKeyFilePath)
	if err != nil {
		return resMsg, err
	}
	info, _ := file.Stat()
	buf := make([]byte, info.Size())
	file.Read(buf)
	//pem解密
	block, _ := pem.Decode(buf)
	//x509解密拿到私钥
	public, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return resMsg, err
	}
	//转换指针类型
	publicKey := public.(*rsa.PublicKey)
	//数据加密
	sourceData, err = rsa.EncryptPKCS1v15(rand.Reader, publicKey, sourceData)
	if err != nil {
		return resMsg, err
	}
	return sourceData, nil
}

//信息对称加密AES
func CliAesEnCrypt(sourceData, key []byte) []byte {
	block, err := aes.NewCipher(key)
	msg := []byte("")
	if err != nil {
		fmt.Println("NewCipher:", err)
		return msg
	}
	//填充分组数据
	sourceData = FillLastGroupData(sourceData, block.BlockSize())
	//创建asc加密模式
	blockModel := cipher.NewCBCEncrypter(block, key)
	//加密数据
	fmt.Println("sourceData", sourceData)
	blockModel.CryptBlocks(sourceData, sourceData)
	return sourceData
}

//信息对称解密AES
//func CliSevAesDeCrypt(sourceData, key []byte) []byte {
//
//}

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
