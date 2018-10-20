package main
import (
	"demo/day1/goroute/calc"
	"fmt"
)
func main(){
	//声明管道
	var pipe chan int
	//初始化管道
	pipe=make(chan int,1);
	 
	go calc.Add(100,11,pipe)
	
	sum:=<-pipe;
	fmt.Println("sum",sum)


}