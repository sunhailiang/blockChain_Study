package main

import (
	"sort"
	"fmt"
)

func main() {
	arr := []int{1, 4, 9, 2, 3, 7, 4, 2, 1, 6, 7, 84, 11, 23, 34, 45, 66}
	//升序-----int
	sort.Ints(arr)
	fmt.Println(arr)
	//降序
	sort.Sort(sort.Reverse(sort.IntSlice(arr)))
	fmt.Println(arr)
	arrstring:=[]string{"a","l","a","s","r","j","g"}
	//升序----string
	sort.Sort(sort.StringSlice(arrstring))
	fmt.Println(arrstring)
	//降序
	sort.Sort(sort.Reverse(sort.StringSlice(arrstring)))
	fmt.Println(arrstring)

    //升序---float
	arrfloat := []float64{1.4, 4.5,2.9, 1.2, 3.7, 4.7, 4.2, 6.2, 1.1, 0.6, 71.2, 18.4, 1.11, 1.23, 34, 45, 66}
	sort.Float64s(arrfloat)
	fmt.Println(arrfloat)
	//降序
	sort.Sort(sort.Reverse(sort.Float64Slice(arrfloat)))
	fmt.Println(arrfloat)
}
