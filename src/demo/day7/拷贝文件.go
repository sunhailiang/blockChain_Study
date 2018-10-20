package main

import (
	"os"
	"io"
)

func main()  {
	CopyFile("d:/readorwrite/lala.txt","d:/readorwrite/read.txt")
}

func CopyFile(targetFileName,sourceFileName string)(written int64,err error)  {
	//打开数据源文件
	src,err:=os.Open(sourceFileName)
	//如果报错，提示错误
	if err!=nil {
       return
	}
	//关闭数据源文件
	defer  src.Close()
	//写入文件
	targetfile,err:=os.OpenFile(targetFileName,os.O_WRONLY|os.O_CREATE,0644)
	if err!=nil {
		return
	}
	defer targetfile.Close()

	return io.Copy(targetfile,src)


}
