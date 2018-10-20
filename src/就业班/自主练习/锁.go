package main

import (
	"sync"
	"runtime"
	"fmt"
	"strings"
)

var mutex sync.Mutex

func main() {
	var ch = make(chan int)
	runtime.GOMAXPROCS(2)
	i:=0
	for  {
		i++
		if i>100 {
			break
		}
		go person1(ch)
		go person2(ch)
	}

	for  {
		;
	}
}

func printer(str string) {
	mutex.Lock()
	for _, v := range str {
		fmt.Println(strings.ToUpper(string(v)))
		//time.Sleep(time.Millisecond * 100)
	}
	mutex.Unlock()
}

func person1(ch <-chan int) {
	printer("aabb")
	//<-ch
}
func person2(ch chan<- int) {
	printer("vvww")
	//ch <- 10
}
