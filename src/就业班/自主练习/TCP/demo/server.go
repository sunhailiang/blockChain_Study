package main

import (
	"net"
	"fmt"
)

func main() {
	listenner()
}

//监听函数
func listenner() {
	//监听端口
	listenner, err := net.Listen("tcp", ":8888")
	if err != nil {
		fmt.Println("listen err:", err)
		return
	}
	//主程序结束时关闭链接
	defer listenner.Close()
	//等待客户端链接
	fmt.Println("等待客户端连接...")
	conn, err := listenner.Accept()
	if err != nil {
		fmt.Println("accept err:", err)
		return
	}

	//接收数据
	fmt.Println("与客户端成功建立连接...")
	accept(conn)

}
//接收数据
func accept(conn net.Conn) {
	buf := make([]byte, 4096)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("read err:", err)
		return
	}
	fmt.Println("读到数据", string(buf[:n]))
	//接收完关闭连接
	defer conn.Close()
}

