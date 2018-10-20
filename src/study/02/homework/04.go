package main

import (
	"fmt"
)

func main()  {
	var str string
	fmt.Println("请输入字符串")
	fmt.Scanf("%q",&str)
	fmt.Println(str)
	getStrLen(&str)
}

func getStrLen(str *string)  {
	byts:=[]byte(*str)
	var sum int
	for i:=0;i< len(byts);i++{
		sum++
	}
	fmt.Println("Result is:",sum)
}
