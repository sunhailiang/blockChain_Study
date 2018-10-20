package main
import "fmt"
//var ch=make(chan bool)
var done=make(chan bool)
func main() {
	//同时打印*，数字，字母
	go printOther()
	go printNum()
	go printA()
	//三个函数都执行完后都有对应的读出解除阻塞，直到所有的函数都执行完
	<-done
	<-done
	<-done
}
func printNum()  {
	var num=0
	for  {
		if num>99 {
			break
		}
		num++
		fmt.Println(num)
	}
	//写入并阻塞
	done<-true
}
func printOther()  {
	for i:=0;i<100;i++{
		fmt.Println("*")
	}
	//写入并阻塞
	done<-true
}
func  printA()  {
	for j:=0;j<100;j++{
      fmt.Println("A")
	}
	//写入并阻塞
	done<-true
}
