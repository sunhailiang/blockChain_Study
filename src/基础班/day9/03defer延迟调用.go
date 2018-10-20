package main

import "fmt"

func main0301() {
	fmt.Println("传智播客")
	defer fmt.Println("IT培训")
	//defer 函数调用
	defer fmt.Println("日薪越亿")
	fmt.Println("高薪就业")
}

func test4() {
	fmt.Println("hello world")

}

var value int

func test5(a int, b int) {
	value = a + b
}

func test6(value int) {
	fmt.Println(value)
}
func main0302() {

	test4()
	//函数中有返回值不能使用defer调用
	defer test5(10, 20)
	test6(value)

}

func main() {
	a := 10
	b := 20
	//定义函数类型变量
	//f:=func(a int ,b int){
	//	fmt.Println(a)
	//	fmt.Println(b)
	//}

	//f(a,b)

	defer func(a int ,b int) {
		fmt.Println(a)
		fmt.Println(b)
	}(a,b)

	a = 100
	b = 200

	fmt.Println(a)
	fm