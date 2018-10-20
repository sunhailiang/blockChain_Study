package main

import "fmt"

//父类
type Opt struct {
	num1 int
	num2 int
}

//加法子类
type AddOpt struct {
	Opt
}

//加法子类
type SubOpt struct {
	Opt
}

//加法方法实现
func (add *AddOpt) Operate() int {
	return add.num1 + add.num2
}

//减法方法实现
func (sub *SubOpt) Operate() int {
	return sub.num1 - sub.num2
}

//对象  空的
type OptFratory struct {
}

//设计模式  对于面向对象基于M(模型)V(视图)C(控制器)有26种
//工厂模式
//num1  值1  num2 值2 op 运算符
func (of *OptFratory) OptCalc(num1 int, num2 int, op string) (value int) {

	//通过运算符 进行分类计算
	//通过接口进行统一处理
	var opter Opter
	//根据不同运算符
	switch op {
	case "+":
		//创建加法对象
		var add AddOpt = AddOpt{Opt{num1, num2}}
		//value = add.Operate()
		opter=&add

	case "-":
		//创建减法对象
		var sub SubOpt = SubOpt{Opt{num1, num2}}
		opter=&sub
	}

	//value =opter.Operate()
	//通过多态来实现接口操作
	value=Fratory(opter)
	return
}

//定义接口
type Opter interface {
	Operate() int
}

//多态实现
func Fratory(o Opter) (value int){
	value=o.Operate()
	return
}
func main() {
	//1、基于继承和方法
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
	//2、基于继承方法和接口
	//	var o Opter
	//	o = &AddOpt{Opt{10, 20}}
	//	value := o.Operate()
	//	fmt.Println(value)
	//
	//	o = &SubOpt{Opt{20, 10}}
	//	value = o.Operate()
	//	fmt.Println(value)

	//3、基于继承 方法 接口 多态和设计模式

	var optf OptFratory
	value := optf.OptCalc(10, 20, "-")
	fmt.Println(value)
}
