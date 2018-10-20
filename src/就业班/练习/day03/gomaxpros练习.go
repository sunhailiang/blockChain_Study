package main

import (
	"runtime"
	"fmt"
)

func main() {
	fmt.Println("cpu",runtime.NumCPU())
	fmt.Println("pros",runtime.GOMAXPROCS(1))
}
