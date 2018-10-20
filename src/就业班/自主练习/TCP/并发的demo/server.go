package main

import (
	"net"
	"fmt"
	"strings"
)

func main() {
	listenner()
}
//监听
func listenner()  {
	listenner,err:=net.Listen("tcp",":8889")
	if err!=nil{
		fmt.Println("listen err:",err)
		return
	}
	//接收多个用户
	for {
		conn,err:=listenner.Accept()
		if err!=nil {
			fmt.Println("accept err:",err)
			return
		}
		//给请求用户新建协程
		go HandleConn(conn)
	}
}

//客户端消息处理
func  HandleConn(conn net.Conn)  {
	//获取客户端网络地址信息
	//addr:=conn.RemoteAddr().String()
	buf:=make([]byte,4096)

	for {
		n,err:=conn.Read(buf)
		if err!=nil {
			fmt.Println("read err:",err)
			return
		}
		//如果用户输入exit表示退出
		fmt.Println("这是个什么玩意儿",string(buf[:n]))
		//if string(buf[:n])!=""&&string(buf[:n-2])=="exit"{
         // fmt.Println(addr,"exit")
         // return
		//}
		//回发给客户端
		conn.Write([]byte(strings.ToUpper(string(buf[:n]))))
	}
	//
	defer conn.Close()

}

