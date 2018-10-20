package main

import "fmt"

func main()  {
	var arr=[10]int{22,44,33,66,999,3,1,3,88,90}
	max:=arr[0]
	for i:=1;i<len(arr);i++  {
		if max<arr[i] {
			max=arr[i]
		}
	}
	fmt.Println("max",max)
}
