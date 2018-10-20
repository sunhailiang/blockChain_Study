package main

import "fmt"

func  main()  {
	var m int
	var n int
	fmt.Println("请输入起始数据")
	fmt.Scanf("%d",&m)
	fmt.Println("请输入结尾数据")
	fmt.Scanf("%d",&n)
	sum(&m,&n)
}
func sum(m *int,n *int)  {

	if(*n>*m){
		var sum int
		for i:=*m;i<=*n;i++{
			sum+=i
		}
		fmt.Println("sum",sum)
	}else{
		fmt.Println("输入有误，请重新尝试")
	}
}
