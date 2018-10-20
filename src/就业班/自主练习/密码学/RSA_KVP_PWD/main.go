package main

import "fmt"

func main() {
	test()
}

func test() {
	sourceData := []byte("这是一段没有结局的表演，包含所有荒谬和疯狂~")
	res, err := RsaPublicEncrypt(sourceData, "publicKey.pem")
	if err != nil {
		fmt.Println("加密出错", err)
		return
	}
	fmt.Println("加密数据", res)

	sData, err := RsaPrivateDecrypt(res, "private.pem")
	if err != nil {
		fmt.Println("解密出错", err)
		return
	}
	fmt.Println("解密结果", string(sData))
}
