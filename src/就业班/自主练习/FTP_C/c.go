package main

import (
	"net"
	"fmt"
)

func main() {
	conn,err:=net.Dial("tcp","127.0.0.1:6666")

	if err!=nil {
		fmt.Println("dial err:",err)
		return
	}
	defer conn.Close()

	conn.Write([]byte("fuck you"))

	buf:=make([]byte,4096)

	n,err:=conn.Read(buf)

	if err!=nil {
		fmt.Println("client read err:",err)
		return
	}
	fmt.Println("服务器回发：",buf[:n])

}
