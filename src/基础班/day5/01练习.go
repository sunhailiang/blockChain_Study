package main

import (
	"fmt"
	"math/rand"
	"time"
)

//通过键盘输入20个小写字母 统计个数
func main0101() {
	//sdjfjsdfuiyeudhajsh
	var arr [20]byte
	for i := 0; i < len(arr); i++ {
		fmt.Scanf("%c",&arr[i])
	}
	// ch[26]  ch[0]-ch[25]
	var ch [26]int  //用来统计字符个数

	//记录字母出现的次数
	for i := 0; i<len(arr);i++  {
		ch[arr[i]-'a']++
	}
	//打印字母出现次数
	for i := 0; i<len(ch);i++  {
		if ch[i]>0{
			fmt.Printf("字母：%c出现：%d次\n",'a'+i,ch[i])
		}
	}
	//for i := 0; i < len(arr); i++ {
	//	fmt.Printf("%c",arr[i])
	//}

}

//随机一注双色球彩票信息 红球6个 1-33 不能重复
func main(){
	//获取随机数种子
	rand.Seed(time.Now().UnixNano())

	var redball [6]int

	for i := 0; i<len(redball);i++  {
		//遍历之前存在的值和新随机数是否有重复
		//redball[i] = rand.Intn(33)+1//0-32
		temp:=rand.Intn(33)+1
		for j := 0; j < i; j++ {
			if temp == redball[j]{
				temp=rand.Intn(33)+1
				j=-1
				continue
			}
		}

		redball[i]=temp
	}

	fmt.Println(redball,"+",rand.Intn(16)+1)

}

