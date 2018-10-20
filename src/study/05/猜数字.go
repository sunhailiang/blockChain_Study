package main

import (
	"math/rand"
	"time"
	"fmt"
)

func main()  {
	matchNum()
}


//猜数字
func matchNum(){
	//设置随机数种子
	rand.Seed(time.Now().UnixNano())
	//产生随机数切片
	sysNum:=make([]int,3)
	sysNum[0]=rand.Intn(8)+1
	sysNum[1]=rand.Intn(9)+1
	sysNum[2]=rand.Intn(9)+1
    fmt.Println(sysNum)
	//存储用户输入的切片
	userNum:=make([]int,3)
	var num int
	var flag int
	for	{
		//用户输入数字
		for i:=0;i<3;i++{
			fmt.Println("请输入三位数")
			fmt.Scanf("%d",&num)
			if num>99&&num<1000{
				break
			}
			fmt.Println("输入错误")
		}
		//对比
		userNum[0]=num/100
		userNum[1]=num/10%10
		userNum[2]=num%10
		for i:=0;i<3;i++{
			if(userNum[i]>sysNum[i]){
				fmt.Printf("输入的第%d位数据过大\n",i+1)
			}else if(userNum[i]<sysNum[i]){
				fmt.Printf("输入的第%d位数据过小\n",i+1)
			}else{
				fmt.Printf("输入的第%d位数据相同\n",i+1)
				flag++
			}
		}

		if flag==3 {
			break
		}
	}
}

