package main

import (
	"fmt"
	"strings"
)

func main() {
    //练习1：从日期字符串（"2008-08-08"）中分析出年、月、日；2008年08月08日。
	//让用户输入一个日期格式如:2008-01-02,你输出你输入的日期为2008年1月2日
      date:=""
      fmt.Println("请输入日期如：2018-10-29")
      fmt.Scanf("%s",&date)
      slice:= strings.Split(date,"-")
      newDate:=slice[0]+"年"+slice[1]+"月"+slice[2]+"日"
      fmt.Printf("您的生日是:%s",newDate)

}
