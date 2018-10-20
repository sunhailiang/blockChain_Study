package main

import (
	"net"
	"os"
	"fmt"
	"io"
)

func main() {
	//输入文件名
	fmt.Println("请输入要传输的文件")
	var path string
	fmt.Scanf("%s", &path)
	fmt.Printf("文件地址:%s",path)
	//确认访问协议和端口
	conn, err := net.Dial("tcp", "127.0.0.1:8001")
	//连接服务器
	connServer(conn, err)
	////获取文件名
	fileInfo := getFileName(path)
	//发送文件名
	fmt.Println("文件名", fileInfo.Name())
	sendFileName(conn, fileInfo)
	//接收服务器回发
	reciveFormServer(conn, path)
	//发送完成关掉连接
	defer conn.Close()
}

//连接服务器
func connServer(conn net.Conn, err error) {

	if err != nil {
		fmt.Println("client conn err:", err)
		return
	}
}

//获取文件名
func getFileName(path string) os.FileInfo {
	//获取文件名
	fileInfo, err := os.Stat(path)
	if err != nil {
		fmt.Println("os.stat err:", err)
		return nil
	}
	return fileInfo
}

//发送文件名
func sendFileName(conn net.Conn, fileInfo os.FileInfo) {
	//发送文件名
	_, err := conn.Write([]byte(fileInfo.Name()))
	if err != nil {
		fmt.Println("client write title err:", err)
		return
	}
}

//接收服务器回发信息
func reciveFormServer(conn net.Conn, path string) {
	buf := make([]byte, 4096)
	//接收服务端回发的确认信息
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("read back title response err:", err)
		return
	}
	//文件名发送成功后发送
	if "ok" == string(buf[:n]) {
		//发送文件
		sendFile(path, conn)
	}
}

//发送文件
func sendFile(path string, conn net.Conn) {
	f, err := os.Open(path)
	if err != nil {
		fmt.Println("open file err:", err)
		return
	}
	//操作完成关闭连接
	defer f.Close()

	buf := make([]byte, 4096)
	for {
		n, err := f.Read(buf)
		if err != nil {
			if err == io.EOF {
				fmt.Println("file read complete")
			} else {
				fmt.Println("read err:", err)
			}
			return
		}
		conn.Write(buf[:n])
	}
}
