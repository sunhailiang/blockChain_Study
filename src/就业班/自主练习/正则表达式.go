package main

import (
	"regexp"
	"fmt"
)

func main() {
	str := "sdasd.fgsdf212f5g6y6hf3.32 3.43 3. 45.66"
	rexp := regexp.MustCompile(`\d+\.\d+`)
	res := rexp.FindAllStringSubmatch(str, -1)
    fmt.Println("res",res)
}
