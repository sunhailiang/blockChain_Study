package main

import (
	"fmt"
	"strings"
)

func main() {
	//从Email中提取出用户名和域名：abc@163.com

	fmt.Println("请输入邮箱")
	var email=""
	fmt.Scanf("%s",&email)

	 slice:=strings.Split(email,"@")
	 fmt.Printf("账号：%s，域名:%s",slice[0],slice[1])




}
