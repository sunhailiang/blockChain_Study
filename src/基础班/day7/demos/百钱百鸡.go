package main

import "fmt"

//cock 公鸡 5钱
//hen 母鸡 3钱
//chicken 小鸡 1/3钱

func main() {
	//checken
	var checken int
	for i:=0;i<=20; i++{
		for j:=0;j<33;j++  {
			checken=100-i-j
			if i*5+j*3+checken/3==100&&checken%3==0 {
				fmt.Printf("公鸡：%d 母鸡：%d 小鸡：%d\n", i, j, checken)
			}
		}
	}
}
