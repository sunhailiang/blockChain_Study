package main

type student2 struct {
	name   string
	age    int
	sex    string
	cscore int
	mscore int
	escore int
}

//为对象赋值
func (s *student2) InitInfo(name string, age int, sex string, cscore int, mscore int, escore int) {
	s.name=name
	s.age=age
	s.sex=sex
	s.mscore=mscore
	s.cscore=cscore
	s.escore=escore
}


func main() {
	//stu := student2{"贾宝玉", 18, "男", 66, 77, 88}
	var stu student2
	//初始化对象信息
	stu.InitInfo("贾宝玉", 18, "男", 66, 77, 88)

	SayHello()
	//stu.SayHello()
	stu.PrintScore()
}
