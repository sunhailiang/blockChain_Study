package main

import (
	"fmt"
)

func main() {
	var arr = [10]int{5, 7, 2, 1, 4, 9, 10, 11, 6, 8}
	for i := 0; i < len(arr)-1; i++ {
		for j:=0;j<len(arr)-1-i;j++  {
			if arr[j]>arr[j+1] {
				arr[j],arr[j+1]=arr[j+1],arr[j]
			}
		}
	}
	var arr2 = []int{5, 7, 2, 1, 4, 9, 10, 11, 6, 8,22,3,4,5555,6666}
	//sort.Reverse(arr2)
	fmt.Println(arr2)
}