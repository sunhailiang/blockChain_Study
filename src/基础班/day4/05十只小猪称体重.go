package main

import "fmt"

func main() {
	arr:=[10]int{9,1,5,6,7,3,10,2,4,8}


	max:=arr[0]

	for i := 1; i < len(arr); i++ {
		if max < arr[i] {
			max=arr[i]
		}
	}
	fmt.Println("最重的小猪体重：",max)
}
