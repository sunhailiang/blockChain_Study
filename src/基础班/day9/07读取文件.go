package main

import (
	"os"
	"fmt"
	"bufio"
)

func main0701() {
	//只读方式打开文件
	fp,err:=os.Open("D:/a.txt")
	/*
	1、文件不存在
	2、文件权限
	3、文件打开上限
	 */
	if err !=nil{
		fmt.Println("文件打开失败")
		return
	}
	b:=make([]byte,100)
//换行也会作为字符一部分进行读取
	fp.Read(b)

	//fmt.Println(b)
	//for i:=0;i<len(b);i++{
	//	if b[i]!=0{
	//		fmt.Printf("%c",b[i])
	//
	//	}
	//}
	fmt.Println(string(b))
}


func main(){
	fp,err:=os.Open("D:/a.txt")
	if err !=nil{
		fmt.Println("文件打开失败")
		return
	}

	//创建切片缓冲区
	r:=bufio.NewReader(fp)
	//读取一行内容
	b,_:=r.ReadBytes('\n')
	//打印切片中字符的ASCII值
	//fmt.Println(b)
	//将切片转成string打印  汉字
	fmt.Print(string(b))
	//如果在文件截取中  没有标志位（分隔符）到文件末尾自动停止  EOF -1 文件结束标志
	b,_=r.ReadBytes('\n')

	fmt.Println(string(b))
}