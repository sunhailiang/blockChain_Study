package main

import (
	"os"
	"fmt"
	"strings"
)

func main() {
	fmt.Println("请输入查询地址")
	var path string
	fmt.Scanf("%s",&path)
	findJpg(path)
}
func findJpg(path string)  {
	//打开目录
	dir,err:=os.OpenFile(path,os.O_RDONLY,os.ModeDir)
	if err!=nil {
		fmt.Println("open dir err:",err)
		return
	}
	//读取文件夹-1表示查找所有
	fileInfos,err:= dir.Readdir(-1)
	//遍历文件夹下所有的资料
	for _,fileinfo:=range fileInfos {
		if !fileinfo.IsDir(){
			if strings.HasSuffix(fileinfo.Name(),".jpg"){
				fmt.Println(fileinfo.Name()+",是jpg文件")
			}
		}
	}
}
