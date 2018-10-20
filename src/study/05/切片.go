package main

import "fmt"

func main()  {
	slice:=[]int{1,2,3,4,5,6}


	var s2=make([]int,3)

	copy(s2,slice)

	fmt.Println("sss",s2)



}
