package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"log"
	"crypto/sha256"
	"fmt"
	"github.com/base58"
	"github.com/ripemd160"
)

type Wallet struct {
	//首字母一定要大写，否则后面gob编码时会出错(用于保存)
	PrivateKey ecdsa.PrivateKey
	//由两个坐标点拼接而成的临时公钥，便于传输，校验时进行拆分，还原成原始的公钥
	PublicKey []byte
}

//创建钱包
func NewWallet() *Wallet {
	//生成私钥
	curve := elliptic.P256()
	privateKey, err := ecdsa.GenerateKey(curve, rand.Reader)
	if err != nil {
		log.Panic("私钥创建失败：", err)
	}
	//生成公钥
	rawPublicKey := privateKey.PublicKey
	//拆分公钥
	publicKey := append(rawPublicKey.X.Bytes(), rawPublicKey.Y.Bytes()...)
	return &Wallet{*privateKey, publicKey}
}

//============使用公钥生成地址============
func (w *Wallet) getAddress() string {
	//对公钥进行哈希处理：RIPEMD160(sha256())
	ripemdHash := HashPublicKey(w.PublicKey)
	version := byte(1)
	payload := append([]byte{version}, ripemdHash[:]...)
	//	//获取校验码
	checkCode := Checksum(payload)
	//拼接version + hash + checksum
	pubKeyHash := append(payload, checkCode...)
	fmt.Println("pubKeyHash : %x\n", pubKeyHash)

	//base58
	address := base58.Encode(pubKeyHash)
	return address
}

//获取校验码: checksum()
func Checksum(payload []byte) []byte {
	//第一次加密
	hash1 := sha256.Sum256(payload)
	//第二次加密
	hash2 := sha256.Sum256(hash1[:])
	//取前四个字符作为校验码
	chekCode := hash2[:4]
	return chekCode
}

//公钥哈希处理
func HashPublicKey(publicKey []byte) []byte {
	//先用sha2加密
	hash256 := sha256.Sum256(publicKey)
	//再用ripemd160加密
	ripemd160Hasher := ripemd160.New()
	_, err := ripemd160Hasher.Write(hash256[:])
	if err != nil {
		log.Panic("ripemd160 加密公钥出错", err)
	}
	ripemd160Hash := ripemd160Hasher.Sum(nil)
	return ripemd160Hash

}
