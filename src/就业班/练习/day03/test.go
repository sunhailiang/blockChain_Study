package main

import "fmt"

func main() {
	for i := 0; i < 50; i++ {
		fmt.Println("主程", i)
		go func(num *int) {
			fmt.Println("协程", *num)
		}(&i)
	}
	for {
		;
	}
}
