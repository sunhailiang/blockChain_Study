package main

import (
	"fmt"
)

func main() {
	var str = "qwqwyuyu"
	var newStr = useFor(str)
	fmt.Println("forStr", newStr)
	var newStr2,strMap= useMap(str)
	fmt.Println("str", newStr2,strMap)
}

//利用for
func useFor(str string)string{
	var byts = []byte(str)
	var strArr string
	for i := 0; i < len(byts); i++ {
		flag := true
		for _, v := range strArr {
			if string(v) == string(byts[i]) {
				flag = false
				break
			}
		}
		if flag {
			strArr += string(byts[i])
		}
	}
	return strArr
}

//利用Map
func useMap(str string) (string,map[string]int) {
	result := ""
	var strMap=make(map[string]int,99)
	temp := map[string]byte{}
	for _, v := range str {
		length := len(temp)
		//此处不重复的话就新添加key，key唯一
		temp[string(v)] = 'a'
		if len(temp) != length {
			result += string(v)
			strMap[string(v)]=1
		}else{
		strMap[string(v)]++
		}
	}
	return result,strMap
}
