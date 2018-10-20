package main

import (
	"strings"
	"fmt"
)

func main0101(){
	str:="helloworld"

	//在一个字符串中查找另外一个字符串是否出现 模糊查找  返回值为bool类型
	value:=strings.Contains(str,"hellworld")
	if value{
		fmt.Println("找到")
	}else{
		fmt.Println("未找到")
	}
	//fmt.Println(value)
}


func main0102(){

	//字符串拼接
	s:=[]string{"1234","3867","9017","1859"}
	//将一个字符串切片 拼接成一个字符串
	str:=strings.Join(s,"-")

	fmt.Println(str)
}

func main0103(){
	str:="日照香炉生紫烟"

	//i:=strings.Index(str,"照香炉")
	//在一个字符串中查找另外一个字符串是否出现  返回值为整型  如果找到具体下标  找不到是-1
	i:=strings.Index(str,"赵香炉")

	fmt.Println(i)

}

func main0104(){
	str:="你瞅啥"

	//将一个字符串重复n次  n取值范围大于等于0
	s:=strings.Repeat(str,0)

	fmt.Println(s)
}