package main
import (
	"fmt"
	"runtime"
)
func main() {
	go func() {
		defer fmt.Println("A.defer")

		func() {
			defer fmt.Println("B.defer")
			runtime.Goexit() // 终止当前 goroutine, import "runtime"
			fmt.Println("B") // 不会执行
		}()
		fmt.Println("A") // 不会执行
	}() 	//不要忘记()
	//死循环，目的不让主goroutine结束
	for  {
	}
}

