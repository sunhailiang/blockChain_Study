package main

import (
	"fmt"
	"os"
	"strings"
	"bufio"
	"io"
)

func main() {
	var sourceSrc string
	var word string
	fmt.Println("输入查询目录")
	fmt.Scanf("%s", &sourceSrc)
	fmt.Println("输入检索词汇")
	fmt.Scanf("%s", &word)
	wordNum := getFileListToFindWords(sourceSrc, word)
	fmt.Printf("单词%s,在当前目录出现%d次", word, wordNum)
}
func getFileListToFindWords(sourceSrc string, word string) int {
	wordNum := 0
	//打开目录
	dir, err := os.OpenFile(sourceSrc, os.O_RDONLY, os.ModeDir)
	if err != nil {
		fmt.Println("dir open err:", err)
		return -1
	}
	//读取目录内容-1读取全部
	fileInfos, err := dir.Readdir(-1)
	//遍历目录中内容
	for _, file := range fileInfos {
		if !file.IsDir() {
			if strings.HasSuffix(file.Name(), ".txt") || strings.HasSuffix(file.Name(), ".xml") {
				wordNum += findWordFromFile(word, sourceSrc+"/"+file.Name())
			}
		} else {
			wordNum += getFileListToFindWords(sourceSrc+"/"+file.Name(), word)
		}
	}
	return wordNum
}
func findWordFromFile(word string, filePath string) int {
	//打开文件
	file, err := os.OpenFile(filePath, os.O_RDWR, 6)
	//创建字典累加value
	wordMap := make(map[string]int)
	if err != nil {
		fmt.Println("open file err", err)
		return -1
	}
	//最终挂不必次
	defer file.Close()
	//创建缓冲区
	buf := make([]byte, 4096)
	//创建阅读器
	reader := bufio.NewReader(file)
	n, err := reader.Read(buf)
	if err != nil && err == io.EOF {
		fmt.Println("read completed")
		return wordMap[word]
	} else if err != nil {
		fmt.Println("read file err:", err)
		return -1
	}
	temp := buf[0:n]
	for _, v := range strings.Fields(string(temp)) {
		if v == word {
			wordMap[v]++
		}
	}
	return wordMap[word]
}
