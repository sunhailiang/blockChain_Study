package main

import (
	"sync"
	"math/rand"
	"fmt"
	"time"
)

//条件变量：并不保证在同一时刻，在同一时刻仅有一个线程（协程），访问某个数据资源，而是在对应的共享数据的状态发生变化时，通知阻塞在某个条件上的线程（协程）
//条件变量不是锁，不能达到同步作用，因此总是与锁一块使用

//声明 var cond sync.Cond

//3个常用方法，Wait，Signal，Broadcast。

//1)func (c *Cond) Wait()
//该函数的作用可归纳为如下三点：
//a)阻塞等待条件变量满足
//b)释放已掌握的互斥锁相当于cond.L.Unlock()。 注意：两步为一个原子操作。
//c)当被唤醒，Wait()函数返回时，解除阻塞并重新获取互斥锁。相当于cond.L.Lock()

//2)func (c *Cond) Signal()
//单发通知，给一个正等待（阻塞）在该条件变量上的goroutine（线程）发送通知。

//3)func (c *Cond) Broadcast()
//广播通知，给正在等待（阻塞）在该条件变量上的所有goroutine（线程）发送通知。

//声明全局条件变量
var cand sync.Cond

func main() {
	cand.L = new(sync.Mutex)         //使用互斥锁
	rand.Seed(time.Now().UnixNano()) //设置随机数种子
	ch := make(chan int,5)
	//创建多个生产者
	for i := 0; i < 5; i++ {
		go producer(ch, i)
	}
	//创建多个消费者
	for i := 0; i < 5; i++ {
		go consumer(ch, i)
	}
	for {
		;
	}
}

//生产者
func producer(ch chan<- int, index int) {
	for {
		cand.L.Lock()
		for len(ch) == 5 {
			cand.Wait() //挂起，唤醒读取端
		}
		//写入数据
		num := rand.Intn(600)
		ch <- num
		fmt.Printf("生产者%d号写入数据%d\n", index, num)
		cand.L.Unlock()
		cand.Signal()
		time.Sleep(time.Millisecond * 300)
	}
}

//消费者
func consumer(ch <-chan int, index int) {
	for {
		cand.L.Lock()
		for len(ch) == 0 {
			cand.Wait() //挂起并唤醒生产者
		}
		num := <-ch
		fmt.Printf("消费者%d号读出数据%d\n", index, num)
		cand.L.Unlock()
		cand.Signal()
		time.Sleep(time.Millisecond * 300)
	}
}
