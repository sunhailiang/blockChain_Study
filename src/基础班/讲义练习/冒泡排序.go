package main

import "fmt"

func main() {
	intarr:=[]int{2,31,23,45,12,76,45,312,34,88,97,654,34}

	for i:=0;i< len(intarr);i++{
		for j:=0;j< len(intarr)-i-1;j++{
			if intarr[j]<intarr[j+1]{
				intarr[j],intarr[j+1]=intarr[j+1],intarr[j]
			}
		}
	}
	fmt.Println(intarr)
}
