package main

import "fmt"

type Student struct {
	name   string
	age    int
	gender string
	csocre int
	mscore int
	escore int
}

func (stu *Student) sayhello() {
	fmt.Printf("大家好，我是%s,我今年%d 岁了，我是%s生", stu.name, stu.age, stu.gender)
}
func (stu *Student) printscore() {
	fmt.Printf("我的平均分是%d", (stu.csocre+stu.escore+stu.mscore)/3)
}

func main() {
	var st Student
	st.gender="男"
	st.age=18
	st.mscore=80
	st.escore=80
	st.csocre=100
	st.name="harry"

	st.sayhello()
	st.printscore()
}
