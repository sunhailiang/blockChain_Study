package main

import (
	"fmt"
	"net"
	"os"
	"io"
)

func main() {
	//输入文件地址获取文件名
	arg := os.Args
	if len(arg) != 2 {
		fmt.Println("输入格式：xxx.go  filepath")
		return
	}
	//获取文件路径
	filePath := arg[1]
	//获取文件信息
	fileInfo, err := os.Stat(filePath)
	Err("os.Stat", err)
	//获取文件名称
	fileName := fileInfo.Name()
	//发起TCP请求
	conn, err := net.Dial("tcp", "127.0.0.1:8008")
	Err("net.Dial", err)
	//发送文件名
	sendFileName(fileName, conn, filePath)
}

//发送文件名称
func sendFileName(fileName string, conn net.Conn, filePath string) {
	buf := make([]byte, 4096)
	//发送文件名
	fmt.Println("拿得到吗？", fileName)
	n, err := conn.Write([]byte(fileName))
	Err("conn.Write\n", err)
	//读取服务器回执
	n, err = conn.Read(buf)
	Err("conn.Read", err)
	if string(buf[:n]) == "ok" {
		sendFile(conn, filePath)
	}
}

//发送文件
func sendFile(conn net.Conn, filePath string) {
	//打开并读取本地文件
	file, err := os.Open(filePath)
	defer file.Close()
	Err("os.Open", err)
	//创建切片缓冲区
	buf := make([]byte, 4096)
	for {
		n, err := file.Read(buf)
		if err != nil {
			if err == io.EOF {
				fmt.Println("本地文件读取完毕")
			} else {
				fmt.Println("file.Read", err)
			}
			return
		}
		//往网络中写数据
		conn.Write(buf[:n])
	}
	defer conn.Close()
	defer file.Close()
}

func Err(str string, err error) {
	if err != nil {
		fmt.Println(str+" err:", err)
		return
	}
}
