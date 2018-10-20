package main

import (
	"fmt"
)

func test(m int)  {
	var b int = 1000
	b += m
}

/*func main()  {
	var a int = 10
	var p *int = &a

	a = 100
	fmt.Println("a = ", a)

	test(10)

	*p = 250  // 借助a 变量的地址，操作a对应空间
	fmt.Println("a = ", a)
	fmt.Println("*p = ", *p)

	a = 1000
	fmt.Println("*p = ", *p)
}*/

/*func test()  {
	p := new(int)			// 默认类型的 默认值
	*p = 1000
}

func main()  {
	var a int = 10
	fmt.Println("&a", &a)

	var p *int

	// 在 heap 上申请一片内存地址空间

	fmt.Printf("%d\n", *p)
	fmt.Printf("%v\n", *p)		// 打印Go语言格式的字符串。

}*/

func swap(a, b int)  {
	a, b = b, a
	fmt.Println("swap  a:", a, "b:", b)
}

func swap2(x, y *int)  {
	*x, *y = *y, *x
}

func main()  {
	a, b := 10, 20

	swap2(&a, &b)			// 传地址值。
	fmt.Println("swap2: main  a:", a, "b:", b)

}




