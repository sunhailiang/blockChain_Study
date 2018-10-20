package main

import (
	"net/http"
	"fmt"
)

func main() {
	//路由匹配
	http.HandleFunc("/",Hello)
	http.HandleFunc("/user/login",Login)
	err:=http.ListenAndServe("0.0.0.0:8001",nil)
	if err!=nil{
		fmt.Println("listen server failed:err",err)
	}
}
func Hello(w http.ResponseWriter,r *http.Request)  {
	fmt.Println("handle hello")
	fmt.Fprintf(w,"返回成功")
}
func Login(w http.ResponseWriter,r *http.Request)  {
	fmt.Println("userXXX logined")
	fmt.Fprintf(w,"欢迎登陆")
}
