package main

import (
	"reflect" //引入反射包
	"fmt"
)
type person struct  {
	name string
	gender string
}

func test(obj interface{})  {
	//获取对象类型
	t:= reflect.TypeOf(obj)
	fmt.Println("类型是:",t)
	//获取对象的值
	fmt.Println("对象值是：",reflect.ValueOf(obj))
	//获取值的类型
	v:=reflect.ValueOf(obj)
	k:=v.Kind()
	fmt.Println("对象值类型是：",k)
}
func main()  {
	var p person=person{
		name:"harry",
		gender:"男",
	}
	test(p)
}
