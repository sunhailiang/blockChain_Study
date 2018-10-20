package main

import (
	"net"
	"fmt"
	"os"
)

func main() {
	//创建监听socket
	listenner, err := net.Listen("tcp", "127.0.0.1:8008")
	defer listenner.Close()
	sevErr("net.Listen", err)
	fmt.Println("等待客户端连接~")
	//等待接收数据-此处阻塞
	conn, err := listenner.Accept()
	sevErr("listenner.Accept", err)
	//读取客户端数据
	reciveFileName(conn)
	//回发数据

}

func reciveFileName(conn net.Conn) {
	buf := make([]byte, 4096)
	n, err := conn.Read(buf)
	remAddr := conn.RemoteAddr()
	fmt.Println("n数据吗？", string(buf[:n]))
	sevErr("conn.Read", err)
	fmt.Printf("已收到客户端:%v的数据%v\n", remAddr, string(buf[:n]))
	//回发数据
	_, err = conn.Write([]byte("ok"))
	sevErr("conn.Write", err)
	defer conn.Close()
	reciveFile(conn, string(buf[:n]))
}
func reciveFile(conn net.Conn, fileName string) {
	//创建存储文件
	fmt.Println("这是个啥", fileName)
	file, err := os.Create(fileName)
	sevErr("os.Create", err)
	//创建切片缓冲区
	buf := make([]byte, 4096)
	for {
		n, _ := conn.Read(buf)
		if n == 0 {
			fmt.Println("文件接收完成")
			return
		}
		//写入本地文件
		file.Write(buf[:n])
	}
	defer file.Close()
	defer conn.Close()
}

//错误处理
func sevErr(str string, err error) {
	if err != nil {
		fmt.Println(str+" err:\n", err)
		return
	}
}
