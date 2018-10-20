package main

import "fmt"

func main() {
	var n float64
	fmt.Scanf("%v", &n)

	resl := getl(&n)
	ress := gets(&n)

	fmt.Println(resl)
	fmt.Println(ress)

}

func getl(r *float64) (papa float64) {
	var pi float64 = 3.14
	papa = 2 * pi * *r
	return
}
func gets(r *float64) float64 {
	var pi float64 = 3.14
	return pi * *r * *r
}
