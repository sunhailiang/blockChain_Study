package main

import "fmt"

func main(){
	var sum int
    for i:=1;i<101;i++{
    	if i>1&&i%2==0{
    		sum-=i
		}else{
			sum+=i
		}
	}
	fmt.Println(sum)
}
