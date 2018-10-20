package main

import "fmt"

//一只公鸡值五钱，一只母鸡值三钱，三只小鸡值一钱，现在要用百钱买百鸡，请问公鸡、母鸡、小鸡各多少只？
//公鸡最多20只 cock
//母鸡最多33只 hen

func main()  {

	chicken:=0
    for cock:=0;cock<20;cock++{
		for hen:=0;hen<33;hen++{
			chicken=100-cock-hen
			 if(cock*5+hen*3+chicken/3==100)&&chicken%3==0{
				 fmt.Printf("公鸡：%d 母鸡：%d 小鸡：%d\n",cock,hen,chicken)
			 }
		}
	}
}