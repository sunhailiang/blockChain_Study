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


func main0105(){

	//str:="helloworlld"
	////字符串替换元素  参数 字符串  被替换内容 替换内容 替换次数
	//s:=strings.Replace(str,"ll","ww",-1)

	str:="我星你大爷"
	//屏蔽关键字
	s:=strings.Replace(str,"你大爷","***",-1)
	fmt.Println(s)

}

func main0106(){
	//str:="www.itcast.cn"
	str:="1234-3867-9017-1859"
	//s:=strings.Split(str,".")
	//字符串切割  返回值是[]string
	s:=strings.Split(str,"-")

	fmt.Println(s)
}

func main0107(){
	str:="====hello======world===="
	//s:=strings.Trim(str," ")
	//去掉字符串前后的内容
	s:=strings.Trim(str,"=")
	fmt.Println(s)
}

func main(){
	str:="    are  u   ok  ?     "
	//去掉字符串中的空格  并返回有效数据切片
	s:=strings.Fields(str)

	fmt.Println(s)
}