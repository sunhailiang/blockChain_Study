package main
import (
	"net"
	"fmt"
	"os"
)

func main() {
	//发起访问连接
	conn, err := net.Dial("tcp", "127.0.0.1:8001")
	cliErr("net.Dial", err)
	//循环发送信息
	go sendMsg(conn)
	//循环接收服务器信息
	recive(conn)
}

//循环发送信息
func sendMsg(conn net.Conn) {
	//创建缓冲区
	buf := make([]byte, 4096)
	for {
		//输入信息
		n, err := os.Stdin.Read(buf)
		if err != nil {
			fmt.Println("os.Stdin.Read err:", err)
			continue
		}
		//写入的信息发送给服务器
		conn.Write(buf[:n])
	}
}

//循环接收服务器返回信息
func recive(conn net.Conn) {
	//创建切片缓冲区
	buf := make([]byte, 4096)
	//循环获取服务器数据
	for {
		n, err := conn.Read(buf)
		if n == 0 {
			fmt.Println("服务器关闭，客户端退出")
			return
		}
		cliErr("conn.Read", err)
		fmt.Println("客户端回发信息：", string(buf[:n]))
	}
}
//封装错误处理
func cliErr(str string, err error) {
	if err != nil {
		fmt.Println(str+"err:", err)
		return
	}
}
