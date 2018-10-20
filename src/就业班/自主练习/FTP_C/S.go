package main

import (
	"net"
	"fmt"
)

func main() {

	//listen
	lisenner,err:=net.Listen("tcp","127.0.0.1:6666")
	if err!=nil{
		fmt.Println("listen err:",err)
		return
	}
	defer lisenner.Close()
	fmt.Println("正在等待客户端连接...")
	//等待接收

	conn,err:=lisenner.Accept()
	if err!=nil {
		fmt.Println("accept err:",err)
		return
	}
	buf:=make([]byte,4096)
	n,err:=conn.Read(buf)
	if err!=nil {
		fmt.Println("read err:",err)
		return
	}

	fmt.Println("接收信息:",string(buf[:n]))

	conn.Write([]byte("fuck u"))

	defer  conn.Close()

	
}
