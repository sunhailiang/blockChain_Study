package main

import "fmt"

//水仙花  一个三位数 100-999 各个位数的立方和等于这个数本身 就是一个水仙花数
func main02() {
	for i:=100; i<1000;i++  {
		//百位
		a:=i/100
		//十位
		b:=i/10%10//b:=i%100/10
		//个位
		c:=i%10
		if a*a*a+b*b*b+c*c*c==i{
			fmt.Println(i)
		}
	}
}
