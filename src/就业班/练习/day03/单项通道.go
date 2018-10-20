package main

import "fmt"

func main() {
	var mainchan=make(chan int,2)

	go chanInput(mainchan)
	go chanOut(mainchan)
	<-mainchan
	<-mainchan
}

func chanInput(in chan int)  {
	for i:=0;i<50;i++ {
		fmt.Println("AAAAA",i)
	}
	in<-10
}

func chanOut(out chan int)  {
	for i:=0;i<50;i++ {
		fmt.Println("BBBBBB",i)
	}
	out<-10
}
