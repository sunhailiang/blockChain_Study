package main

import "fmt"

func main() {
	arr:=[10]string{"a","b","c","d","e","f","g","h","i","j"}
	start:=0
	end:= len(arr)-1
	for {
		if start>end {
			break
		}
		arr[start],arr[end]=arr[end],arr[start]
		start++
		end--
	}
	fmt.Println(arr)
}
