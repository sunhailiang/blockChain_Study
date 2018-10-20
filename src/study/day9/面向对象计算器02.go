package main

import "fmt"

//定义父类输入数字
type inputNum struct {
	num1 int
	num2 int
}

//定义统一计算接口
type calcer interface {
	Operate() int
}

//定义加减乘除结构体
type add struct {
	inputNum
}
type sub struct {
	inputNum
}
type mutil struct {
	inputNum
}
type chu struct {
	inputNum
}

//统一实现接口
//加
func (add *add) Operate() int {
	return add.num1 + add.num2
}

//减
func (sub *sub) Operate() int {
	return sub.num1 - sub.num2
}

//乘
func (mutil *mutil) Operate() int {
	return mutil.num1 * mutil.num2
}

//除
func (chu *chu) Operate() int {
	return chu.num1 / chu.num2
}

//创建工厂结构体
type Factory struct {
}

//工厂函数
func (fac *Factory) optFactory(num1, num2 int, opt string) (value int) {
	var cal calcer
	switch opt {
	case "+":
		var optAdd = add{inputNum{num1, num2}}
		cal = &optAdd
	case "-":
		var optSub = sub{inputNum{num1, num2}}
	       cal=&optSub
	case "*":
		var optMult=mutil{inputNum{num1,num2}}
		cal=&optMult
	case "/":
		var optDiv=mutil{inputNum{num1,num2}}
		cal=&optDiv
	}

	return moreStatus(cal)
}

//利用接口实现多态
func moreStatus(cal calcer) int {
	return cal.Operate()
}

func main() {
    var factory Factory
    result:= factory.optFactory(10,30,"*")
    fmt.Println(result)
}
