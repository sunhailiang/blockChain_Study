package main

import (
	"fmt"
	"strings"
)

func main()  {
	str:="操蛋的一天开始了"
   ad:= strings.Split(str,"")
   for index,v:=range ad{
   	   fmt.Printf("index：%d,value:%s \n",index,v)
   }
	fmt.Println(ad)
}
