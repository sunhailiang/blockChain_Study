package main

import (
	"net"
	"fmt"
)

func main() {
	connServer()
}

//发起连接
func connServer()  {
	conn,err:=net.Dial("tcp","127.0.0.1:8888")
	if err!=nil {
		fmt.Println("client conn server err:",err)
		return
	}
	sendMsg(conn)
}
//发送请求
func sendMsg(conn net.Conn)  {
	_,err:=conn.Write([]byte("are you ok??"))
	if err!=nil {
		fmt.Println("write err:",err)
		return
	}
	//发送完，关闭链接
	conn.Close()
}