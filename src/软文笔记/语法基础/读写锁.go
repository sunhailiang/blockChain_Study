package main

import (
	"sync"
	"math/rand"
	"fmt"
	"time"
)

//读写锁：类似互斥锁，但是读写锁并行性更高，读共享，写独占
//锁只有一把：读，写是这把锁两种状态
//写加锁：其他所有加锁线程都会被阻塞
//读操作：支持并发
//读写锁由结构体类型sync.RWMutex表示
/*一组是对写操作的锁定和解锁，简称“写锁定”和“写解锁”：
func (*RWMutex)Lock()
func (*RWMutex)Unlock()
“读锁定”与“读解锁”：
func (*RWMutex)RLock()
func (*RWMutex)RUloc()
*/

var rwmuutex sync.RWMutex //全局读写锁
var count int             //全局变量
func main() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 20; i++ {
		go write(i + 1)
	}
	for i := 0; i < 20; i++ {
		go read(i + 1)
	}
	for {
		;
	}
}

func write(index int) {
	rwmuutex.Lock()
	count = rand.Intn(1000)
	num := count
	fmt.Printf("goroutine %d 写入数据%d\n", index, num)
	rwmuutex.Unlock()
	time.Sleep(time.Millisecond * 300)
}

func read(index int) {
	rwmuutex.RLock()
	num := count
	fmt.Printf("goroutine %d读取数据%d\n", index, num)
	rwmuutex.RUnlock()
	time.Sleep(time.Millisecond * 300)
}
