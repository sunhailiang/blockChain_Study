package main

import (
	"fmt"
	"strconv"
	"net/http"
	"io"
	"regexp"
	"strings"
	"os"
)

// 获取一个网页所有的内容， result 返回
func HttpGet(url string) (result string, err error) {
	resp, err1 := http.Get(url)
	if err1 != nil {
		err = err1
		return
	}
	defer resp.Body.Close()
	buf := make([]byte, 4096)
	for {
		n, err2 := resp.Body.Read(buf)
		if n == 0 {
			break
		}
		if err2 != nil && err2 != io.EOF {
			err = err2
			return
		}
		result += string(buf[:n])
	}
	return
}

func SaveJoke2File(idx int, fileTitle, fileContent []string)  {
	path :=  strconv.Itoa(idx) + "页.txt"
	f, err := os.Create(path)
	if err != nil {
		fmt.Println("Create err:", err)
		return
	}
	defer f.Close()
	n := len(fileTitle)
	for i:=0; i<n; i++ {
		f.WriteString(fileTitle[i] + "\n" + fileContent[i] + "\n")
		f.WriteString("---------------------------------------------------\n")
	}
}

// 抓取一个网页，带有 10 个段子 —— 10 URL
func SpiderPage(idx int, page chan int)  {
	// 拼接URL
	url := "https://www.pengfu.com/xiaohua_" + strconv.Itoa(idx) + ".html"

	// 封装函数获取段子的URL
	result, err := HttpGet(url)
	if err != nil {
		fmt.Println("HttpGet err:", err)
		return
	}
	// 解析、编译正则
	ret := regexp.MustCompile(`<h1 class="dp-b"><a href="(?s:(.*?))"`)

	// 提取需要信息 —— 每个段子的 URL
	alls := ret.FindAllStringSubmatch(result, -1)

	// 创建用户存储 title 、content 的 切片, 初始容量为0
	fileTitle := make([]string, 0)
	fileContent := make([]string, 0)

	for _, jokeURL := range alls {
		//fmt.Println("jokeURL:", jokeURL[1])
		title, conntent, err := SpiderJokePage(jokeURL[1])			// 爬取一个笑话页面的 title 和 content
		if err != nil {
			fmt.Println("SpiderJokePage err:", err)
			continue
		}
/*		fmt.Println("title:", title)
		fmt.Println("conntent:", conntent)*/
		fileTitle = append(fileTitle, title)		// 将title 追加到 切片末尾
		fileContent = append(fileContent, conntent)	// Content 追加到 切片末尾
	}
/*	fmt.Println("title:", fileTitle)
	fmt.Println("conntent:", fileContent)*/
	SaveJoke2File(idx, fileTitle, fileContent)

	// 防止主go程提前结束
	page <- idx
}

// 爬取一个笑话页面的 title 和 content
func SpiderJokePage(url string) (title, content string, err error) {
	result, err1 := HttpGet(url)
	if err1 != nil {
		err = err1
		return
	}
	// 编译解析正则表达式 —— title
	ret1 := regexp.MustCompile(`<h1>(?s:(.*?))</h1>`)
	// 提取 title
	alls := ret1.FindAllStringSubmatch(result, 1)		// 有2 处，取第一个
	for _, tmpTitle := range alls {
		title = tmpTitle[1]

		title = strings.Replace(title, "\t", "", -1)
		break
	}
	// 编译解析正则表达式 —— content
	ret2 := regexp.MustCompile(`<div class="content-txt pt10">(?s:(.*?))<a id="prev" href="`)
	// 提取 title
	alls2 := ret2.FindAllStringSubmatch(result, 1)		// 有2 处，取第一个
	for _, tmpContent := range alls2 {
		content = tmpContent[1]
		content = strings.Replace(content, "\n", "", -1)
		content = strings.Replace(content, "\t", "", -1)
		content = strings.Replace(content, "&nbsp; ", "", -1)
		break
	}
	return
}

func toWork(start, end int)  {
	fmt.Printf("正在爬取 %d 到 %d 页...\n", start, end)

	page := make(chan int)

	for i:=start; i<=end; i++ {
		go SpiderPage(i, page)
	}

	for i:=start; i<=end; i++ {
		fmt.Printf("第 %d 页爬取完毕\n", <-page)
	}
}

func main()  {
	// 指定爬取起始、终止页
	var start, end int
	fmt.Print("请输入爬取的起始页（>=1）:")
	fmt.Scan(&start)
	fmt.Print("请输入爬取的终止页（>=start）:")
	fmt.Scan(&end)

	toWork(start, end)
}
