package main

import (
	"fmt"
	"time"
)

func main() {


	go test()

	for i:=0;i<10; i++{
		fmt.Println("main j",i)
		time.Sleep(time.Millisecond*400)
	}

}

func test()  {
	for i:=0;i<10;i++ {
		fmt.Println("test 中的i",i)
		time.Sleep(time.Millisecond*200)
	}
}
