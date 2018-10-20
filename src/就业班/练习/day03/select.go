package main

import "fmt"

func main() {
	in:=make(chan int)
	quit:=make(chan bool)
	go selectTest(in,quit)

	go inputTest(in,quit)

	for  {
		;
	}

}

func selectTest(in <-chan int,quit <-chan bool)  {
	for i:=0;i<10;i++{
		select {
		case  num:=<-in:
			fmt.Println("读到数据",num)
		case <-quit:
		  return
		}
	}
}

func inputTest(out chan <- int,quit chan <-bool)  {
	for i:=0;i<10;i++{
		out<-i
		fmt.Println("写入数据",i)
	}
	quit<-true
}