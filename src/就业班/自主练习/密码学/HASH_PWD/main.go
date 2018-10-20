package main

import "fmt"

func main() {
	testMD5()
	testSHA1()
	testSHA256()
	testSHA512()
}
func testMD5() {
	sourceData := "这是一段没有结局的表演~包含所有荒谬和疯狂~MD5"
	resA := GetMD5Str([]byte(sourceData))
	fmt.Println("加密MD5_A", resA)
	resB := GetMD5Str_B([]byte(sourceData))
	fmt.Println("加密MD5_B", resB)
}

func testSHA1() {
	sourceData := "这是一段没有结局的表演~包含所有荒谬和疯狂~SHA1"
	resA := GetSha1Str([]byte(sourceData))
	fmt.Println("加密SHA1", resA)
	resB := GetSha1Str_B([]byte(sourceData))
	fmt.Println("加密SHA1_B", resB)
}

func testSHA256() {
	sourceData := "这是一段没有结局的表演~包含所有荒谬和疯狂~SHA256"
	resA := GetSha256Str([]byte(sourceData))
	fmt.Println("加密Sha256", resA)
	resB := GetSha256Str_B([]byte(sourceData))
	fmt.Println("加密Sha256_B", resB)
}
func testSHA512() {
	sourceData := "这是一段没有结局的表演~包含所有荒谬和疯狂~Sha512"
	resA := GetSha512Str([]byte(sourceData))
	fmt.Println("加密Sha512Str", resA)
	resB := GetSha512Str_B([]byte(sourceData))
	fmt.Println("加密Sha512Str_B", resB)
}
