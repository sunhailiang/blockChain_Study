package main
import (
	"strings"
	// "strconv"
	
	"fmt"
)
func main(){
	// var name string="harry"
    //1判断字符串开头
	// res:= strings.HasPrefix(name,"d");
	// fmt.Println("字符串开头是否包含这个字符？",res)
	
	
	//2判断字符串结尾
	// res:=strings.HasSuffix(name,"y")
	// fmt.Println("字符串结尾是否包含这个字符？",res)

	//3判断字符在字符串中首次出现的位置
	// index:=strings.Index(name,"r")
	// fmt.Println("在字符串中首次出现的位置",index)

	//4最后出现的位置
	// lastindex:=strings.LastIndex(name,"r")
	// fmt.Println("在字符串中首次出现的位置",lastindex)

	// var str string=" th,is is str,i ng ";
	// //去掉字符串首尾空白
	// res:=strings.TrimSpace(str)
	// fmt.Println("去掉首位空格",res)

	// //去掉字符串首尾字符
	// res:=strings.Trim(str,"g ")
	// fmt.Println("去掉首尾空格",res)
	
	// //切掉开头字符
	// res1:=strings.TrimLeft(str," t")
	// fmt.Println("切掉开头空格",res1)

	// //切掉结尾字符
	// res2:=strings.TrimRight(str,"g ")
	// fmt.Println("切掉结尾空格",res2)

	// //返回空格分隔的切片数组
	// res3:=strings.Fields(str)
	// fmt.Println("返回空格分隔的切片",len(res3))

	// //返回条件分割切片数组
	// res4:=strings.Split(str,",")
	// fmt.Println("返回条件分隔的切片",len(res4))


	//join
	//  arr:=[]string{"a","b","c"}
	//  res:=strings.Join(arr,"d")
	//  fmt.Println("返回join后的新数组？",res)

	//itoa整数转成字符串
	// var num int=3000;
	// fmt.Println(strconv.Itoa(num))

	// //Atoi是将字符串转成整数
	// var num int=3000;
	// fmt.Println(strconv.Atoi(num))

	//字符串替换
	// var str string="this is a string";
//    fmt.Println("字符串替换",strings.Replace(str,"s","b",2)) //参数1：目标值符传，2替换的内容，3替换结果，4替换多少个？

	// var str string="this is a string";
	//  fmt.Println("s的个数",strings.Count(str,"s"));
	 

	 var str string="this is a string";
	 fmt.Println("转大写",strings.ToUpper(str));
	 

	 var strUP string="THIS IS A STRING";
     fmt.Println("转小写",strings.ToLower(strUP));
}