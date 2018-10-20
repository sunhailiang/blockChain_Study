package main
//先创建客户端结构体
type Client struct {
	Name string
	Addr string
	C chan  string
}

//储存在线用户
var onlineMap map[string]Client
//全局channel
var massage=make(chan string)
//



func main() {
	
}
