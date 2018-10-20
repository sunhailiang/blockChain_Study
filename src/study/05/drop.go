package main

import (
	"time"
	"fmt"
	"math/rand"
)
//福利彩票  双色   红球有6个  1-33  不能重复  蓝球 1-16 随机一个
//rand.Intn(33)+1//1-33
func main() {

	start:=time.Now().UnixNano()
	fmt.Println("dddd",start)
	demo()
	end:=time.Now().UnixNano()
	fmt.Println("qq",start)
	fmt.Println("dd",end)

}

func demo(){
	arr := [6]int{}
	var count int
	var index int
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(arr); i++ {
		if i == 0 {
			arr[i] = rand.Intn(33) + 1
		} else {
			arr[i] = rand.Intn(33) + 1
			for {
				for j := 0; j < i; j++ {
					index++
					if arr[j] != arr[i] {
						count++
					}
				}
				if count != i {
					arr[i] = rand.Intn(33) + 1
					count = 0
				}
				if count == i {
					break
				}
			}
		}
	}
	for i:=0;i<50;i++ {
		fmt.Println("d",i)
	}

}
//循环一百万

//dd 1531372701386721200
//dd 1531372704057815700
//循环十万
//dd 1531372749368415400
//dd 1531372749575010200
//循环1万
//dd 1531372769975557300
//dd 1531372770002612700
//循环100
//dd 1531372797245593400
//dd 1531372797246595500
//循环50次
//dd 1531373261611774600
//dd 1531373261611774600

//你的循环不足50...自己想


