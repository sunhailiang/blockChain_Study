package main

import (
	"net"
	"fmt"
	"os"
)
//交互了两次？？？？？？？？？？
func main() {
	//主动发起连接
	conn, err := net.Dial("tcp", "127.0.0.1:8889")
	if err != nil {
		fmt.Println("client conn err:", err)
		return
	}
	go sendMsg(conn)
	reciveMsg(conn)

}
func sendMsg(conn net.Conn) {
	str := make([]byte, 4096) //创建切片缓冲区存储数据
	for {
		n, err :=os.Stdin.Read(str)//获取键盘输入内容
		if err != nil {
			fmt.Println("client read err:", err)
			return
		}
		//发送给服务器
		_, err = conn.Write(str[:n])
		if err != nil {
			fmt.Println("client write err", err)
			return
		}
	}
}

func reciveMsg(conn net.Conn)  {
	 //创建接收数据的切片缓冲区
	 buf:=make([]byte,4096)
	 for{
	 	n,err:=conn.Read(buf)
	 	if err!=nil{
	 		fmt.Println("recive err:",err)
	 		return
		}
		fmt.Println("服务端回发数据",string(buf[:n]))
	 }
	//发送完成后关闭
	defer conn.Close()
}
