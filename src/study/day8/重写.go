package main

import "fmt"

type stu struct {
	id int
	name string
}
type stu2 struct {
	stu
	gender string
}

func (s stu) study()  {
	fmt.Println(s.name)
}

func (s stu2)study()  {
	fmt.Println(s.name)
}
func main() {
	var s stu2
	s.name="harry"
	s.study()
}
