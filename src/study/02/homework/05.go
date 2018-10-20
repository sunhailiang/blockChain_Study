package main

import (
	"fmt"
)

func main()  {
	var x int
	var y int
	fmt.Println("请输入起始数据")
	fmt.Scanf("%d",&x)
	fmt.Println("请输入结尾数据")
	fmt.Scanf("%d",&y)
	getNum(&x,&y)
}
func getNum(x *int,y *int)  {
	if *x<*y{
		Num(*x,*y)
	}else{
		Num(*y,*x)
}
}

func  Num(small int,big int)  {
	 arr:=[]int{}
	 var max int
	for i:=1;i<=small;i++  {
		if small%i==0&&big%i==0{
			arr=append(arr,i)
		}
	}
    for i:=0;i< len(arr);i++{
    	if max<arr[i]{
    		max=arr[i]
		}
	}
	fmt.Println("最大公约数是:",max)
}