package main

import "fmt"

func Swap(p *[9]int) {
	for i, v := range *p {
		(*p)[0] = 89
		fmt.Printf("%d,%d",i,v)
	}
}
func main() {
	a:= [9]int{1, 2, 3, 4, 5, 4, 4, 4, 4}
	Swap(&a)
}
