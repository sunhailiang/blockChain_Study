package main

import "fmt"

/*func main()  {
	arr := [10]int {1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	s := arr[1:3:5]
	fmt.Println("s = ", s)
	fmt.Println("len（s） = ", len(s))
	fmt.Println("cap（s） = ", cap(s))

	s := arr[1:5:7]
	fmt.Println("s = ", s)
	fmt.Println("len（s） = ", len(s))		// 5-1 == 4
	fmt.Println("cap（s） = ", cap(s))		// 7-1

	s2 := s[0:6]
	fmt.Println("s = ", s2)
	fmt.Println("len（s） = ", len(s2))		// 6-0 == 6
	fmt.Println("cap（s） = ", cap(s2))
}*/

/*func main() {
	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	s := arr[2:5:5]					// {3， 4, 5}
	fmt.Println("s=", s)
	fmt.Println("len(s)=", len(s))
	fmt.Println("cap(s)=", cap(s))

	s2 := s[2:7] 					// {34567} {56789}
	fmt.Println("s=", s2)
	fmt.Println("len(s)=", len(s2))
	fmt.Println("cap(s)=", cap(s2))
}
*/

/*func main()  {
	// 1. 自动推导赋初值
	s1 := []int {1, 2, 4, 6}
	fmt.Println("s1 = ", s1)

	s2 := make([]int, 5, 10)
	fmt.Println("len=", len(s2), "cap=", cap(s2))

	s3 := make([]int, 7)
	fmt.Println("len=", len(s3), "cap=", cap(s3))

}*/

func main()  {
	s1 := []int {1, 2, 4, 6}  	// 创建一个有初始值的切片

	s1 = append(s1, 888)
	s1 = append(s1, 888)
	s1 = append(s1, 888)
	s1 = append(s1, 888)
	s1 = append(s1, 888)

	fmt.Println("s1=", s1)
}