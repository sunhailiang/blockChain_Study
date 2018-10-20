package main

import (
	"time"
	"fmt"
	"runtime"
)

func main() {
	go func() {
		for i:=0;i<100;i++ {
			fmt.Println("first",i)
			//time.Sleep(time.Millisecond*100)
		}
	}()
    go func() {
    	runtime.Gosched()
		for i:=0;i<100;i++ {
			fmt.Println("second",i)
		}
	}()
	time.Sleep(time.Second)
}
