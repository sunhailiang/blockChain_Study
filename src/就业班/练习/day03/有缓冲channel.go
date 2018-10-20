package main

import "fmt"

func main() {
	channel := make(chan int, 50)

	go func() {
		for i := 0; i < 10; i++ {
			channel <- i
			fmt.Println("go chan", len(channel), cap(channel))
		}
		close(channel)
	}()

	//for member:=range channel{
	//	fmt.Println(member)
	//}

	//for i:=0;i<10;i++{
     //  num:=<-channel
     //  fmt.Println("main",num)
	//}
	for  {
		if num,ok:=<-channel;ok==true{
			fmt.Println("输出channel数据",num)
		}else{
			break
		}
	}
	for  {
		;
	}

}
