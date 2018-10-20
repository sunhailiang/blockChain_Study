package main

import "fmt"

//定义结构体存储5名学生 三门成绩 求出每名学生的总成绩和平均成绩
type stud struct {
	id    int
	name  string
	score []int //结构体成员为数组|切片
}

func main() {
	arr:=[]stud{stud{101,"小明",[]int{100,99,88}},
		stud{102,"小红",[]int{88,56,83}},
		stud{103,"小刚",[]int{18,57,81}},
		stud{104,"小强",[]int{48,66,93}},
		stud{105,"小花",[]int{98,50,89}}}
	//五名学生
	for i:=0;i<len(arr);i++{
		//三门成绩
		sum:=0
		for j:=0;j<len(arr[i].score);j++{
			sum+=arr[i].score[j]
		}
		fmt.Printf("第%d名学生总成绩为：%d 平均成绩：%d\n",i+1,sum,sum/3)
	}
}
