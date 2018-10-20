package main

import "fmt"

func main(){

	fmt.Println("请输入比较数字")
	var a int
	var b int
	var c int
	fmt.Scanf("%d %d %d",&a,&b,&c)

	if a>b{
		if a>c{
			fmt.Println("a是个胖子")
		}else{
			fmt.Println("c是个胖子")
		}
	}else{
		if b>c{
			fmt.Println("b是个胖子")
		}else{
			fmt.Println("c是个胖子")
		}
	}

}
