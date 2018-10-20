package main

import "fmt"

type stu struct {
	name string
	age int
	score int
}

func main(){
	m:=make(map[int][]stu,6)
	m[101][0]=stu{"曹操",50,88}
	m[101][1]=stu{"张辽",38,98}

	for k,v:=range m {
		for i,data:=range v{
			fmt.Println("key:",k,"index",i,"value:",data)
		}
	}

}
