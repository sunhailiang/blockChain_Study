package main
//
//import "fmt"
//
////父类结构体
//type inputNums struct {
//	num1 int
//	num2 int
//}
//
////定义通用计算协议
//type calc interface {
//	Operate() int
//}
//
////加法
//type add struct {
//	inputNums
//}
//
////实现统一计算接口
//func (ad *add) Operate() int {
//	return ad.num1 + ad.num2
//}
//
////减法
//type sub struct {
//	inputNums
//}
//
////实现统一计算接口
//func (sub *sub) Operate() int {
//	return sub.num1 - sub.num2
//}
//
////乘法
//type mutl struct {
//	inputNums
//}
//
////实现乘法
//
//func (mutl *mutl) Operate() int {
//	return mutl.num1 * mutl.num2
//}
//
////除法
//type chu struct {
//	inputNums
//}
//
////实现除法
//func (chu *chu) Operate() int {
//	return chu.num1 / chu.num2
//}
//
////声明工厂结构体
//type optFactory struct {
//}
//
////实现工厂方法
//var ca calc
//
//func (optFac *optFactory) optFactory(num1, num2 int, opt string) (resValue int) {
//	switch opt {
//	case "+":
//		var addres = add{inputNums{num1, num2}}
//		ca = &addres
//	case "-":
//		var subres = sub{inputNums{num1, num2}}
//		ca = &subres
//	case "*":
//		var mutilres = mutl{inputNums{num1, num2}}
//		ca = &mutilres
//	case "/":
//		var chu = chu{inputNums{num1, num2}}
//		ca = &chu
//	}
//	return factory(ca)
//}
//
//func factory(c calc) int {
//	return c.Operate()
//}
//
//func main() {
//  var factory optFactory
//   res:=factory.optFactory(10,20,"/")
//   fmt.Println(res)
//}
