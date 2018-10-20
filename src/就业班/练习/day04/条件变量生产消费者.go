package main

import (
	"sync"
	"math/rand"
	"fmt"
	"time"
)
//定义条件变量
var cond sync.Cond

func main() {
	//添加随机数种子
	rand.Seed(time.Now().UnixNano())
	//创建互斥锁和条件变量
	cond.L = new(sync.Mutex)
	//主程阻塞
	var quit = make(chan bool)
	//创建双向通道
	var produce = make(chan int, 4)
	//创建五个消费者
	for i := 0; i < 4; i++ {
		go Consumer(produce, i+1)
	}
	//创建3个生产者
	for i := 0; i < 3; i++ {
		go Producer(produce, i+1)
	}
	//阻塞服务器
	<-quit
}

func Producer(in chan<- int, proId int) {
	for {
		cond.L.Lock() //加一把锁
		for len(in) ==4 {
			cond.Wait() //如果缓冲区数据满了，挂起
		}
		//如果没有挂起则继续添加数据
		num := rand.Intn(1000)
		in <- num
		fmt.Printf("%d生产者生产数据%d，缓存区当前还有数据%d\n", proId, num, len(in))
		//缓冲区没满，解锁
		cond.L.Unlock()
		//唤醒阻塞的消费者//广播我已经完成
		cond.Signal()
		//生产完成休息一会
		time.Sleep(time.Second)
	}
}
func Consumer(out <-chan int, cons int) {
	for {
		cond.L.Lock()
		//如果缓冲区没数据，挂起
		for len(out) == 0 {
			cond.Wait()
		}
		//如果有数据就读出
		num := <-out
		fmt.Printf("消费者%d,消费数据%d,缓冲区剩余数据%d个\n", cons, num, len(out))
		//解锁
		cond.L.Unlock()
		//唤醒
		cond.Signal()
		time.Sleep(time.Second)

	}
}
