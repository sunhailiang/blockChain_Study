package main

import "fmt"

func main() {
	//cock 公鸡 5钱
	//hen 母鸡 3钱
	//chicken 小鸡 1/3钱

	chicken:=0
	count:=0
	for cock:=0;cock<=20;cock++{
		for hen:=0;hen<=33;hen++{
			count++
			//计算剩余小鸡个数 100-公鸡-母鸡
			chicken=100-cock-hen
			if cock*5+hen*3+chicken/3==100 && chicken%3==0{
				fmt.Printf("公鸡：%d 母鸡：%d 小鸡：%d\n",cock,hen,chicken)
			}
			//for chicken:=0;chicken<=100;chicken+=3{
			//	count++
			//	if (cock+hen+chicken==100) && (cock*5+hen*3+chicken/3==100){
			//		//fmt.Println("公鸡：",cock)
			//		//fmt.Println("母鸡：",hen)
			//		//fmt.Println("小鸡：",chicken)
			//		fmt.Printf("公鸡：%d 母鸡：%d 小鸡：%d\n",cock,hen,chicken)
			//
			//	}
			//}
		}
	}

	fmt.Println("执行次数：",count)
}
