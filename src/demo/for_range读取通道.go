package main

import "fmt"

func main() {
	ch := make(chan int)
	testFor(ch)
	for item := range ch {
		fmt.Println("通道取值", item)
	}
}
func testFor(ch chan int) {
	i := 100
	for {
		if i < 0 {
			//要使用for range 获取channel值，必须要使用close显式通知对方写入结束
			close(ch)
			break
		}
		ch <- i
		i++
	}
}
