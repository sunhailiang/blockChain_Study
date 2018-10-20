package main

import "fmt"

//  1-100  逢7 7的倍数 含有7的（17 27 72 73 74）敲桌子
//7的倍数
//个位带7
//十位带7
func main() {
	for i:=1;i<=100;i++{
		if i%7==0 || i%10==7 || i/10==7{
			fmt.Println("敲桌子")
		}else{
			fmt.Println(i)
		}

	}
	
}
