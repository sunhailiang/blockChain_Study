package main

import (
	"os"
	"fmt"
	"bufio"
	"io"
	"strings"
)

func main() {
	//翻译词典
	checkVal()
}

//查找value
func checkVal() {
	for {
		fmt.Println("请输入您要查询的词汇;")
		words := ""
		fmt.Scanf("%s", &words)
		dic := makeDic()
		flag := false
		for k, v := range dic {
			if string(strings.TrimSpace(k)) == words {
				fmt.Println(v)
				flag = true
				break
			}
		}
		if !flag {
			fmt.Println("查无此货")
		}
		if words == "q" {
			break
		}
	}
}

//制作字典
func makeDic() map[string]string {
	fs, err := os.Open("D:/blackhourse/上课全部资料/dictest.txt")
	if err != nil {
		fmt.Println("open fiald err:", err)
		return nil
	}
	//创建缓存区
	r := bufio.NewReader(fs)
	var dict = make(map[string]string)
	mapKey := "null"
	for {
		str, err := r.ReadBytes('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("read failed err:", err)
			break
		}
		//判断行首是否是#符号，有则添加做key
		if strings.Contains(string(str), "#") {
			newKey := strings.TrimLeft(string(str), "#")
			mapKey = newKey
		} else if strings.Index(string(str), ":") > -1 { //查找行是否包含:符号,有则用来截取翻译
			if mapKey != "null" {
				vStr := strings.Split(string(str), ":")[1]
				dict[mapKey] = vStr
				mapKey = "null"
			} else {
				println("key有误")
				break
			}
		}
	}
	//关闭
	defer fs.Close()
	//返回字典
	return dict
}
