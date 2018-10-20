package main

import "fmt"

func main() {
	channel := make(chan int, 100)

	go func() {
		for i := 0; i < 20; i++ {
			channel<-i
			fmt.Println("写入",i)
		}
		close(channel)
	}()

	for {
		if num,ok:=<-channel;ok==true {
			fmt.Println("拿到数据",num)
		}else{
			fmt.Println("读完了",num)
			break
		}
	}
}
