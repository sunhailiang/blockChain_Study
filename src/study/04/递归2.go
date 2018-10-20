package main

import "fmt"

func main(){
	s:=test(3)//3+2+1
	fmt.Println(s)
}

func test(n int)int {
	if n==1{
		return 1
	}
	fmt.Println("what？",n)
	return n+test(n-1)
}
