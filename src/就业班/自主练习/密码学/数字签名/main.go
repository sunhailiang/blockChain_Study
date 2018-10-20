package main

import (
	"crypto/rsa"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"os"
	"crypto/sha256"
	"crypto"
	"fmt"
)

func main() {
	sourceData := []byte("这一份机密文件")
	signData, err := SignatureRSA(sourceData)
	if err != nil {
		fmt.Println("加密出错:", err)
	}
	err = VerifyRSA(sourceData, signData)
	if err != nil {
		fmt.Println("校验出错:", err)
	}
	fmt.Println("校验正确:")

}

func SignatureRSA(sourceData []byte) ([]byte, error) {
	msg := []byte("")
	//从文件读取私钥
	file, err := os.Open("E:/go/src/sign/privateKey.pem")
	if err != nil {
		return msg, err
	}
	info, err := file.Stat()
	if err != nil {
		return msg, err
	}
	buf := make([]byte, info.Size())
	file.Read(buf)
	//解析
	block, _ := pem.Decode(buf)
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return msg, err
	}
	//哈希加密
	myHash := sha256.New()
	myHash.Write(sourceData)
	hashRes := myHash.Sum(nil)
	//对哈希结果进行签名
	res, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hashRes)
	if err != nil {
		return msg, err
	}
	defer file.Close()
	return res, nil
}

func VerifyRSA(sourceData, signedData []byte) error {
	file, err := os.Open("E:/go/src/sign/publicKey.pem")
	if err != nil {
		return err
	}
	info, err := file.Stat()
	if err != nil {
		return err
	}
	buf := make([]byte, info.Size())
	file.Read(buf)
	//pem解密
	block, _ := pem.Decode(buf)
	publicInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return err
	}
	publicKey := publicInterface.(*rsa.PublicKey)
	//元数据哈希加密
	mySha := sha256.New()
	mySha.Write(sourceData)
	res := mySha.Sum(nil)

	//校验签名
	err = rsa.VerifyPKCS1v15(publicKey, crypto.SHA256, res, signedData)
	if err != nil {
		return err
	}
	defer file.Close()
	return nil

}

//生成密钥对
func GenerateRsaKey(bit int) error {
	private, err := rsa.GenerateKey(rand.Reader, bit)
	if err != nil {
		return err
	}
	//x509私钥序列化
	privateStream := x509.MarshalPKCS1PrivateKey(private)
	//将私钥设置到pem结构中
	block := pem.Block{
		Type:  "Rsa Private Key",
		Bytes: privateStream,
	}
	//保存磁盘
	file, err := os.Create("E:/go/src/sign/privateKey.pem")
	if err != nil {
		return err
	}
	//pem编码
	err = pem.Encode(file, &block)
	if err != nil {
		return err
	}
	//=========public=========
	public := private.PublicKey
	//509序列化
	publicStream, err := x509.MarshalPKIXPublicKey(&public)
	if err != nil {
		return err
	}
	//公钥赋值pem结构体
	pubblock := pem.Block{Type: "Rsa Public Key", Bytes: publicStream,}
	//保存磁盘
	pubfile, err := os.Create("E:/go/src/sign/publicKey.pem")
	if err != nil {
		return err
	}
	//pem编码
	err = pem.Encode(pubfile, &pubblock)
	if err != nil {
		return err
	}
	return nil

}
