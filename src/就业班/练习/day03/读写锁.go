package main

import (
	"sync"
	"fmt"
	"math/rand"
	"time"
)

var count int
var rwlock sync.RWMutex //创建全局读写锁
func main() {
	for i := 0; i < 10; i++ {
		go read(i)
	}

	for i := 0; i < 10; i++ {
		go write(i)
	}
	for {
		;
	}
}

func read(count int) {
	rwlock.RLock()
	fmt.Println("正在读取数据...", count)
	num := count
	fmt.Printf("读 goroutine %d 读取数据结束，读到 %d\n", count, num)
	defer rwlock.RUnlock()
}
func write(count int) {
	rwlock.Lock()
	fmt.Println("正在写数据...")
	rand.Seed(time.Now().UnixNano())
	num := rand.Intn(100)
	count = num
	fmt.Printf("写 goroutine %d 写数据结束，写入新值 %d\n", count, num)
	defer rwlock.Unlock()
}
