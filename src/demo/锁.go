package main

import (
	"fmt"
	"sync"
	"math/rand"
	"time"
)
//同步组
var wg sync.WaitGroup
//互斥锁
var matex sync.Mutex

func main() {

	wg.Add(3)
		go win("一")
		go win("二")
		go win("三")
		go win("四")

	wg.Wait()
	fmt.Println("结束")

}
var ticket=100
func win(name string) {

	//设置随机数种子
    rand.Seed(time.Now().UnixNano())
    for{
    	//互斥锁，相当于同步队列
		 matex.Lock()
    	if ticket>0{
    		 time.Sleep(time.Millisecond*time.Duration(rand.Intn(1000)))
    		 fmt.Println(name+"出售:",ticket)
    		 ticket--
		}else{
			fmt.Println("售票结束")
			//售票完全结束后解锁
			matex.Unlock()
			break
		}
		//单次任务解锁
		matex.Unlock()
	}
	wg.Done()
}
