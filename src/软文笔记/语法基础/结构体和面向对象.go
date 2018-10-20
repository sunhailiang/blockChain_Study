package main

import "fmt"

//结构体
//一种数据类型，其地位等价于int string  bool...
//通常全局定义
//可以使用==   !=进行比较
//相同类型的结构体可以相互赋值（类型，个数，顺序一致）
//结构体的变量地址等于首个成员的变量地址

//定义结构体
type person struct {
	name string
	age  int
	id   int
	job  string
}

/*func main() {
	//初始化自定义结构体类型的对象
	//1，顺序初始化，即按照成员位置依次赋值
	/*p := person{"harry", 18, 1} //顺序初始化数据
	fmt.Println("姓名:", p.name, "年龄:", p.age, "编号", p.id)
	//开辟存储空间来赋值
	p2 := new(person)
	p2.name = "fuck"
	p2.age = 19
	p2.id = 2
	fmt.Println("姓名:", p2.name, "年龄:", p2.age, "编号", p2.id)
	//通过指针索引成员变量赋值
	var p3 person
	p3.id = 3
	p3.age = 20
    p3.name="harry"
	fmt.Println("姓名:", p3.name, "年龄:", p3.age, "编号", p3.id)



	//结构体传参
	//1，结构体属于值类型，常规传参都是传的结构体的值，但是结构体大小难定，性能差
	//2，常用结构体传参都是采用地址传递，这样我们只需要传递一个8字节的指针地址即可
 	var p person
	structTest(&p)//此处传的地址用来操作结构体
	fmt.Println(p)//输出{欧阳锋 45 4}

}

func structTest(obj *person)  { //*号通过地址直接访问只针对象
	  obj.name="欧阳锋"
	  obj.id=4
	  obj.age=45
}
*/

//面向对象
//1,封装，所谓封装即函数封装，目的是将有特定功能的代码块写成函数体，以便多方使用，提高代码复用率，灵活性，降低代码冗余
//所谓函数封装,main函数就是一个函数体
//2，继承
//继承通过匿名字段来实现
//也就是直接使用父结构体
//3，重写-即多态体现

//定义接口
type Handler interface {
	study()
}

//定义结构体
type student struct {
	person //此处通过一个父类的匿名字段就可以继承父类
}

//定义警察类型结构体
type policeMan struct {
	person //此处通过一个父类的匿名字段就可以继承父类
}

//方法
//所谓方法即某结构体具有的某些功能，在代码体现上就是将函数值给某个结构体
//实现接口方法
func (p *person) study() { //此处将函数挂载到person上也就是person这个结构体的方法
	fmt.Printf("我%d号学生，我叫%s，我今年%d岁,我是一名%s\n", p.id, p.name, p.age, p.job) //通过访问student结构体就能获得父类的属性体现继承
}
func main() {

	var i Handler
	s := &student{person{"harry", 28, 1, "学生"}}
	i = s
	i.study() //调用接口使用同一个函数，多态
	//此处在多个子类同时调用父类的函数进行重写，这就是一种多态的体现，即同一个函数能展示子类不同的内容
	var p = &policeMan{person{"赵日天", 28, 2, "警察"}}
	i = p
	i.study() //调用接口使用同一个函数，多态
}
