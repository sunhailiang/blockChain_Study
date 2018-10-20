package main

import "fmt"

func main() {
	str := "helleweordere.e"
	for i,v:=range str{
		if v=='e'{
			fmt.Println(i)
		}
	}
}
