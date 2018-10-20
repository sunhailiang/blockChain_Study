package main

import (
	"就业班/自主练习/密码学/密码学练习01/DES_PWD_CBC"
	"fmt"
	"就业班/自主练习/密码学/密码学练习01/AES_PWD_CBC"
	"就业班/自主练习/密码学/密码学练习01/AES_PWD_CTR"
	"就业班/自主练习/密码学/密码学练习01/DES_PWD_CTR"
	"就业班/自主练习/密码学/密码学练习01/RSA_KVP_PWD"
)

func main() {
	//Aes_cbc()
	//desCbc()
	//aes_ctr()
	//des_ctr()
	kvpTest()
}

func Aes_cbc() {
	sourceData := []byte("这是一段没有解决的表演，包含所有荒谬和疯狂")
	key := []byte("abcdefgh12345678")
	enRes := aes.AesEnCrypt(sourceData, key)
	//aes算法，cbc分组模式
	fmt.Println("aes,cbc分组加密结果", enRes)
	deRes := aes.AesDeCrypt(enRes, key)
	fmt.Println("aes,cbc分组解密结果", string(deRes))
}

func desCbc() {
	sourceData := []byte("不要拘谨，MMP~")
	key := []byte("12345678")
	enRes := des_cbc.DesEnCrypt(sourceData, key)
	fmt.Println("des,cbc分组加密结果", enRes)
	deRes := des_cbc.DesDeCrypt(enRes, key)
	fmt.Println("des,cbc分组加密结果", string(deRes))
}

func aes_ctr() {
	sourceData := []byte("不要拘谨，MMP~")
	key := []byte("12345678ABCDEFGH")
	enRes := aesctr.AesEnCrypt(sourceData, key)
	fmt.Println("aes,ctr分组加密结果", enRes)
	deRes := aesctr.AesEnCrypt(enRes, key)
	fmt.Println("aes,ctr分组加密结果", string(deRes))
}

func des_ctr() {
	sourceData := []byte("不要拘谨")
	key := []byte("12345678")
	enRes := desctr.DesEnCrypt(sourceData, key)
	fmt.Println("des,ctr分组加密结果", enRes)
	deRes := desctr.DesDeCrypt(enRes, key)
	fmt.Println("des,ctr分组加密结果", string(deRes))
}

func kvpTest() {
	//kvp.RasGenKey(6144)
	sourceData := []byte("AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaabbbbbbbbbbbbbbbbbaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaabbbbbbbbbbbbbbbbbaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaabbbbbbbbbbbbbbbbbaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaabbbbbbbbbbbbbbbbb")
	publicKeyFilePath := "public.pem"
	privateKeyFilePath := "private.pem"
	enRes, err := kvp.RsaPublicEncrypt(sourceData, publicKeyFilePath)
	if err != nil {
		fmt.Println("加密出错:", err)
		return
	}
	fmt.Println("加密结果:", enRes)

	txt, err := kvp.RsaPrivateDecrypt(enRes, privateKeyFilePath)
	if err != nil {
		fmt.Println("解密出错:", err)
		return
	}
	fmt.Println("解密结果:", string(txt))

}
