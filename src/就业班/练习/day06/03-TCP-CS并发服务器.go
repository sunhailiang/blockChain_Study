package main

import (
	"net"
	"fmt"
	"strings"
)

func HandlerConnect(conn net.Conn) {

	defer conn.Close()

	// 获取连接的客户端 Addr
	addr := conn.RemoteAddr()
	fmt.Println(addr, "客户端成功连接！")

	// 循环读取客户端发送数据
	buf := make([]byte, 4096)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("conn.Read err:", err)
			return
		}
		fmt.Println("服务器读到数据：", string(buf[:n]))	// 使用数据

		// 小写转大写，回发给客户端
		conn.Write([]byte(strings.ToUpper(string(buf[:n]))))
	}
}

func main()  {
	// 创建监听套接字
	listener, err := net.Listen("tcp", "127.0.0.1:8001")
	if err != nil {
		fmt.Println("net.Listen err:", err)
		return
	}
	defer listener.Close()

	// 监听客户端连接请求
	for {
		fmt.Println("服务器等待客户端连接...")
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener.Accept err:", err)
			return
		}
		// 具体完成服务器和客户端的数据通信
		go HandlerConnect(conn)
	}
}
