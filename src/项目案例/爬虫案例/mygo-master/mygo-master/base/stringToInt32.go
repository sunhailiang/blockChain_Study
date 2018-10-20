package mygo

import (
	"fmt"
	"strconv"
	"strings"
)

func SwxctxInt32Param(s string) int32 {
	//give up space again(inspect)
	var r string = strings.TrimSpace(s)
	var resu int64
	var err error
	var exists bool = strings.Contains(r, ".")
	if exists {
		//find . for str
		sub_num := strings.Count(r, ".")
		if sub_num > 1 {
			fmt.Println("The numbers are incorrect...")

		} else {
			//. index
			index := strings.Index(r, ".")
			//string to int32(index[0]-.[index])
			front_num, err := strconv.ParseInt(r[0:index], 10, 32)
			if err != nil {
				panic(fmt.Errorf("parse int32 param error, %s", err))
			}
			//.[index]+1->.[index]+2 string to int32
			after_num, err := strconv.ParseInt(r[index+1:index+2], 10, 32)
			if err != nil {
				panic(fmt.Errorf("parse int32 param error, %s", err))
			}
			//front_num+1
			if after_num >= 5 {
				resu = front_num + 1
			} else {
				resu = front_num
			}
		}
	} else {
		resu, err = strconv.ParseInt(r, 10, 32)
		if err != nil {
			panic(fmt.Errorf("parse int32 param error, %s", err))
		}
	}
	return int32(resu)
}

func main() {
	var a string = "623.4"
	// var b int32
	// b = xAdsDigitalFormat(a)
	// fmt.Print("%d",b)
	res := xAdsDigitalFormat(a)
	fmt.Print(res)
}
