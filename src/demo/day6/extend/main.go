package main

import "fmt"

type person struct {
	 	id int
	 	name string
	 	age int
}
 type student struct {
 	*person
 	score float64
 }

func main()  {
    var stu student
    stu.person=new(person)
    stu.id=1
    stu.name="harry"
    stu.age=27
    stu.score=99.5
    var p person
    p=*stu.person
    fmt.Printf("编号：%d  姓名：%s, 年龄：%d",p.id,p.name,p.age)
}