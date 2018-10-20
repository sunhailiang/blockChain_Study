package main

import "fmt"

//============================数组===================================
//func main() {
//数组：相同数据类型的集合
//数据类型:值类型，存于栈
//===================

//基础定义
//var  intArr [10]int //定义一个长度为10的int类型的数组,当然了也可以换成其他类型，string，bool都行...
//定义完成后，在内存中开辟了10个连续的内存空间，每个数据都存在相应的空间内，每个数据称为数组元素，数组包含的元素个数成为数组的长度
//内部函数len()来获取数据的长度
//arrLen:=len(intArr)//arrlen就是数组的长度源码：返回int类型：func len(v Type) int
//intArr[0]=0
//intArr[1]=1
//intArr[2]=2
//intArr[3]=3
// .
// .
// .
// int[9]=9
//根据数组下标，依次赋值...
//下标的范围是从0开始到数组长度减1的位置
//如果不懂下标，自行学习...
//===========================
//当然了多次赋值比较烦，在特定情况下也可以借助for循环来赋值
//for i:=0;i< len(intArr);i++{
//	intArr[i]=i
//}

//1,数组遍历输出
//for i:=0;i<len(intArr);i++{
//	fmt.Printf("intArr[%d]=%d\n",i,intArr[i])//同样也是通过下标输出每一项的值
//}
//数组遍历方法2 rang
//for index,value:=range intArr{
//	 fmt.Printf("下标index:%d,对应下标的值:%d\n",index,value)
//}

//带有初始化值的数组定义
//var  strArr =[10]string{"A","B","C","D","E","F","G","H","I","J"}
//===================经典案例：冒泡排序，了解一下========================
//var  intArr =[10]int{2,44,3,6,7,1,4,11,34,55}
//for i:=0;i< len(intArr);i++{
//	for j:=0;j<len(intArr)-1-i;j++ {
//		if intArr[j]<intArr[j+1] {
//			intArr[j],intArr[j+1]=intArr[j+1],intArr[j]
//		}
//	}
//}
//fmt.Println(intArr)

//关于数组传参，有区别于其他语言，go中数组是个值类型，默认传参都是copy一份数组的值给形参，并不会影响原数组，如果想直接通过自定义函数直接处理函数外的数组需要借助指针
//如：
//var strArr=[9]string{"a","b","c","d","e","f","g","h","i"}
//fmt.Println("反转：",reverse(strArr))
//}
//数组反转
//func reverse(strArr [9] string)[9]string{
//	for i:=0;i< len(strArr)/2;i++ {
//		timp:=strArr[i]
//		strArr[i]=strArr[len(strArr)-i-1]
//		strArr[len(strArr)-i-1]=timp
//	}
//	return strArr
//}

//=====================切片=============================
//不是一个数组的指针，是一种数据结构体，用来操作数组内部元素。	runtime/slice.go
// type slice struct {
//       	*p
//	       len
//	      cap
//}
func main() {
	//1，目前在大多数go开发过程中都使用切片来替代数组
	//2，切片数组的区别 1，数组定长
	//2，切片不定长且可以追加元素，追加时自动扩容，可以看成动态数组
	//3，定义：数组var num [10] int 指定固定长度
	//         切片var num []int   不指定长度随意追加成员
	//4，切片的元素追加
	//var num []int
	//num = append(num, 10, 11, 23, 44, 55,44,89,55,66,66,34,54,77) //追加多个成员
	//fmt.Println("num", num)
	//fmt.Printf("长度%d\n", len(num))//长度即成员个数 len()获取数组长度,此处说的是已经存储的空间
	//fmt.Printf("容量%d\n", cap(num))//容量扩容，每次扩容长度+2 cap()获取数组容量，所谓容量指的是整个切片已经赋值和空闲的存储空间

	//切片定义方式：
	//slice:=[]int{}
	//fmt.Printf("空切片的默认值是:%v",slice)
	//定义方式2
	//var slice2 []int
	//fmt.Printf("空切片的默认值是:%v",slice2)
	//定义方式3
	// slice3:=make([]int,0) //用make定义充分解释了切片是应用类型数据，在堆内存开辟空间
	//fmt.Printf("空切片的默认值是:%v",slice3)

	//1，给切片赋值，可以像数组一样通过下标
	//2，可以通过append追加
	// slice:=[]int{}
	//slice=append(slice, 1,2,3,4,5,6)

	//切片截取
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	//1，从头开始到指定位置 slice[:n]---包左不包右
	fmt.Println("从开始截取到下标为2的位置", slice[:2]) //输出：[1 2]
	//2，从指定位置到结尾slice[n:]
	fmt.Println("从下标为2的位置到最后", slice[2:]) //输出：[3 4 5 6 7 8 9 0]
	//3，从指定位置开始到指定位置结束slice[n,m]
	fmt.Println("从下标为2到下标5中间所有", slice[2:5]) //输出：[3 4 5]
	//4,将所有的都切了
	fmt.Println("全部转为切片", slice[:]) //输出：[1 2 3 4 5 6 7 8 9 0]

	//常用函数append()追加成员,copy()复制数据
	//copy（目标位置切片， 源切片）
	//s1 := slice[1:5]
	//s2 := slice[7:]
	//copy(s1, s2)//
	//fmt.Println("s1", s1)

	//切片做参数：
	//1,切片是个引用类型，作为参数时传递的是切片的内存地址
	sliceTest(slice)           //此处传递地址
	fmt.Println("反转结果", slice) //输出：[0 9 8 7 6 5 4 3 2 1] 将切片地址传进去函数内部操作直接影响原切片

}
func sliceTest(slice []int) {
	for i := 0; i < len(slice)/2; i++ {
		time := slice[i]
		slice[i] = slice[len(slice)-i-1]
		slice[len(slice)-i-1] = time
	}
}
