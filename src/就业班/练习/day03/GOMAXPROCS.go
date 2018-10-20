package main

import (
	"fmt"
	"runtime"
)

func main() {
	//n := runtime.GOMAXPROCS(1) 	// 设置最大cpu核心1也就是goroutine按照go调度器的安排交替时间段执行..
	//打印结果：111111111111111111110000000000000000000011111...

	n := runtime.GOMAXPROCS(2) //次数设置cpu核数2也就是存在2个并行任务，所以存在01010101
	//打印结果：010101010101010101011001100101011010010100110...
	fmt.Printf("n = %d\n", n)

	for {
		go fmt.Print(0)
		fmt.Print(1)
	}
}
