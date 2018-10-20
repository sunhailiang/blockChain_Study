package main

import (
	"fmt"
	"runtime"
)

func main() {
	go func(str string) {
		for i:=0;i<2;i++{
			fmt.Println(str)
		}
	}("hello")

	for i:=0;i<2;i++{
		//1，让出cup时间片让其他任务执行，其他任务执行后再回到当前任务
		//2，如果没加，上面子协程还没执行主goroutine已经结束了
		runtime.Gosched()
		fmt.Println("word")
	}
}
