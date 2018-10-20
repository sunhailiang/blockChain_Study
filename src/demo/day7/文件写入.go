package main

import (
	"os"
	"fmt"
	"bufio"
)

func main(){
	outputFile,err:=os.OpenFile("d:/readorwrite/fuck.txt",os.O_WRONLY|os.O_CREATE,0666)
	if err!=nil{
		fmt.Printf("An error occurred with cread ion \n");
		return
	}
	//操作完成关闭文件
	defer outputFile.Close()
	outputWriter:=bufio.NewWriter(outputFile)
	outputString:="fuck you \n"
	for i:=0;i<100;i++ {
		outputWriter.WriteString(outputString)
	}
	outputWriter.Flush()

}
