package main

import (
	"os"
	"fmt"
	"io"
)

func main() {

	//文件拷贝  从一个文件中拷贝数据到新的文件中
	fp1, err1 := os.Open("D:\\0705基础班\\01_Go语言基础第09天（面向对象下）\\03视频\\02接口定义和使用.avi")
	fp2, err2 := os.Create("D:/test.avi")

	if err1 != nil || err2 != nil {
		if err1!=nil{
			fmt.Println("err1",err1)
		}
		if err2!=nil{
			fmt.Println("err2",err2)
		}
		fmt.Println("拷贝失败")
		return
	}

	defer fp1.Close()
	defer fp2.Close()

	b:=make([]byte,1024*1024)

	for{
		//块读取
		n,err:=fp1.Read(b)
		if err==io.EOF{
			break
		}
		fp2.Write(b[:n])
	}
	fmt.Println("拷贝完成")
}
