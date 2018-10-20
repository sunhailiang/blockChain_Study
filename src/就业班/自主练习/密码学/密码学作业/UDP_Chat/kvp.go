package main

import (
	"crypto/rsa"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"os"
)

func main() {
	RsaGenKey(4096)
}
func RsaGenKey(bit int) error {
	private, err := rsa.GenerateKey(rand.Reader, bit)
	if err != nil {
		return err
	}
	//使用509协议进行私钥序列化
	privateKey := x509.MarshalPKCS1PrivateKey(private)
	//pem编码
	block := pem.Block{
		Type:  "private Key RSA",
		Bytes: privateKey,
	}
	//写入磁盘
	file, err := os.Create("E:/go/src/clientkey/privateKey.pem")
	if err != nil {
		return err
	}

	err = pem.Encode(file, &block)
	if err != nil {
		return err
	}
	//================获取公钥================
	public := private.PublicKey
	//x509协议序列化
	publicKey, err := x509.MarshalPKIXPublicKey(&public)
	if err != nil {
		return err
	}
	//pem编码
	pubblock := pem.Block{
		Type:  "public key rsa",
		Bytes: publicKey,
	}
	//写入磁盘
	publicKeyFile, err := os.Create("E:/go/src/clientkey/publicKey.pem")
	if err != nil {
		return err
	}
	err = pem.Encode(publicKeyFile, &pubblock)
	if err != nil {
		return err
	}
	return nil
}
