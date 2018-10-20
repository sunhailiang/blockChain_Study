package main

import "fmt"

/*
  记者：我是记者  我的爱好是偷拍 我的年龄是34 我是一个男狗仔
  程序员：我叫孙全 我的年龄是23 我是男生 我的工作年限是 3年

 */

 type Person1 struct {
	name string
	age int
	sex string
 }

type Rep struct {
	Person1
	hobby string
}

type Dep struct {
	Person1
	worktime int
}
func (p Person1)SayHello(){
	fmt.Printf("大家好，我是%s,我是%s生，我今年%d岁了",p.name,p.sex,p.age)
}

func (r Rep)SayHi(){
	r.SayHello()
	fmt.Println("我的爱好是",r.hobby)
}
func (d Dep)SayHi(){
	d.SayHello()
	fmt.Println("我的工作年限是",d.worktime)
}
func main() {

	var r Rep=Rep{Person1{"卓伟",40,"男"},"偷拍"}
	var d Dep =Dep{Person1{"汤姆逊",68,"男"},40}
	r.SayHi()
	d.SayHi()

}
