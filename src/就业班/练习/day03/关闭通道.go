package main

import "fmt"

//全局通道
var mainCh=make(chan  bool,2)
func main() {

	go func() {
		i:=0
		for {
			if i>50{
				//close(mainCh)
				break
			}
			fmt.Println(i)
			i++
		}
		mainCh<-true
		mainCh<-true
	}()


	go func() {
	   	  fmt.Println("BBB")
		   <-mainCh
   }()
	<-mainCh
  fmt.Println("over")
}
