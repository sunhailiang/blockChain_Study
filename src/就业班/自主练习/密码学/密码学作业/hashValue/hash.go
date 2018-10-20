package main

import (
	"os"
	"fmt"
	"io"
	"crypto/sha256"
)

func main() {
	GetHashVal()
}
func GetHashVal() {

	filePath := "E:/go/src/test.mp4"
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("openFaild", err)
		return
	}
	buf := make([]byte, 1024*1024)
	MySha:=sha256.New()
	for {
		_, err := file.Read(buf)
		if err == io.EOF {
			break
		} else {
			MySha.Write(buf)
		}
	}
	res:=MySha.Sum(nil)
	fmt.Println("最终读到什么？", res)
}

