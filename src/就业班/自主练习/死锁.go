package main

import "fmt"

//func main() {
//	ch:=make(chan  int)
//	ch<-10
//	num:=<-ch
//	fmt.Println("num:",num)
//}
//func main() {
//	ch := make(chan int)
//	//放在go程下面即可
//	fmt.Println("num",1)
//	<-ch
//    //位置不同
//	go func() {
//		ch <- 10
//	}()
//
//}

func main()  {
	ch1:=make(chan  int)
	ch2:=make(chan int)

	go func() {
		for {
			select {
			  case num:=<-ch2:
			  	ch1<-2
			  	fmt.Println("num1",num)
			}
		}
	}()

	for {
		select {
		case num:=<-ch1:
			ch2<-2
			fmt.Println("num2",num)
		}
	}
}
