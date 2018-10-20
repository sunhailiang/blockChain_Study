package main

import (
	"strconv"
	"fmt"
)

func main0201() {
	//1、将其他类型转成字符串
	//将bool类型转成字符串
	//s:=strconv.FormatBool(false)
	//fmt.Println(s)

	//2、将整型转成字符串
	//formatInt(数据，进制) 二进制  八进制  十进制 十六进
	//s:=strconv.FormatInt(123,10)
	//s:=strconv.Itoa(123456)
	//fmt.Println(s)

	//3、浮点型转成字符串
	//formatfloat(数据,'f'，保留小数位置（谁四舍五入）,位数（64 32）)
	s:=strconv.FormatFloat(1.155,'f',2,64)
	fmt.Println(s)
}

func main0202(){
	//1、将字符串转成bool
	//str:="true"
	////只能将“flase”“true”转成bool类型 有多余的数据 是无效的
	//b,err:=strconv.ParseBool(str)
	//if err!=nil{
	//	fmt.Println(err)
	//}else {
	//	fmt.Println(b)
	//}

	//2、将字符串转成整型
	//str:="1234"
	//value,_:=strconv.ParseInt(str,10,64)
	//fmt.Println(value)
	//str:="123"
	//value,_:=strconv.Atoi(str)
	//fmt.Println(value)

	//3、将字符串转成浮点型

	//str:="12345"
	//
	//value,_:=strconv.ParseFloat(str,64)
	//fmt.Println(value)

}

func main(){

	b:=make([]byte,0,1024)
	//将bool类型放在指定切片中
	b=strconv.AppendBool(b,false)
	b=strconv.AppendInt(b,123,10)
	b=strconv.AppendFloat(b,1.234,
		'f',5,64)
	b=strconv.AppendQuote(b,"hello")
	fmt.Println(string(b))
}