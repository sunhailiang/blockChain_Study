package main

import "fmt"

//map
//(映射，字典)是一种内置的数据结构，无序的键值对集合
//key唯一

func main() {
	//map基础定义：
	// 数据类型:引用类型
	//容量不够自动扩容
	//方式1
	//var mp map[int]string //此只声明，并未初始化，此时没有指向的内存空间不能使用
	//方式2
	//mp:=map[int]string{}//此处初始化数据为空的map
	//方式3
	//mp:=make(map[int]string,4)//此处第二个参数指定长度

	//map赋值
	//通过下标赋值
	//mp:=make(map[int]string)
	//mp[0]="a"
	//mp[1]="b"
	//fmt.Println("a",mp)

	mp := map[int]string{1: "a", 2: "b", 3: "c", 4: "d", 5: "e"}
	//删除
	// map的元素：delete()操作是安全的，即使元素不在map中也没有关系；如果查找删除失败将返回value类型对应的零值。
	//delete(mp, 0) //删除map中的元素
	//map遍历
	//for k, v := range mp {
	//	fmt.Println(k, v) //输出：1 a  2 b  3 c  4 d  5 e
	//}

	//map做函数参数
	//map是引用类型所以传递的是指针地址，直接操作内存地址存储的map对象
	mapTest(mp)
	fmt.Println(mp)//输出：map[1:a 2:b 3:c 4:d 5:e 0:你看看是不是改了原map的成员了？？？]

}

func mapTest(mp map[int]string)  {
     if len(mp)>0{
     	mp[0]="你看看是不是改了原map的成员了？？？"
	 }
}
