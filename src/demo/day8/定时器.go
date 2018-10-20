package main

import (
	"time"
	"fmt"
)

func main() {
	t := time.NewTicker(time.Second)
	//每秒循环一次
	for v:=range t.C{
          fmt.Println("hello,",v)
	}
	//不推荐使用，但是使用后注意手动停止
	t.Stop()
}
