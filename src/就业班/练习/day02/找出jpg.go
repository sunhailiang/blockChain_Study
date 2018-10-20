package main

import (
	"os"
	"fmt"
	"strings"
)

func main() {
	findJPG()
}
func findJPG()  {
	 var resurcePath="E:/"
	 files,err:=os.OpenFile(resurcePath,os.O_RDONLY,os.ModeDir)
	 if err!=nil{
	 	fmt.Println("open err:",err)
	 	return
	 }
	 //操作完关闭
	 defer files.Close()
	 fs,err:=files.Readdir(-1)
	 if err!=nil{
		 fmt.Println("read err:",err)
		 return
	 }
	 for _,fileInfo:=range fs{
	 	if strings.HasSuffix(fileInfo.Name(),".png"){
	 		fmt.Println(fileInfo.Name()+"是一个png文件")
			copyTo(fileInfo,resurcePath)
		}
	 }
}

func copyTo(fileName os.FileInfo,path string)  {
	buf:=make([]byte,4096)
	fs,err:=os.OpenFile(path+fileName.Name(),os.O_RDWR,6)
	if err!=nil{
		fmt.Println("open err:",err)
		return
	}
	//创建文件
	dstDir:=path+"dst/"
	wfs,err:=os.Create(dstDir+fileName.Name())
	for {
		n,err:=fs.Read(buf)
		if err!=nil{
			fmt.Println("open err:",err)
			return
		}
		temp:=buf[:n]
		wfs.Write(temp)
	}
	wfs.Close()
	fs.Close()
}
