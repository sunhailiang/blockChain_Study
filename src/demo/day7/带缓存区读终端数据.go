package main

import (
	"bufio"
	"os"
	"fmt"
)

func main()  {
	//创建读数据对象 os.stdin用来读取终端信息
	 reader :=bufio.NewReader(os.Stdin)
	 str,err:=reader.ReadString('\n')
	if err!=nil {
		fmt.Println("read string failed,err",err)
		return
	}
	fmt.Printf("read str succ:%s \n",str)
}
