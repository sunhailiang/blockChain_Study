package main

import "fmt"

//括号 （） 结构体成员.  数组下标[]

//单目运算符
//逻辑非! 取地址& 取值* 自增++ 自减--

//双目运算符
//乘除 * / %
//加减 + -
//关系 == != > >= < <=
//逻辑 || &&
//赋值 = += -= *= /= %=
func main() {
	a:=10
	b:=20
	c:=30
	//d:=a+b*c
	//var d int
	//d=(a+b)*c


	//fmt.Println(d)
	fmt.Println(a+b>=c && !(b>c))
}
