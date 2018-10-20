package main

import (
	"fmt"
)

func main()  {
	var m int
	var n int
	fmt.Println("请输入起始数据")
	fmt.Scanf("%d",&m)
	fmt.Println("请输入结尾数据")
	fmt.Scanf("%d",&n)
    getSum(&m,&n)
}

func getSum(m *int,n *int)  {
	if(*n>*m){
       Sum(*n,*m)
	}else{
		Sum(*m,*n)
	}
}

func Sum(n int,m int)  {
	var sum int
	for i:=m;i<=n;i++{
		if i%7!=0&&i%5!=0{
			sum+=i
		}
	}
	fmt.Println("sum:",sum)
}