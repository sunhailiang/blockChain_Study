package main

import (
	"fmt"
	"strconv"
	"net/http"
	"io"
	"regexp"
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

// 抓取一个网页，带有 10 个段子 —— 10 URL
func SpiderPage(idx int)  {
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

	for _, jokeURL := range alls {
		fmt.Println("jokeURL:", jokeURL[1])
	}
}

func toWork(start, end int)  {
	fmt.Printf("正在爬取 %d 到 %d 页...\n", start, end)

	for i:=start; i<=end; i++ {
		SpiderPage(i)
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
