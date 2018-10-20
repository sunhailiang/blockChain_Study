package main

import "fmt"

func main() {
	var factory factory
	res := factory.calcFactory(20, 20, "*")
	fmt.Println("结果:", res)
}

//定义计算接口
type caculator interface {
	calc() int
}

//params
type params struct {
	num1 int
	num2 int
}

//add
type addObj struct {
	params
}

//实现加法接口
func (add *addObj) calc() int {
	return add.num1 + add.num2
}

//sub
type subObj struct {
	params
}

//实现减法接口
func (sub *subObj) calc() int {
	return sub.num1 - sub.num2
}

//mult
type multObj struct {
	params
}

//实现乘法接口
func (mult *multObj) calc() int {
	return mult.num1 * mult.num2
}

//div
type divObj struct {
	params
}

//实现除法接口
func (div *divObj) calc() int {
	return int(div.num1 / div.num2)
}

//factory
type factory struct {
}

func (fac *factory) calcFactory(num1, num2 int, char string) int {
	var temp caculator
	switch char {
	case "+":
		doadd := addObj{params{num1, num2}}
		temp = &doadd
	case "-":
		dosub := subObj{params{num1, num2}}
		temp = &dosub
	case "*":
		domult := multObj{params{num1, num2}}
		temp = &domult
	case "/":
		dodiv := divObj{params{num1, num2}}
		temp = &dodiv
	}
	return getResObj(temp)
}

func getResObj(cal caculator) int {
	return cal.calc()
}
