package main

import(
  "fmt"	
)

func main(){
	defer func(){
		if err:=recover();err!=nil{
			fmt.Println("异常警报:",err)
		}
	}()

	b:=0;
	a:=100/b;
	fmt.Println(a)
}