package main

import (
	"net/http"
	"fmt"
	"io"
	"strconv"
	"regexp"
	"strings"
	"os"
)

func main() {
	// 指定爬取起始、终止页
	var start, end int
	fmt.Print("请输入爬取的起始页（>=1）:")
	fmt.Scan(&start)
	fmt.Print("请输入爬取的终止页（>=start）:")
	fmt.Scan(&end)
	//开始爬取
	working(start, end)
}
func working(start, end int) {
	var page = make(chan int)
	fmt.Printf("正在爬取%d到%d页...", start, end)
	for i := start; i < end; i++ {
		go spiderPage(i, page)
	}
	for i := start; i < end; i++ {
		fmt.Printf("第%d页爬取完成\n", <-page)
	}

}
func spiderPage(index int, page chan<- int) {
	url := "https://www.pengfu.com/xiaohua_" + strconv.Itoa(index) + ".html"
	res, err := httpGet(url)
	if err != nil {
		fmt.Println("httpGet err:", err)
		return
	}
	rxpUrl := regexp.MustCompile(`<h1 class="dp-b"><a href="(?s:(.*?))"`)
	jokePageUrl := rxpUrl.FindAllStringSubmatch(res, -1)
	kv := make(map[string]string)
	for i := 0; i < len(jokePageUrl); i++ {
		getJokeContentUrl := jokePageUrl[i][1]
		title, content := spiderJokePage(getJokeContentUrl)
		kv[title] = content
	}
	saveJoks(index, kv)
	//防止主程结束
	page <- index
}

func httpGet(url string) (result string, err error) {
	resp, err_ := http.Get(url)
	defer resp.Body.Close()
	if err != nil {
		err = err_
		return
	}
	buf := make([]byte, 4096)
	for {
		n, _err := resp.Body.Read(buf)
		if n == 0 {
			break
		}
		if _err != nil && _err != io.EOF {
			err = _err
			return
		}
		result += string(buf[:n])
	}
	return
}

func spiderJokePage(url string) (title, content string) {
	result, err := httpGet(url)
	if err != nil {
		fmt.Println("JokePage getHttp err", err)
		return
	}
	//提取页面内容
	rxpTit := regexp.MustCompile(`<h1>(?s:(.*?))</h1>`)
	rxpContent := regexp.MustCompile(`<div class="content-txt pt10">(?s:(.*?))<a id="prev" href="`)
	title = strings.Replace(rxpTit.FindAllStringSubmatch(result, -1)[0][1], "\t", "", -1)
	content = strings.Replace(rxpContent.FindAllStringSubmatch(result, -1)[0][1], "\t", "", -1)
	content = strings.Replace(content, "\n", "", -1)
	content = strings.Replace(content, "<br><br />", "", -1)
	return

}
func saveJoks(index int, kv map[string]string) {
	file, err := os.Create("第" + strconv.Itoa(index) + "页.txt")
	if err != nil {
		fmt.Println("os.Create err", err)
		return
	}
	defer file.Close()
	file.WriteString("title" + "\t\t\t\t\t" + "内容\n")
	for k, v := range kv {
		file.WriteString(k + "\t\t\t\t\t" + v + "\n")
	}
}
