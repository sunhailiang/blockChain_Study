package main

import "fmt"

func main() {
	arr:=[10]int{9,1,5,6,7,3,10,2,4,8}
	for i:=0;i< len(arr); i++{
		for j:=0;j< len(arr)-1-i; j++{
			if arr[j]>arr[j+1]{
				arr[j],arr[j+1]=arr[j+1],arr[j]
			}
		}
	}
	fmt.Println(arr)
	
}
