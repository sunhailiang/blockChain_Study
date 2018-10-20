package main

import "fmt"

func main() {
	AES_Test()
}
func AES_Test()  {
	sourceData:=[]byte("这是一段没有解决的表演，包含所有荒谬和疯狂~")
	key:=[]byte("qwertyuiasdfghjk")
	encipherRes:=EncryptAES(sourceData,key)
	fmt.Println("加密结果:",encipherRes)
	decipherRes:=DecryptAES(encipherRes,key)
	fmt.Println("解密结果:",string(decipherRes))
}