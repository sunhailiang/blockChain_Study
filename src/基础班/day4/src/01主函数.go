package main

import "userinfo"

//全局变量
var num int = 123
//在同级别目录报名要相同
func main() {
	add(10,20)

	userinfo.Login()
	userinfo.DeleteUser()
	userinfo.SelectUser()
}
