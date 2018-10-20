package main

import (
	"fmt"
	"time"
)

func  main()  {
	//声明，并初始化
	intChan:=make(chan int,10)

	 //使用goroute
	 //读
	go read(intChan)
    //写
	go write(intChan)

	time.Sleep(time.Second*10)
}
//写入channel
func write(ch chan int)  {
   for i:=0;i<20;i++{
   	   ch<-i
	}
}
//读channel
func read(ch chan int){
   for{
   	var b int
   	b=<-ch
   	fmt.Println("读数据",b)
   }
}
