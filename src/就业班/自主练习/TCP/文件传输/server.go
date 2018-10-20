package main

import (
	"net"
	"fmt"
	"os"
	"io"
)

func main() {
	//创建监听
	listenner, err := net.Listen("tcp", ":8001")
	if err != nil {
		fmt.Println("conn err:", err)
		return
	}

	waitForClient(listenner)

	defer listenner.Close()

}

//等待连接客户端
func waitForClient(listenner net.Listener) {
	conn, err := listenner.Accept()
	if err != nil {
		fmt.Println("accept err:", err)
		return
	}
	reciveFileName(conn)
	defer conn.Close()
}

//读取文件名
func reciveFileName(conn net.Conn) {

	buf := make([]byte, 4096)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("fileNameread err:", err)
		return
	}
	fileName := string(buf[:n])
	fmt.Println("有反应么？",fileName)
	conn.Write([]byte("ok"))
	//接收文件
	recivFile(fileName, conn)
}

//读取文件

func recivFile(fileName string, conn net.Conn) {
	f, err := os.Create(fileName)
	if err != nil {
		fmt.Println("create file err:", err)
		return
	}
	defer f.Close()

	//接收并写入创建的文件中
	buf := make([]byte, 4096)

	for{
		n, err := conn.Read(buf)
		if err != nil {
			if err == io.EOF {
				fmt.Println("file read complete")
			} else {
				fmt.Println("file read err:", err)
			}
			return
		}
		fmt.Println("有东西吗？",buf[:n])
		f.Write(buf[:n])
	}


}
