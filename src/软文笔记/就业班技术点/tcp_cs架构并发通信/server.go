package main
import (
	"fmt"
	"net"
	"strings"
)

//tcp：（transmission control protocol）网络传输控制协议，是一种面向链接可靠，基于字节的传输通信协议
//tcp/ip:模型大致分为四个层次，应用，传输，网络，链路
//链路层(遵循协议ARP:作用借助IP获取mac地址)：以太网规定联网必须使用网卡接口，网卡使不同计算机之间互相连接。因此数据传输，数据包必须是网卡之间的传递，而网卡之间的传输是因为网卡的mac地址，这里的mac地址就是数据收发的物理地址
//网络层(核心协议IP):通过ip在网络环境中找到主机
//传输层(TCP/UDP协议)：通过（port端口）锁定目标机器的某个进程
//应用层（FTP/HTTP协议）:对数据进行封装，解封，规定数据格式，没有封装的数据不能在网络环境中传递

//socket（套接字）用于描述IP地址和端口，可以实现不同程序之间的数据通信
//在tcp/ip协议中，ip+tcp/udp+端口号来确定网络通信的一个进程，IP+端口号就对应一个socket
//socket一般成对出现
//socket:一般两种格式：
//流式：面向链接的socket，针对于面向链接的TCP服务应用
//数据报式：无连接socket,应用于无连接的UDP服务应用

func main() {
	//创建监听socket
	listenner, err := net.Listen("tcp", "127.0.0.1:8001")
	serverErr("net.Listen", err)
	defer listenner.Close()
	fmt.Println("等待客户端连接...")
	//循环接收信息
	for {
		conn, err := listenner.Accept()
		serverErr("listenner.Accept", err)
		//开启处理信客户端信息的并发go协程
		go handlerConn(conn)
	}

}

//处理客户端信息
func handlerConn(conn net.Conn) {
	defer conn.Close()
	//创建缓冲区
	buf := make([]byte, 4096)
	//获取连接的IP端口
	addr := conn.RemoteAddr()
	fmt.Println(addr, ": 已经连接成功")
	//循环读取信息
	for {
		n, err := conn.Read(buf)
		//主动退出
		if "exit\r\n" == string(buf[:n]) || "exit\n" == string(buf[:n]) {
			fmt.Println(addr, "已经退出~")
			return
		}
		//关闭窗口
		if n == 0 {
			fmt.Println(addr, "已经关闭窗口，退出连接")
			return
		}
		serverErr("conn.Read", err)
		fmt.Println("服务端已接收", string(buf[:n]))
		//回复客户端
		conn.Write([]byte(strings.ToUpper(string(buf[:n]))))
	}

}

//封装错误处理
func serverErr(str string, err error) {
	if err != nil {
		fmt.Println(str+"err:", err)
		return
	}
}
