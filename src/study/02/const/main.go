package main
import "fmt"
func  main()  {
	//常量
	const (
		//1，同一行值相同
		//2，从0开始下一行+1，给定初始值不会按照初始值+1
		num1=iota
		num2
		num3
		num4
	)
	fmt.Println("num1",num1)
	fmt.Println("num2",num2)
	fmt.Println("num3",num3)
	fmt.Println("num4",num4)
}
