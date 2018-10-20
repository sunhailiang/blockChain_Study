package main

import (
	"os"
	"fmt"
	"bufio"
	"io"
	"strings"
)

func main() {
	//让用户输入一句话,判断这句话中有没有邪恶,如果有邪恶就替换成这种形式然后输出,如:老王很邪恶,输出后变成老王很**
	for {
		fmt.Println("您想说什么?")
		inputs := ""
		fmt.Scanf("%s", &inputs)
		matchWords(inputs)
	}
}

func matchWords(input string)  {
	//打开并读取文件
	fs, err := os.Open("E:/mgc.txt")
	if err != nil {
		fmt.Println("open failed")
		return
	}
	buf := bufio.NewReader(fs)
	//放读出的词汇
	slice := []string{}
	for {
		words, err := buf.ReadBytes('\n')
		if err == io.EOF||err != nil {
			break
		}
		slice = append(slice, string(words))
	}
	if len(slice) > 0 {
		for i := 0; i < len(slice); i++ {
			//fmt.Printf("%s,%d",slice[i],len(strings.TrimSpace(slice[i])))
			newSts:=strings.TrimSpace(slice[i])
			if strings.Contains(input, string(newSts)) {
				resStr:=strings.Replace(input, string(newSts), "***", -1)
				fmt.Printf("%s,..在说脏话，狗头铡伺候~\n",resStr)
			}
		}
	}
}
