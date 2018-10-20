package main

import "fmt"

//============================================
//指针基本概念
//指针：就是地址，指向内存地址空间，这个地址往往是在内存中存储的另一个变量的值的起始位置
//指针变量：即存储地址的变量

//go指针特点
//1,默认值:nil
//2,&取变量地址
//3,*取目标对象的值
//4,直接用.来访问成员
/*func main() {
	//代码段01-解释&，*号的应用
	var num = 100  //声明int 类型num
	fmt.Printf("&a=%p\n", &num) //输出：&a=0xc042050080

	//代码段02-声明指针类型
	var p *int=nil
	p=&num
	fmt.Printf("a=%d,*p=%d\n",num,*p)//输出：a=100,*p=100，此处将指针地址赋值给p，用*p访问对象依旧指向同一个值

    //代码段03
	*p=123 //此处直接通过指针修改了对象的值，因为p和a此时指向的对象是同一个，所以a的值也变了
	fmt.Printf("a=%d,*p=%d\n",num,*p)//输出：a=123,*p=123

}*/

//============================================================
//new(T)表达式关于指针的应用
//1,new(T)创建的对象为堆内存上，作用就是创建T类型的匿名变量为该类型变量开辟一块内存空间，然后将这块空间的内存地址作为结果返回。*T
//2,开辟的内存空间的值就是该变量类型的默认值如：new（int）则默认值为0，如果new(bool)则默认值为false
/*func main() {
	var obj *int               //创建int类型的指针
	obj = new(int)             //让*obj指针指向int变量
	fmt.Println("*obj=", *obj) //*obj= 0

	//demo2---演示bool，其他数据类型自行测试
	var obj2 *bool
	obj2 = new(bool)
	fmt.Println("*obj2", *obj2)
	//直接修改指针对象的值
	*obj2 = true
	fmt.Println("*obj2", *obj2)
//注释：我们使用new函数无需担心其内存的生命周期或怎样将其销毁删除，go语言的内存管理系统会帮我们大理一切
}*/

//==================================================
//指针做函数参数
//
/*func main() {
	var a, b = 10, 99
	test(a, b)//输出：a=10,b=99 值类型传参[值类型传参本质是将变量的值copy一份作为参数传给函数]
	fmt.Printf("a=%d,b=%d\n", a, b)

	test2(&a, &b)//输出：&a=99,&b=10 指针传参[指针类型传参本质就是将指向具体内存空间地址当作参数传递，形参通过*直接访问变量的值，函数内部堆形参的操作就是在操作内存空中间的变量值]
	fmt.Printf("&a=%d,&b=%d\n", a, b)

}
func test(a, b int) {
	a, b = b, a
}
func test2(a, b *int) {
	*a, *b = *b, *a
}*/

//==================指针常见错误使用方式====================
//空指针和野指针
func main() {
	//1,空指针:即未被初始化的指针
	var obj *int
	//fmt.Println("*p", *obj) //输出：invalid memory address or nil pointer dereference，无效地址，即没有可指向的内存空间

	//空指针解决方案
	obj=new(int)
	fmt.Println("*p",*obj)

	//2，野指针：即指向未知的内存地址空间
	//var p *int
	//*p = 0x042058080//未知空间
	//fmt.Println("*p",*p)

}

//相关拓展知识块================栈帧==============
//1:用来给函数提供内存空间，取内存于stack上
//2:函数调用时产生栈帧，函数调用结束释放栈帧
//3:栈帧的存储：1，局部变量2，形参（两者地位等同）3,内存字段描述值
