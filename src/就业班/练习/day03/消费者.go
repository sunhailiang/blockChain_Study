package main

import "fmt"

func producer2 (in chan<- int)  {
	for i:=0;i<20;i++ {
		in<-i*i
		fmt.Println("生产数据,",i*i)
	}
	close(in)
}
func  customer2(out <-chan int)  {
	for num:=range out{
		fmt.Println("读出数据",num)
	}
}

func main() {
	ch:=make(chan int)
     go producer2(ch)
	customer2(ch)
}
