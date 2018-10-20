package main

import (
	"os"
	"fmt"
)

func main() {
	args := os.Args

	if len(args) < 2 {
		fmt.Println("正确输入格式：xxxxx.go  文件名")
		return
	}
	//提取文件名
	fileInfo, err := os.Stat(args[1])
	Err("os.Stat", err)
	//获取文件名
	fmt.Println("文件名:", fileInfo.Name())
	//获取文件大小
	fmt.Println("文件名:", fileInfo.Size())

}
func Err(str string, err error) {
	if err != nil {
		fmt.Println(str+" err:", err)
		return
	}
}
