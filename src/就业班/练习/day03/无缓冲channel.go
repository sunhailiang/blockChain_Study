package main

import "fmt"

var mainChan=make(chan  int,10)
func main() {
	go test3()
	go test4()

	<-mainChan
	<-mainChan
	fmt.Println("main over")
}

func test3() {
	i := 0
	for {
		if i > 50 {
			break
		}
		fmt.Println("test", i)
		i++
	}
	mainChan<-10
}
func test4() {
	i := 0
	for {
		if i > 50 {
			break
		}
		fmt.Println("BBBB")
		i++
	}
	mainChan<-11
}
