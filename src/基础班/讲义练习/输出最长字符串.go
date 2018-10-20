package main

import (
	"strings"
	"fmt"
)

func main() {
	var strArr = []string{"马龙", "迈克尔乔丹", "科比布mmmmmm 莱恩特", "雷吉dddd米勒", "蒂姆邓肯"}
	var max = len(strings.Split(strArr[0], ""))
	var longeststr = strArr[0]
	fmt.Println()
	for i := 0; i < len(strArr); i++ {
		if max < len(strings.Split(strArr[i], "")) {
			max = len(strings.Split(strArr[i], ""))
			longeststr = strArr[i]
		}
	}
	fmt.Println("最长的", longeststr)
}
