package main

import (
	"net"
	"fmt"
)

func main() {
  listenner,err:=net.Listen("tcp",":8888")

  if err!=nil{
  	fmt.Println("listen err",err)
  	return
  }

  //主协程刚结束结束listenner
  defer  listenner.Close()

  fmt.Println("服务器等待客户端链接...")

  //接收客户端请求
  conn,err:=listenner.Accept()
  if err!=nil{
  	fmt.Println("accept err,",err)
  	return
  }
  defer  conn.Close()//使用结束断块连接
  fmt.Println("连接建立成功...")
  //接收数据
  buf:=make([]byte,1024)//创建缓冲区
  n,err:=conn.Read(buf)
	if err!=nil {
		fmt.Println("read err,",err)
		return
	}
	fmt.Println("服务器读到:",string(buf[:n]))

}
