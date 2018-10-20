package main

import "fmt"

func main()  {
	G(3)
	fmt.Println(sum)
}
var sum int=1
func G(n int) {
	if n==1 {
		return
	}
    sum*=n
	G(n-1)

}