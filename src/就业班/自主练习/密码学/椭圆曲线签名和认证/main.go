package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"os"
	"crypto/sha1"
	"math/big"
	"fmt"
)

func main() {
	privateKeyFilePath := "E:/go/src/ecc/eccprivateKey.pem"
	publicKeyFilePath := "E:/go/src/ecc/eccpublicKey.pem"
	sourceData := []byte("测试加密数据")
	r, s := EccSignature(sourceData, privateKeyFilePath)
	res := EccVerify(r, s, sourceData, publicKeyFilePath)
	if res {
		fmt.Println("认证成功")
	}
}

//生成密钥对
func GenerateEccKey() error {
	//使用ecdsa生成密钥对
	privateKey, err := ecdsa.GenerateKey(elliptic.P521(), rand.Reader)
	if err != nil {
		return err
	}
	//使用509
	private, err := x509.MarshalECPrivateKey(privateKey) //此处
	if err != nil {
		return err
	}
	//pem
	block := pem.Block{
		Type:  "esdsa private key",
		Bytes: private,
	}
	file, err := os.Create("E:/go/src/ecc/eccprivateKey.pem")
	if err != nil {
		return err
	}
	err = pem.Encode(file, &block)
	if err != nil {
		return err
	}
	file.Close()

	//处理公钥
	public := privateKey.PublicKey

	//x509序列化
	publicKey, err := x509.MarshalPKIXPublicKey(&public)
	if err != nil {
		return err
	}
	//pem
	public_block := pem.Block{
		Type:  "ecdsa public key",
		Bytes: publicKey,
	}
	file, err = os.Create("E:/go/src/ecc/eccpublicKey.pem")
	if err != nil {
		return err
	}
	//pem编码
	err = pem.Encode(file, &public_block)
	if err != nil {
		return err
	}
	return nil
}

//ecc签名--私钥
func EccSignature(sourceData []byte, privateKeyFilePath string) ([]byte, []byte) {
	//1，打开私钥文件，读出内容
	file, err := os.Open(privateKeyFilePath)
	if err != nil {
		panic(err)
	}
	info, err := file.Stat()
	buf := make([]byte, info.Size())
	file.Read(buf)
	//2,pem解密
	block, _ := pem.Decode(buf)
	//x509解密
	privateKey, err := x509.ParseECPrivateKey(block.Bytes)
	if err != nil {
		panic(err)
	}
	//哈希运算
	hashText := sha1.Sum(sourceData)
	//数字签名
	r, s, err := ecdsa.Sign(rand.Reader, privateKey, hashText[:])
	if err != nil {
		panic(err)
	}
	rText, err := r.MarshalText()
	if err != nil {
		panic(err)
	}
	sText, err := s.MarshalText()
	if err != nil {
		panic(err)
	}
	defer file.Close()
	return rText, sText
}

//ecc认证

func EccVerify(rText, sText, sourceData []byte, publicKeyFilePath string) bool {
	//读取公钥文件
	file, err := os.Open(publicKeyFilePath)
	if err != nil {
		panic(err)
	}
	info, err := file.Stat()
	if err != nil {
		panic(err)
	}
	buf := make([]byte, info.Size())
	file.Read(buf)
	//pem解码
	block, _ := pem.Decode(buf)

	//x509
	publicStream, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		panic(err)
	}
	//接口转换成公钥
	publicKey := publicStream.(*ecdsa.PublicKey)
	hashText := sha1.Sum(sourceData)
	var r, s big.Int
	r.UnmarshalText(rText)
	s.UnmarshalText(sText)
	//认证
	res := ecdsa.Verify(publicKey, hashText[:], &r, &s)
	defer file.Close()
	return res
}
