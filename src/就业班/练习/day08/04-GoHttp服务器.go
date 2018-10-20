package main

import (
	"net/http"
	"fmt"
)

func myHandle(w http.ResponseWriter, r *http.Request) {
	// w : 写给客户端的数据内容
	w.Write([]byte("this is a Web server"))

	// r: 从客户端读到的内容
	fmt.Println("Header:", r.Header)
	fmt.Println("URL:", r.URL)
	fmt.Println("Method:", r.Method)
	fmt.Println("Host:", r.Host)
	fmt.Println("RemoteAddr:", r.RemoteAddr)
	fmt.Println("Body:", r.Body)
}

func main()  {
	// 注册回调函数， 该函数在客户端访问服务器时，会自动被调用
	//http.HandleFunc("/itcast", myHandle)
	http.HandleFunc("/", myHandle)

	// 绑定服务器监听地址
	http.ListenAndServe("127.0.0.1:8000", nil)
}
