package main

import (
	"fmt"
	"os"
	"strings"
	"io"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	var sourceSrc string
	var fileType string
	var targetPath string
	fmt.Println("输入原目录")
	fmt.Scanf("%s", &sourceSrc)
	fmt.Println("输入文件类型")
	fmt.Scanf("%s", &fileType)
	fmt.Println("输入目标目录")
	fmt.Scanf("%s", &targetPath)

	getFileList(sourceSrc, fileType, targetPath)

}

func getFileList(sourceSrc string, fileType string, targetPath string) {
	//打开指定路径文件夹
	dir, err := os.OpenFile(sourceSrc, os.O_RDONLY, os.ModeDir)
	if err != nil {
		fmt.Println("open dir err:", err)
		return
	}
	//读取文件夹-1读取全部
	fileInfos, err := dir.Readdir(-1)

	if err != nil {
		fmt.Println("read dir err:", err)
		return
	}

	for _, fileInfo := range fileInfos {
		if !fileInfo.IsDir() {
			if strings.HasSuffix(fileInfo.Name(), fileType) {
				copyToDir(sourceSrc, fileInfo, targetPath)
			}
		}
	}
}

func copyToDir(sourceSrc string, file os.FileInfo, targetDir string) {
	sourcefile, err := os.OpenFile(sourceSrc+"/"+file.Name(), os.O_RDWR, 6)
	if err != nil {
		fmt.Println("file open err:", err)
		return
	}
	//创建缓冲区
	buf := make([]byte, 4096)
	//创建文件
	targetFile, err := os.Create(targetDir + "/" + file.Name())
	for {
		n, err := sourcefile.Read(buf)
		if err != nil && err == io.EOF {
			fmt.Println("read completed...")
			return
		} else if err != nil {
			fmt.Println("read err:", err)
			return
		}
		temp := buf[:n]
		targetFile.Write(temp)
	}
}
