package main

import (
	"sync"
	"fmt"
	"math/rand"
	"time"
)

var rwlock sync.RWMutex //全局读写锁
var val int

func read(idx int) {
	for {
		rwlock.RLock() //添加读锁
		num := val
		fmt.Printf("%d号go程读出数据%d\n", idx, num)
		time.Sleep(time.Second)
		rwlock.RUnlock()

	}
}

func write(idx int) {
	for {
		num := rand.Intn(100)
		rwlock.Lock() //添加写锁
		val = num
		fmt.Printf("%d号go程写入数据%d\n", idx, val)
		time.Sleep(time.Millisecond*400)
		rwlock.Unlock()
	}
}

func main() {

	rand.Seed(time.Now().UnixNano())

	for i:=0;i<5;i++{
		go read(i+1)
	}
	for i:=0;i<5;i++{
		go write(i+1)
	}

	for  {
		;
	}

}
