package main

import "fmt"

//有个切片，找出第二大的数，并且打印出来
func main()  {
	slice:=[]int {5,100,32,45,21,67,32,68,41,99,13,71,2222,444}
	res:=getSecNum(slice)
	fmt.Println("res",res)
}

func getSecNum(slice []int)[]int{
	//用冒泡，从小到大，倒数第二个就是老二
	for i:=0;i< len(slice);i++{
		for j:=0;j< len(slice)-1-i;j++{
			if slice[j]>slice[j+1]{
				slice[j],slice[j+1]=slice[j+1],slice[j]
			}
		}
	}
	return slice[len(slice)-2:len(slice)-1]
}


