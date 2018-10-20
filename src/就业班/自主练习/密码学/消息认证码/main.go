package main

import (
	"crypto/sha512"
	"crypto/hmac"
	"fmt"
)

func main() {
	src := []byte("在消息认证码中，需要发送者和接收者之间共享密钥，而这个密钥不能被主动攻击者Mallory获取。如果这个密钥落入Mallory手中，则Mallory也可以计算出MAC值，从而就能够自由地进行篡改和伪装攻击，这样一来消息认证码就无法发挥作用了。")
	key := []byte("helloworld")
	hamc1 := GennerateHmac(src, key)
	bl := CheckHmac(src, key, hamc1)
	fmt.Printf("校验结果: %t\n", bl)
}

func GennerateHmac(sourceData, key []byte) []byte {
	//创建哈希接口
	myHash := hmac.New(sha512.New, key)
	//	//添加数据
	myHash.Write(sourceData)
	//	//生成散列值
	return myHash.Sum(nil)
}
func CheckHmac(socureData, key, hashMac []byte) bool {
	myHash := hmac.New(sha512.New, key)
	myHash.Write(socureData)
	resHash := myHash.Sum(nil)
	return hmac.Equal(resHash, hashMac)
}
