package main

import (
	"net"
	"fmt"
)

func main()  {
	// 指定服务器 通信协议、IP地址、port。 创建一个用于监听的 socket
	listener, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("net.Listen err:", err)
		return
	}
	defer listener.Close()

	fmt.Println("服务器等待客户端建立连接...")
	// 阻塞监听客户端连接请求, 成功建立连接，返回用于通信的socket
	conn, err := listener.Accept()
	if err != nil {
		fmt.Println("listener.Accept() err:", err)
		return
	}
	defer conn.Close()
	fmt.Println("服务器与客户端成功建立连接！！！")
	// 读取客户端发送的数据
	buf := make([]byte, 4096)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("conn.Read err:", err)
		return
	}
	// 处理数据—— 打印
	fmt.Println("服务器读到数据：", string(buf[:n]))
}
