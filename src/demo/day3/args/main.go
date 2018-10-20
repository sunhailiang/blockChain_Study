package main

import(
	"fmt"
)

func main(){
	fmt.Println("和:",getvalue(0,1,2,3,4,5,6,7,8,9))
	fmt.Println("字符串结果：",concat("我是","你","大","爷"))
}

//多个数想加的值
func getvalue(arg...int)(res int){
	  
	for i:=0;i<len(arg);i++{
         res+=arg[i]
	  }
	  return;
}
//多个字符串拼接
func concat(arg...string)(str string){
	for i:=0;i<len(arg);i++{
		str+=arg[i];
	}
	return;
}