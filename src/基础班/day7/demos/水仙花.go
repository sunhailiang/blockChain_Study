package main

import "fmt"

//水仙花  一个三位数 100-999 各个位数的立方和等于这个数本身 就是一个水仙花数
func main() {
	getNums()
}

func getNums() {
	for i := 100; i < 999; i++ {
        var num1=i/100
        var num2=i%100/10
        var num3=i%10
        if num1*num1*num1+num2*num2*num2+num3*num3*num3==i{
        	fmt.Println(i)
		}
	}
}
