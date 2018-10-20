package main

import "fmt"

var sum int

func producer(out chan<- int) {
	for i := 0; i <= 100;i++{
		out <- i
	}
	close(out);
}

// 此chanel 只能读，不能写
func consumer(in <-chan int) {
	for num := range in {
		sum += num
	}
	fmt.Println("sum =", sum)
}

func main() {

	ch := make(chan int) // 创建一个双向通道
	go producer(ch)      // 协程1，生产者，生产数字，写入channel
	go consumer(ch)      // 协程2，消费者1
	consumer(ch)         // 主协程，消费者。从channel读取内容打印
	for {
		;
	}
}
