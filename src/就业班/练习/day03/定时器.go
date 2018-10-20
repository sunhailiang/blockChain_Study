package main

import (
	"time"
	"fmt"
)

func main() {
	mytimer:=time.NewTimer(time.Second*5)
	nowTime:= <-mytimer.C
	fmt.Println("NewTimer",nowTime)




	nowTimer2:=<-time.After(time.Second*3)
	fmt.Println("After",nowTimer2)
}
