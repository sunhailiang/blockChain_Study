package main

import "fmt"

type demoA struct {
	id int
}
type demoB struct {
	name string
}

func (d demoB)fuck()  {

}
type demoC struct {
	demoA
	demoB
	gender string
}

func (c demoC)fuck()  {
	fmt.Println(c.name+"fuckyou")
}
func main() {
	var demo demoC
	demo.id=11
	demo.name="harry"
	demo.gender="男"
	demo.fuck()


	fmt.Printf("id：%d,姓名：%s,性别：%s",demo.id,demo.name,demo.gender)
}
