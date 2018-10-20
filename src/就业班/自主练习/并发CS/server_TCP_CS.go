package main

import (
	"net"
	"fmt"
	"strings"
)

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:6666")
	showErr("net.Listen", err)

	for {
		conn, err := listener.Accept() //等待接收
		showErr("listener.Accept", err)
			go handlerConn(conn)
	}
	defer listener.Close()
}
func showErr(info string, err error) {
	if err != nil {
		fmt.Println(info+":", err)
		return
	}
}

func handlerConn(conn net.Conn) {
	defer conn.Close()
	addr := conn.RemoteAddr()
	fmt.Println("客户端已经连接:", addr)
	buf := make([]byte, 4096)

	for {
		n, err := conn.Read(buf)
		if "exit\n" == string(buf[:n]) {
			fmt.Println(addr, "客户退出")
			return
		}
		if n == 0 {
			fmt.Printf("检测到客户端%s关闭窗口，断开连接",addr)
			return
		}
		showErr("conn.Read", err)
		fmt.Println("服务端接收数据:", string(buf[:n]))
		conn.Write([]byte(strings.ToUpper(string(buf[:n]))))
	}

}
