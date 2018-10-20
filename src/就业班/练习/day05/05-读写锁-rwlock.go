package main

import (
	"math/rand"
	"time"
	"fmt"
	"sync"
)

var rwMutex sync.RWMutex		// 锁只有一把， 2 个属性 r w

var value int		// 定义全局变量，模拟共享数据

func readGo05(idx int)  {
	for {
		rwMutex.RLock()			// 以读模式加锁
		num := value
		fmt.Printf("----%dth 读 go程，读出：%d\n", idx, num)
		rwMutex.RUnlock()		// 以读模式解锁
		time.Sleep(time.Second)
	}
}

func writeGo05(idx int)  {
	for {
		// 生成随机数
		num := rand.Intn(1000)
		rwMutex.Lock()			// 以写模式加锁
		value = num
		fmt.Printf("%dth 写go程，写入：%d\n", idx, num)
		time.Sleep(time.Millisecond * 300)		// 放大实验现象
		rwMutex.Unlock()
	}
}

func main()  {
	// 播种随机数种子
	rand.Seed(time.Now().UnixNano())

	for i:=0; i<5; i++ {			// 5 个 读 go 程
		go readGo05(i+1)
	}
	for i:=0; i<5; i++ {			//
		go writeGo05(i+1)
	}
	for{
		;
	}
}
