package main

import "fmt"

func main()  {
	for i:=0;i<1000;i++  {
		if i%2==1&&i%3==2&&i%5==4&&i%6==5&&i%7==0{
			fmt.Println(i)
		}
	}

}
