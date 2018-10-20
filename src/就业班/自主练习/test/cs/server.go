package main

import (
	"net"
	"fmt"
	"os"
	"strings"
)

func main() {
	//监听
	listenner, err := net.Listen("tcp", "127.0.0.1:6666")
	Err("net.Listen", err)
	//阻塞，等待接收
	//起go程序处理单独客户信息
	for {
		conn, err := listenner.Accept()
		Err("listenner.Accept", err)
		go dealRequest(conn)
	}
	defer listenner.Close()
}

func dealRequest(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, 4096)
	for {
		n, err := conn.Read(buf)
		if n == 0 {
			fmt.Println("客户端退出")
			return
		}
		if "exit\r\n" == string(buf[:n]) || "exit\n" == string(buf[:n]) {
			fmt.Println("客户主动退出")
			return
		}
		Err("conn.Read", err)
		//接收信息并打印
		fmt.Println("接收到客户端信息：", string(buf[:n]))
		//回写给客户端
		conn.Write([]byte(strings.ToUpper(string(buf[:n]))))
	}
}

func Err(str string, err error) {
	if err != nil {
		fmt.Println(str+" err:", err)
		os.Exit(1)
	}
}
