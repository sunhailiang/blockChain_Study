package main

import "fmt"

func main() {
	//{"red", "", "black", "", "", "pink", "blue"}
	//——> {"red", "black", "pink", "blue"}

	var strArr=[]string{"red", "", "black", "", "", "pink", "blue"}
	strArr=getNoSpaceSlice(strArr)
	fmt.Println(strArr)

}

func getNoSpaceSlice(slice []string) []string{
	var newSlice [] string
	for i:=0;i< len(slice);i++ {
		if slice[i]!="" {
			newSlice=append(newSlice, slice[i])
		}
	}
	return  newSlice
}
