package main

import (
	"net"
	"fmt"
	"bytes"
	"os"
	"encoding/pem"
	"crypto/x509"
	"crypto/rsa"
	"crypto/rand"
	"crypto/aes"
	"crypto/cipher"
)

func main() {
	udpAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:3456")
	if err != nil {
		fmt.Println("ResolveUDPAddr:", err)
		return
	}
	conn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		fmt.Println("ListenUDP:", err)
		return
	}
	for {
		buf := make([]byte, 1024*4)
		n, Clientaddr, err := conn.ReadFromUDP(buf)
		if err != nil {
			fmt.Println("ReadFromUDP:", err)
			return
		}
		fmt.Printf("接收到来自客户端%v的信息%s", Clientaddr, string(buf[:n]))
		go func() {
			conn.WriteToUDP([]byte(""), Clientaddr)
		}()

	}
}


//公钥解密
func SevRsaPublicEncrypt(sourceData []byte, ClientpublcKey string) {

}

//公钥加密
func SevRsaPrivateDecrypt(sourceData []byte, publicKeyFilePath string) ([]byte, error) {
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
func SevAesEnCrypt(sourceData, key []byte) []byte {
	block, err := aes.NewCipher(key)
	msg := []byte("")
	if err != nil {
		fmt.Println("NewCipher:", err)
		return msg
	}
	//填充分组数据
	sourceData = SevFillLastGroupData(sourceData, block.BlockSize())
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
func SevFillLastGroupData(sourceData []byte, blockSize int) []byte {
	neeToFillDataLen := blockSize - len(sourceData)%blockSize
	//将neeToFillDataLen作为因子填入
	neeToFillDataLenText := bytes.Repeat([]byte{byte(neeToFillDataLen)}, neeToFillDataLen)
	//将最后一个添加的组拼接到源数据尾部
	NewData := append(sourceData, neeToFillDataLenText...)
	return NewData
}

//删除最后一个分组补位的数据
//sourceData   源
func SevCutFillData(sourceData []byte) []byte {
	//获取要切除切片的长度
	Cutlen := int(sourceData[len(sourceData)-1])
	//从源数据切除这个长度
	resSourceData := sourceData[:len(sourceData)-Cutlen]
	//返回最终的源数据
	return resSourceData
}
