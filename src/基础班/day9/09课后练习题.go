package main

import (
	"os"
	"fmt"
	"bufio"
	"io"
)

func main() {

	fp,err:=os.Open("D:/dict.txt")
	if err!=nil{
		fmt.Println("文件打开失败\n")
		return
	}
	defer fp.Close()


	dict:=make(map[string]string)

	r:=bufio.NewReader(fp)

	for {
		word,_:=r.ReadBytes('\n')
		trans,err:=r.ReadBytes('\n')
		dict[string(word[1:len(word)-2])]=string(trans[6:])
		if err==io.EOF{
			break
		}

	}
	//for k,v:=range dict{
	//	fmt.Printf("%s%s",k,v)
	//}
	fmt.Println(dict["a"])

}
