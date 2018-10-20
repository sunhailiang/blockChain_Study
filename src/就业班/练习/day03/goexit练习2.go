package main

import (
	"fmt"
	"runtime"
)

func main() {
	go func() {
		j:=0
		for  {
			j++
			if j>30 {
				break
			}
			test2()
			fmt.Println("go",j)
			defer fmt.Println("协程结束")
		}
		fmt.Println("")
	}()
	for {
		;
	}
}
func  test2()  {
	i:=0

	for {
		i++
		if i>50 {
			break
		}
		if i>30 {
			 runtime.Goexit()
		}
		fmt.Println("test",i)
	}

	defer fmt.Println("执行完毕")
}
