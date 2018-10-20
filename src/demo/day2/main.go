package main

import "fmt"

func list(n int){
	for i:=0;i<n;i++{
		a:=(n-i);
		b:=i
		fmt.Println("n",n,"a",a,"b",b)
	}
}

func main(){
	list(10)
}