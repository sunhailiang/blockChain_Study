package main

import (
	"net"
	"fmt"
	"strings"
	"time"
)

//用户结构体类型makeMsg
type Client struct {
	C    chan string
	Name string
	Addr string
}

//定义全局Map储存在线用户key:IP+PORT   value Client
var onlineMap map[string]Client

//定义全局管道处理消息
var massage = make(chan string)

//有信息就写给客户端
func WriteMsgToClient(clnt Client, conn net.Conn) {
	//循环跟踪
	for msg := range clnt.C {
		conn.Write([]byte(msg + "\n"))
	}
}

//定制信息格式
func makeMsg(clnt Client, msg string) (buf string) {
	buf = "[" + clnt.Addr + "]" + clnt.Name + ":" + msg
	return
}

func handleConnect(conn net.Conn) {
	defer conn.Close()
	//获取新连接用户的IP/PORT
	userAddr := conn.RemoteAddr().String()
	//给新用户创建单独结构体
	clnt := Client{make(chan string), userAddr, userAddr}
	//每个IP对应一个用户
	onlineMap[userAddr] = clnt
	//创建一个协程转给当前的用户发送信息
	go WriteMsgToClient(clnt, conn)
	//广播用户上线
	massage <- makeMsg(clnt, "login")
	//客户端主动退出，服务器剔除
	isCloseChan := make(chan bool)
	//客户端操作活跃，保持链接
	isLiveChan := make(chan bool)
	//广播给所有用户(聊天室)
	go func() {
		buf := make([]byte, 4096)
		for {
			n, err := conn.Read(buf)
			if n == 0 {
				isCloseChan <- true
				fmt.Printf("用户%s退出登录\n", clnt.Name)
				return
			}
			Err("conn.Read", err)
			msg := string(buf[:n-1])
			//获取在线用户列表
			if len(msg) == 3 && msg == "who" {
				conn.Write([]byte("user list \n"))
				for _, v := range onlineMap {
					userInfo := v.Name + ":" + v.Addr + "\n"
					conn.Write([]byte(userInfo))
				}
			} else if len(msg) >= 8 && msg[:6] == "rename" {
				newName := strings.Split(msg, "|")[1]
				clnt.Name = newName
				onlineMap[clnt.Addr] = clnt
				conn.Write([]byte("rename success"))
			} else {
				massage <- makeMsg(clnt, msg)
			}
			isLiveChan <- true
		}
	}()
	//超时退出
	for {
		select {
		case <-isCloseChan:
			delete(onlineMap, userAddr) //删除用户
			massage <- makeMsg(clnt, "logout")
			return
		case <-isLiveChan:
			//条件满足重置After，保持链接
		case <-time.After(time.Second * 10):
			delete(onlineMap, userAddr) //删除用户
			massage <- makeMsg(clnt, "time out leave")
			return
		}
	}
}

//做在线用户信息转发
func manager() {
	//开辟空间
	onlineMap = make(map[string]Client)
	//循环读取通道信息
	for {
		msg := <-massage
		for _, clnt := range onlineMap {
			clnt.C <- msg
		}
	}
}
func main() {
	listenner, err := net.Listen("tcp", "127.0.0.1:8001")
	Err("net.Listen", err)
	defer listenner.Close()
	//创建协程处理信息
	go manager()
	//循环接收客户端请求链接
	for {
		conn, err := listenner.Accept()
		//此处出错，要继续接收其他客户端链接
		if err != nil {
			fmt.Println("listenner.Accept", err)
			continue
		}
		defer conn.Close()

		//给请求的过来的用户处理信息
		go handleConnect(conn)

	}
}
//通用错误处理封装
func Err(str string, err error) {
	if err != nil {
		fmt.Println(str+" err:", err)
		return
	}
}
