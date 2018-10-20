package main

import (
	"net"
	"fmt"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:6666")
	errShow("net.Dial", err)
	defer conn.Close()
	go func() {
		buf := make([]byte, 4096)
		for {
			n, err := os.Stdin.Read(buf)
			if err != nil {
				fmt.Println("os.Stdin.Read err", err)
				continue
			}
			conn.Write(buf[:n])
		}
	}()

	buf := make([]byte, 4096)
	for {
		n, err := conn.Read(buf)
		//如果服务器关闭了，自动返回
		if n == 0 {
			fmt.Println("检测到服务器关闭，客户端也关闭")
			return
		}
		errShow("conn.Read", err)
		fmt.Println("服务器回发信息：", string(buf[:n]))
	}

}
func errShow(str string, err error) {
	if err != nil {
		fmt.Println(str+"err:", err)
		return
	}
}
