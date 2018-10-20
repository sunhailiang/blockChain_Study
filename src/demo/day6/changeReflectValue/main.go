package main
import (
	"reflect"
	"fmt"
)
func main() {
	//=====基本类型操作=====
	var i int
	//整型
	changeValue(&i)
	fmt.Println(i)
	//字符串
	var s string = "fuck"
	changeValue(&s)
	fmt.Println(s)
	//bool
	var b bool = true
	changeValue(&b)
	fmt.Println(b)
	//float
	var f float64 = 3.33
	changeValue(&f)
	fmt.Println(f)

	//===结构体操作====
	//如果想通过反射修改结构体的值因为跨包，所以变量要大写
	var p Person = Person{
		Gender: "男",
		Name:   "harry",
		Age:    27,
	}
	//传递引用类型
	changeValue(&p)
	fmt.Println(p)
	//此时相关的值也会被改变
	//p.Getinfo(p.Name, p.Age, p.Gender)
}

type Person struct {
	Gender string `json:"person_gender"`
	Name   string  `json:"person_name"`
	Age    int   `json:"person_age"`
}
func (p Person) Getinfo(name string, age int, gender string) {
	p.Age = age
	p.Gender = gender
	p.Name = name
	fmt.Println("===reflect use methods start===")
	fmt.Printf("姓名：%s,性别：%s,年龄：%d\n", p.Name, p.Gender, p.Age)
	fmt.Println("===reflect use methods end===")
}
func changeValue(obj interface{}) {
	//反射获取对象的成员
	var v = reflect.ValueOf(obj).Elem()
	//判断数据类型，专门处理结构体
	if v.Kind() == reflect.Struct {
		//如果是个结构体获取结构体字段数量
		num := v.NumField()
		//获取方法数量
		methoeds:=v.NumMethod()
		fmt.Printf("此结构体共有字段%d个\n", num)
		fmt.Printf("此结构体共有方法%d个\n", methoeds)
		fmt.Println("value", v)
		//通过名称匹配要修改的字段,并且设置相关的值
		v.FieldByName("Gender").SetString("女")
		v.FieldByName("Age").SetInt(18)
		v.FieldByName("Name").SetString("杨无敌")
        //获取json打包的Tag描述---Tag在TypeOf这个函数下面
		for i:=0;i<v.NumField();i++  {
			fmt.Println(reflect.TypeOf(obj).Elem().Field(i).Tag.Get("json"))
		}
		//直接使用结构体内的方法
		var params []reflect.Value
		params=append(params,v.FieldByName("Gender"))
		params=append(params,v.FieldByName("Age"))
		params=append(params,v.FieldByName("Name"))
		//重点，此处的call如果调用的函数有参数，则需要把参数填入，没有则用nil
		v.Method(0).Call(params)
	} else {
		switch v.Kind() {
		case reflect.String:
			v.SetString("不处理字符串")
		case reflect.Int:
			v.SetInt(1234567)
		case reflect.Bool:
			v.SetBool(false)
		case reflect.Float64:
			v.SetFloat(99.999)
		}
	}
}
