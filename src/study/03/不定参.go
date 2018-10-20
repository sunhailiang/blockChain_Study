package main

import "fmt"

func main()  {

	res:=Concat("你","瞅啥？","我就是","你","爸爸")
	fmt.Println(res)
}
func Concat(args ...string)string{
	var str string
	for _,v:=range args{

		 str+=v
	}
	return  str
}
func concat(args ...string) string {
	var str string
	for _,v:=range args{

		str+=v
	}
	return  str
}