package main

import "fmt"

func main() {
	arr:=[]int{1,2,3,4,5}
	//fmt.Printf("栈：实参地址:%p\n",&arr)
	fmt.Printf("堆：实参地址:%p\n",*&arr)
	test002(&arr)
	//fmt.Printf("栈：修改后的地址:%p\n",*&arr)
	fmt.Printf("堆：修改后的地址:%p\n",*&arr)
	fmt.Println("arr",arr)
}
//传入切片:此时二级指针堆地址都是一样的，此时发生对传进切片修改，因为堆地址一样，所以原切片同时改变
//传入切片地址：此时在函数中指向堆地址一致，一旦发生append扩容，地址指向堆的空间就会发生变化
func test002(a *[]int){
	(*a)[0]=3333  //-->0xc04207e030
	//fmt.Printf("栈：形参地址:%p\n",*a)
	fmt.Printf("堆：形参地址:%p\n",*a)
	//a=append(a,22,22,22,22)//---->0xc042084050  //此处append数据地址发生了变化，与传进来的堆地址发生了变化
	*a=append(*a,3,4,5,6,7)
	//(*a)[0]=3333//---->0xc042084050
	fmt.Printf("堆：扩容前的形参地址:%p\n",*a)
	//fmt.Printf("栈：扩容后的形参地址:%p\n",*a)
	fmt.Printf("堆：扩容后的形参地址:%p\n",*a)
}
