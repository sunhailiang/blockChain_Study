package main

import "fmt"

func main() {
	var strArr=[]string{"人","好","是","我"}

	 var  start int
	 var end = len(strArr)-1

	for{
		if start>end {
			break
		}
		strArr[start],strArr[end]=strArr[end],strArr[start]
		start++
		end--
	}
	fmt.Println(strArr)

}
