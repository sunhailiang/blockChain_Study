package main

import (
	"fmt"
	"time"
)

func main() {
	var  ch=make(chan int,2)
	var  t=make(chan bool)
    go func() {
		for {
			select {
			case num:=<-ch:
				fmt.Println("num",num)
			case <-time.After(time.Second*5):
				fmt.Println("finish")
				t<-true
				goto lable
			}
		}
		lable:
	}()
	for i:=0;i<2;i++{
		ch<-i
		time.Sleep(time.Second*2)
	}

	<-t
}
