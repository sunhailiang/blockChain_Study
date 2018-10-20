package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)
	o := make(chan bool)
	go func() {
		for {
			select {
			case v := <-c:
				fmt.Println(v)
			case <-time.After(5 * time.Second)://只要满足条件，即可执行解除阻塞
				fmt.Println("timeout")
				o <- true
				break
			}
		}
	}()
	//c <- 666 // 注释掉，引发 timeout
	<-o
}
