package main

import "fmt"

type person struct {
	Name string
	age int
	gender string
}
func main()  {
	var personChan chan interface{}
	personChan=make(chan interface{},10)

	p:=person{Name:"harry",age:16,gender:"男"}
	//将内容写入管道
	personChan<-p
	//输出管道内容
	var p2 interface{}
	p2=<-personChan

	var p3 person
	//将管道输出的interfase类型转成结构体
	p3,ok:=p2.(person)
	//判断是否转换成功
	if !ok{
      fmt.Println("can not reverse")
	}
	fmt.Println("sss",p3)
}