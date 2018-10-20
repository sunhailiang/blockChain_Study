package main

import "fmt"

func main() {
desTest()
}

//测是des加解密
func desTest() {
	sourceData := []byte("黄河之水天上来，奔流到海不复还")
	key := []byte("abcdefgh")
	enRes := encryptDES(sourceData, key)
	fmt.Println("加密结果", enRes)
	deRes := decryptDES(enRes, key)
	fmt.Println("解密结果", string(deRes))
}
