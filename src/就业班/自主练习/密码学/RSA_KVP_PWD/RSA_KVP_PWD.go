package main

import (
	"crypto/rsa"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"os"
)

//生成私/公钥并写入磁盘
func RasGenKey(bit int) error {
	//1,使用rsa中的GenerateKey方法生成私钥
	privateKey, err := rsa.GenerateKey(rand.Reader, bit)//bit：1024的倍
	if err != nil {
		return err
	}
	//2,通过x509标准将得到的ras私钥序列化为ASN.1的DER编码字符串
	privateStream := x509.MarshalPKCS1PrivateKey(privateKey)
	//3,将私钥串设置到pem块中
	block := pem.Block{
		Type:  "RAS Private Key",
		Bytes: privateStream,
	}
	//4,将pem文件写入磁盘保存
	privateFile, err := os.Create("private.pem")
	if err != nil {
		return err
	}
	err = pem.Encode(privateFile, &block)
	if err != nil {
		return err
	}

	//=============生成公钥=============
	//1，从私钥中提取公钥信息
	publicKey := privateKey.PublicKey
	//2，通过x509标准得到rsa公钥序列化字符串
	publicKeyStream, err := x509.MarshalPKIXPublicKey(&publicKey) //此处注意，必须取地址才行
	if err != nil {
		return err
	}
	//3,将公钥字符串设置到pem格式中
	publicBlock := pem.Block{
		Type:  "RSA Public Key",
		Bytes: publicKeyStream,
	}
	//4,通过pem将设置好的数据编码并写入磁盘
	publicFile, err := os.Create("publicKey.pem")
	if err != nil {
		return err
	}
	err = pem.Encode(publicFile, &publicBlock)
	if err != nil {
		return err
	}
	//获取完毕之后关闭文件
	defer func() {
		privateFile.Close()
		publicFile.Close()
	}()
	return nil
}

//===========公钥加密==============
func RsaPublicEncrypt(sourceData []byte, publicKeyFilePath string) ([]byte, error) {
	msg := []byte("")
	//1,根据路径读出公钥,获取pem字符串
	publickeyFile, err := os.Open(publicKeyFilePath)
	if err != nil {
		return msg, err
	}
	//1.1获取文件文件大小
	info, err := publickeyFile.Stat()
	if err != nil {
		return msg, err
	}
	//开辟缓冲区
	buf := make([]byte, info.Size())
	publickeyFile.Read(buf)
	//2，得到的字符串解码
	block, _ := pem.Decode(buf)
	//3，使用x509标准将公钥解析出来
	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return msg, err
	}
	publicKey := pub.(*rsa.PublicKey) //转成实际类型指针
	//4，使用公钥通过rsa进行数据加密
	msg, err = rsa.EncryptPKCS1v15(rand.Reader, publicKey, sourceData)
	if err != nil {
		return msg, err
	}
	return msg, nil
}

//===========私钥解密==============
func RsaPrivateDecrypt(sourceData []byte, privateKeyFilePath string) ([]byte, error) {
	msg := []byte("")
	//1,打开私钥文件
	file, err := os.Open(privateKeyFilePath)
	if err != nil {
		return msg, err
	}
	//2，读取私钥文件内容
	fileInfo, err := file.Stat()
	if err != nil {
		return msg, err
	}
	buf := make([]byte, fileInfo.Size())
	file.Read(buf)
	//3,将读出的私钥字符串进行解码
	block, _ := pem.Decode(buf)
	//4,通过x509协议还原私钥数据
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return msg, err
	}
	//5，通过密钥对数据解密
	msg, err = rsa.DecryptPKCS1v15(rand.Reader, privateKey, sourceData)
	if err != nil {
		return msg, err
	}
	return msg, nil

}
