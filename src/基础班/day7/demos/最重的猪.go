package main

import "fmt"

func main() {
	arr:=[10]int{9,1,5,6,7,3,10,22,4,8}
    max:=arr[0]
	for i:=0;i< len(arr); i++{
		if max<arr[i] {
			max=arr[i]
		}
	}
	fmt.Println(max)
	
}
