package main

import "fmt"

//父结构体
type person struct {
	name string
	gender string
	age int
}
//子类结构体
type student struct {
	//继承
	*person  //此处是个空指针
	id  int
	score int
}
func (s *student)eat()  {
    //初始化方式1
    s.person=new(person)//为空指针开辟堆空间，用于存放
	//s.age=18
	//s.score=99
	//s.gender="女"
	//s.name="harry"
	//s.id=101
	//fmt.Printf("ID:%d,姓名：%s,年龄：%d,性别：%s,分数：%d",s.id,s.name,s.age,s.gender,s.score)
     fmt.Printf("%v",&s.person)
    }
func main() {
    var  stu student
    //初始化方式二
    //stu=student{&person{"fuck","女",99},101,99}
    stu.eat()
}
