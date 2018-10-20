package main

import (
	"math/rand"
	"time"
	"fmt"
)

func main() {

	//var  arr  [2][4] int =[2][4]int{{1,2,3,4},{5,6,7,8}}
	//6个球，1~33不可重复，蓝色球一个，可以与红球重复，共七个数
	var nums [6]int
	//定义随机数种子
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 6; i++ {
		num := rand.Intn(33) + 1

		for j := 0; j < i; j++ {
			if nums[j] == num {
				j = -1
				continue
			}
		}
		nums[i] = num
	}
	fmt.Println(nums)
}
