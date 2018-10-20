package main

import (
	"fmt"
	"time"
	"math/rand"
	"sync"
)

var cond sync.Cond // 定义全局条件变量

func producer08(out chan<- int, idx int) {
	for {
		// 先加锁
		cond.L.Lock()
		// 判断缓冲区是否满
		for len(out) == 5 {
			cond.Wait() // 1. 2. 3.
		}
		num := rand.Intn(800)
		out <- num
		fmt.Printf("生产者%dth，生产：%d\n", idx, num)
		// 访问公共区结束，并且打印结束，解锁
		cond.L.Unlock()
		// 唤醒阻塞在条件变量上的 消费者
		cond.Signal()
		time.Sleep(time.Millisecond * 200)
	}
}

func consumer08(in <-chan int, idx int) {
	for {
		// 先加锁
		cond.L.Lock()
		// 判断 缓冲区是否为空
		for len(in) == 0 {
			cond.Wait()
		}
		num := <-in
		fmt.Printf("-----消费者%dth，消费：%d\n", idx, num)
		// 访问公共区结束后，解锁
		cond.L.Unlock()
		// 唤醒 阻塞在条件变量上的 生产者
		cond.Signal()
		time.Sleep(time.Millisecond * 200)
	}
}

func main() {
	product := make(chan int, 5)
	rand.Seed(time.Now().UnixNano())

	// 指定条件变量 使用的锁
	cond.L = new(sync.Mutex) // 互斥锁 初值 0 ， 未加锁状态

	for i := 0; i < 5; i++ {
		go producer08(product, i+1) // 1 生产者
	}
	for i := 0; i < 5; i++ {
		go consumer08(product, i+1) // 3 个消费者
	}

	for {
		;
	}
}
