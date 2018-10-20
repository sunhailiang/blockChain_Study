package main

import "fmt"

type Opt struct {
	num1 int
	num2 int
}

type AddOpt struct {
	Opt
}
type SubOpt struct {
	Opt
}

func (add *AddOpt) Operate() int {
	return add.num1 + add.num2
}
func (sub *SubOpt) Operate() int {
	return sub.num1 - sub.num2
}

//定义接口
type Opter interface {
	Operate()int
}
func main() {

	//var add AddOpt
	//add.num1 = 10
	//add.num2 = 20
	//value := add.Operate()
	//fmt.Println(value)
	//
	//var sub SubOpt
	//sub.num1=20
	//sub.num2=10
	//value=sub.Operate()
	//fmt.Println(value)

	var o Opter
	o=&AddOpt{Opt{10,20}}
	value:=o.Operate()
	fmt.Println(value)

	o=&SubOpt{Opt{20,10}}
	value=o.Operate()
	fmt.Println(value)

}
