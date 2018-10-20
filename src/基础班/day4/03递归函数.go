package main

import "fmt"

//在函数定义时调用函数本身 递归函数
//死递归
func test3(a int){
	//在递归函数中一个要有出口  return
	if a==0{
		return
	}
	a--
	fmt.Println(a)
	test3(a)
}

func main0301() {

	test3(10)
}


//计算n的阶乘
var sum int = 1
func test4(n int){
	if n==1{
		return
	}

	test4(n-1)
	sum*=n
}

func main(){
	test4(3)

	fmt.Println(sum)
}