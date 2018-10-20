package main

import "fmt"

func main()  {

	fmt.Println("请输入年份")

	var year int

	fmt.Scanf("%d",&year)

	check:=year%400==0||year%4==0&&year%100!=0

	fmt.Println(check)

}