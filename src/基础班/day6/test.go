package main

import (
	"fmt"
)

func main()  {
	var str="asdfghjklasdghjkl"

	var cMap=make(map[string]int)
	for _,v:=range str{
		cMap[string(v)]++
	}
	var resStr=""
	for k,_:=range cMap{
		resStr+=k
	}
	fmt.Println("str",resStr)
	fmt.Println(cMap)

}
