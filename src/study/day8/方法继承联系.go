package main

import "fmt"

/*
  记者：我是记者  我的爱好是偷拍 我的年龄是34 我是一个男狗仔
  程序员：我叫孙全 我的年龄是23 我是男生 我的工作年限是 3年

 */

 type  Person struct {
 	name string
 	age int
 	gender string
 }

 type  person1 struct {
 	Person
	 workTime int
 }
 type person2 struct {
 	Person
 	hobby string
 }

 //父类方法

func (p Person)sayhello()  {
	fmt.Printf("我是%s,我是%s生，我今年%d岁",p.name,p.gender,p.age)
}

//子类

func (p person2)sayHi()  {
	p.sayhello()
	fmt.Printf("我的爱好是%s\n",p.hobby)
}
//子类

func (p person1)sayHi()  {
	p.sayhello()
	fmt.Printf("我已经工作%d年",p.workTime)
}
func main() {
	var p1 person1=person1{Person{"小马",28,"男"},5}
	var p2 person2=person2{Person{"小卓",45,"男"},"偷拍"}
	p1.sayHi()
	p2.sayHi()
}
