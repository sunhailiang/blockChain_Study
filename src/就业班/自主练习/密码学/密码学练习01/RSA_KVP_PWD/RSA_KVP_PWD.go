package kvp

import (
	"crypto/rsa"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"os"
)

//生成公钥私钥
func RasGenKey(bit int) error {
	//1,使用Ras的GennerateKey生成私钥
	PrivateKey, err := rsa.GenerateKey(rand.Reader, bit)
	if err != nil {
		return err
	}
	//使用x509协议进行进行私钥序列化
	privateStream := x509.MarshalPKCS1PrivateKey(PrivateKey)
	//使用pem编码
	block := pem.Block{
		Type:  "private key",
		Bytes: privateStream,
	}
	//写入磁盘
	privateFile, err := os.Create("private.pem")
	if err != nil {
		return err
	}
	//进行文件编码
	err = pem.Encode(privateFile, &block)
	if err != nil {
		return err
	}
	//=====获取公钥=======
	publicKey := PrivateKey.PublicKey
	//使用x509协议进行公钥序列化
	publicStream, err := x509.MarshalPKIXPublicKey(&publicKey)
	if err != nil {
		return err
	}
	//使用pem编码
	publicbBlock := pem.Block{
		Type:  "RSA public key",
		Bytes: publicStream,
	}
	//写入磁盘
	publicFile, err := os.Create("public.pem")
	if err != nil {
		return err
	}
	//pem文件编码
	err = pem.Encode(publicFile, &publicbBlock)
	if err != nil {
		return err
	}
	return nil
}

//公钥加密
func RsaPublicEncrypt(sourceData []byte, publicKeyFilePath string) ([]byte, error) {
	msg := []byte("")
	//打开公钥文件
	file, err := os.Open(publicKeyFilePath)
	if err != nil {
		return msg, err
	}
	//1.1获取文件大小
	info, err := file.Stat()
	if err != nil {
		return msg, err
	}
	//读取内容
	// 开辟缓冲区
	buf := make([]byte, info.Size())
	file.Read(buf)
	//pem解码
	block, _ := pem.Decode(buf)
	//x509解码
	public, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return msg, err
	}
	//转换指针类型
	publicKey := public.(*rsa.PublicKey)
	//rsa数据加密
	msg, err = rsa.EncryptPKCS1v15(rand.Reader, publicKey, sourceData)
	if err != nil {
		return msg, err
	}
	return msg, err
}

func RsaPrivateDecrypt(sourceData []byte, publicKeyFilePath string) ([]byte, error) {
	msg := []byte("")
	//获取获取私钥文件并取出来
	file, err := os.Open(publicKeyFilePath)
	if err != nil {
		return msg, err
	}
	//1.1获取文件信息
	info, err := file.Stat()
	if err != nil {
		return msg, err
	}
	buf := make([]byte, info.Size())
	file.Read(buf)
	//pem解密
	block, _ := pem.Decode(buf)
	//x509解密获取privateKey
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return msg, err
	}
	//使用privateKey对数据进行加密
	msg, err = rsa.DecryptPKCS1v15(rand.Reader, privateKey, sourceData)
	if err != nil {
		return msg, err
	}
	return msg, err
}
