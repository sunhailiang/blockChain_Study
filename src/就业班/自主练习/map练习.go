package main

import (
	"strings"
	"fmt"
)

func main() {
	//"I love my work and I love my family too"
	str := "I love my work and I love love my family too love"
	res := wcFunc(str)
	fmt.Println(res)
}

//func wcFunc(paramStr string) map[string]int {
//	strArr := strings.Split(paramStr, " ")
//	strMap := make(map[string]int)
//	for i := 0; i < len(strArr); i++ {
//		flag := len(strMap)
//		strMap[strArr[i]] = 1
//		currentFlag := len(strMap)
//		if currentFlag == flag {
//			strMap[strArr[i]]++
//		}
//	}
//	return strMap
//}

func wcFunc(paramStr string) map[string]int {
	strArr:=strings.Fields(paramStr)
	strMap := make(map[string]int)
	for i,_:=range strArr {
		strMap[strArr[i]]++
	}
	return  strMap
}
