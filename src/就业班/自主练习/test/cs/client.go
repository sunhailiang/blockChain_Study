package main

import (
	"net"
	"fmt"
	"os"
)

func main() {
	//主动发起请求
	conn, err := net.Dial("tcp", "127.0.0.1:6666")
	clientErr("net.Dial", err)
	buf := make([]byte, 4096)
	defer conn.Close()
	go func() {
		for {
			n, err := os.Stdin.Read(buf)
			if err != nil {
				fmt.Println("os.Stdin.Read err:", err)
				continue
			}
			n, err = conn.Write(buf[:n])
			if err != nil {
				fmt.Println("写错了？", err)
			} else {
				fmt.Println("写了多少？", n)
			}
		}
	}()
	for {
		n, err := conn.Read(buf)
		if n==0 {
			fmt.Println("服务器关闭,客户端被动断开")
			return
		}
		clientErr("conn.Read", err)
		fmt.Println("接收到服务器信息:", string(buf[:n]))
	}
	clientErr("conn.Write", err)
}
func clientErr(str string, err error) {
	if err != nil {
		fmt.Println(str+" err:", err)
		os.Exit(1)
	}
}
