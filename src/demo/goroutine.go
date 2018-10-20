package main

import (
	"fmt"
	"sync"
)

//同步等待组，等待所有的
var wg1 sync.WaitGroup

func main() {
	wg1.Add(3)
	go test()
	go test2()
	go test3()
	//此处开始等待上面的协程全部执行完成后,在执行其他代码
	//counter为0的时候，解除阻塞
	wg1.Wait()
}
func test() {
	for i := 0; i < 100; i++ {
		fmt.Println("*")
	}
	//add函数在放入协程数量后，此函数放在协程后面的代码中，counter-1
	wg1.Done()
}
func test2() {
	for i := 0; i < 100; i++ {
		fmt.Println("$")
	}
	wg1.Done()
}
func test3() {
	for i := 0; i < 100; i++ {
		fmt.Println("@")
	}
	wg1.Done()
}
