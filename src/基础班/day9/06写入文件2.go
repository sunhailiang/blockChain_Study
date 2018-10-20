package main

import (
	"os"
	"fmt"
)

func main0601() {
	//var fp *File
	//os.Create 创建文件时 文件不存在会创建一个新文件  如果文件存在会覆盖源内容
	fp,err:=os.Create("D:/a.txt")
	if err!=nil{
		fmt.Println("文件创建失败")
		return
	}
	defer fp.Close()

	//写入文件
	//\n在文件中是换行  window文本文件换行是\r\n
	//n,_:=fp.WriteString("itcast\r\n")

	//fmt.Println(n)
	//在go语言中 一个汉字是3个字符
	n,_:=fp.WriteString("美女")
	fmt.Println(n)
}
func main(){
	//创建文件
	fp,err:=os.Create("D:/a.txt")
	if err!=nil{
		fmt.Println("文件创建失败")
		return
	}
	//关闭文件
	defer fp.Close()
	//1、将字符切片写入文件中
	//b:=[]byte{'h','e','l','l','o'}
	//
	////文件对象.write(字符切片)
	//n,_:=fp.Write(b)
	//fmt.Println(n)

	//2、将字符串转成字符切片写入文件中
	//str:="helloworld"
	str:="锄禾日当午"
	//字符串和字符切片允许转换
	b:=[]byte(str)

	fp.Write(b)
}