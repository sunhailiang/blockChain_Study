package main

import (
	"net"
	"fmt"
	"time"
)

func main() {
	udpAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:6666")
	Err("ResolveUDPAddr", err)
	conn, err := net.ListenUDP("udp", udpAddr)
	Err("net.ListenUDP", err)
	for {
		buf := make([]byte, 4096)
		n, clientAddr, err := conn.ReadFromUDP(buf)
		Err("ReadFromUDP", err)
		fmt.Printf("接收到客户端%v的信息:%s", clientAddr, string(buf[:n]))
		go func() {
			t := time.Now().String()
			_, err := conn.WriteToUDP([]byte(t), clientAddr)
			Err("WriteToUDP", err)
		}()
	}
}

func Err(str string, err error) {
	if err != nil {
		fmt.Println(str+" err:", err)
		return
	}
}
