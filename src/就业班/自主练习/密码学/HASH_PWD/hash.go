package main

import (
	"encoding/hex"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
)

//====================MD5=================
//同样是MD5的哈希算法，这种方式只能添加一次加密数据
func GetMD5Str(sourceData []byte) string {
	//加密数据
	md5Res := md5.Sum(sourceData)
	//格式化数据
	resStr := hex.EncodeToString(md5Res[:])
	return resStr
}

//同样是MD5的哈希算法，这种方式只能添加多次加密数据
func GetMD5Str_B(sourceData []byte) string {
	//1，创建哈希接口
	myHash := md5.New()
	//2,添加数据(方式一)
	//io.WriteString(myHash,string(sourceData))
	//2，添加数据（方式二）
	myHash.Write(sourceData)
	//3,获取加密结果
	res := myHash.Sum(nil) ////此处可以额外添加数据，合并加密结果
	//4，散列格式化
	return hex.EncodeToString(res)
}

//=====================SHA1======================
//SHA1加密方式(单一加密场景应用)-----------------已被攻破不安全的加密方式
func GetSha1Str(sourceData []byte) string {
	//1，SHA1加密
	shaRes := sha1.Sum(sourceData)
	//2，格式化数据
	resStr := hex.EncodeToString(shaRes[:])
	return resStr
}
func GetSha1Str_B(sourceData []byte) string {
	//创建哈希接口
	mySha := sha1.New()
	//添加数据
	mySha.Write(sourceData)
	//获取加密结果
	res := mySha.Sum(nil) //此处可以额外添加数据，合并加密结果
	//散列格式化
	return hex.EncodeToString(res)
}

//===============SHA2==================
//SHA2[SHA512]----尚未被破解，安全的加密方式
//一，单次加密应用场景
func GetSha256Str(sourceData []byte) string {
	//使用sha256加密算法直接对数据加密
	sha256Res := sha256.Sum256(sourceData)
	//散列格式化
	return hex.EncodeToString(sha256Res[:])
}

//二，多次加密应用场景
func GetSha256Str_B(sourceData []byte) string {
	//创建hash接口
	mySha256 := sha256.New()
	//添加数据
	mySha256.Write(sourceData)
	//加密并返回结果
	res := mySha256.Sum(nil) //此处可以额外添加数据，合并加密结果
	//散列格式化
	return hex.EncodeToString(res)
}

//SHA512哈希加密方法单次加密应用场景
func GetSha512Str(sourceData []byte) string {
	//直接加密数据
	res := sha512.Sum512(sourceData)
	//散列格式化，并返回数据
	return hex.EncodeToString(res[:])
}

func GetSha512Str_B(sourceData []byte) string {
	//创建哈希接口
	mySha512 := sha512.New()
	//添加数据
	mySha512.Write(sourceData)
	//加密数据并返回结果
	res := mySha512.Sum(nil) //此处可以额外添加数据，合并加密结果
	return hex.EncodeToString(res)
}
