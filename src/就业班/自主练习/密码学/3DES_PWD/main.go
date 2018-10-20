package main

import "fmt"

func main() {
	ThreeDesTest()
}

func ThreeDesTest() {
	sourceDate := []byte("你烫过的衣服仿佛还有温度，却抵挡不住我我寒冷孤独~")
	key := []byte("abcdefgh12345678abcdefgh")
	EnsourceDate := Encrypt3DES(sourceDate, key)
	fmt.Println("加密结果", EnsourceDate)
	res := Decrypt3DES(EnsourceDate, key)
	fmt.Println("解密结果", string(res))
}
