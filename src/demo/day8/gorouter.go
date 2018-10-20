package main

import (
	"runtime"
	"fmt"
)

func main()  {
	//设置Golang运行的CPU核心
	cpuNum:=runtime.NumCPU()
	//
	runtime.GOMAXPROCS(cpuNum)
	fmt.Println("cpu",cpuNum)
}
