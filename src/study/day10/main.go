package main

import (
	"strconv"
	"fmt"
)

func main() {
	content := make([]byte, 0, 1024)
	fmt.Println(string(strconv.AppendBool(content, true)))
	fmt.Println(string(strconv.AppendFloat(content, 3.1415926, 'f', 4, 64)))
}
