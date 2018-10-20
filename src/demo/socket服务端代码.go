package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("start server...")
	//监听端口
	listen,err:=net.Listen("tcp","0.0.0.0:50000")
	if err!=nil {
		fmt.Println("listen failed err:",err)
		return
	}
	for {
		//不断接收请求，不阻塞
		conn,err:=listen.Accept()
		if err!=nil{
			fmt.Println("Accept failed,err:",err)
			continue
		}
		//如果没有错误就执行操作
		go process(conn)
	}
}
//执行具体业务
func process(conn net.Conn)  {
	//执行完成最终关闭连接
	defer conn.Close()
	for {
       buf:=make([]byte,512)
       _,err:=conn.Read(buf)
       if err!=nil{
       	   fmt.Println("read failed err:",err)
       	   return
	   }
	   fmt.Println("read:",string(buf))
	}
}
