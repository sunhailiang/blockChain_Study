package main

import (
	"math/rand"
	"time"
	"fmt"
)
func main() {
	//猜数字
	//设置随机数种子
	rand.Seed(time.Now().UnixNano())
	//产生随机数
	var randNum = rand.Intn(899) + 100
	//放置随机数数组
	var randArr [3] int
	//放手动输入的数字
	var userNum [3]int
	randArr[0] = randNum / 100
	randArr[1] = randNum % 100 / 10
	randArr[2] = randNum % 10
	var num int
	var flag int
	for {
		for {
			fmt.Println("请输入合法三位数")
			fmt.Scan(&num)
			if num >= 100 && num <= 999 {
				break
			}
			fmt.Println("输入有误，重新输入")
		}
		userNum[0] = num / 100
		userNum[1] = num % 100 / 10
		userNum[2] = num % 10
		for i := 0; i < 3; i++ {
			if randArr[i] > userNum[i] {
				fmt.Printf("您输入的%d位过小\n", i+1)
			} else if randArr[i] < userNum[i] {
				fmt.Printf("您输入的%d位过大\n", i+1)
			} else {
				fmt.Printf("恭喜您的第%d位相同\n", i+1)
				flag++
			}
		}
		if flag == 3 {
			fmt.Println("恭喜您,全部猜对了")
			break
		}else{
			flag=0
		}
	}
}
