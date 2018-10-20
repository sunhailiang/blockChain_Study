package main

import (
	"fmt"
	"strconv"
	"net/http"
	"io"
	"regexp"
	"os"
)

func main() {
	var start, end int
	fmt.Printf("请输入爬取的起始页( >= 1 )：")
	fmt.Scan(&start)
	fmt.Printf("请输入爬取的终止页( >= 起始页)：")
	fmt.Scan(&end)
	Worker(start, end)
}

func Worker(start, end int) {
	page := make(chan int)
	for i := start; i <= end; i++ {
		go Spider(i, page)
	}
	for i := start; i <= end; i++ {
		fmt.Printf("第%d页数据爬取完成", <-page)
	}
}
func Spider(index int, page chan int) {
	url := "https://www.douyu.com/gapi/rkc/directory/2_201/" + strconv.Itoa(index)
	err, res := HttpGet(url)
	//fmt.Println("啥玩意儿",res)
	if err != nil {
		fmt.Println("HttpGet err:", err)
		return
	}
	//正则匹配内容
	rxp := regexp.MustCompile(`"rs1":"(?s:(.*?))"`)
	urls := rxp.FindAllStringSubmatch(res, -1)
	for i := 0; i < len(urls); i++ {
		go func(i int, urls []string) {
			_, result := HttpGet(urls[1])
			saveImg(i+1, result)
		}(i, urls[i])
	}
	<-page
}

func HttpGet(url string) (resErr error, result string) {
	res, err := http.Get(url)
	defer res.Body.Close()
	if err != nil {
		fmt.Println("httpGet err", err)
		return
	}
	buf := make([]byte, 4096)
	for {
		n, err := res.Body.Read(buf)
		if n == 0 {
			fmt.Println("读完,Tip:", err)
			break
		}
		if err != nil && err != io.EOF {
			resErr = err
			return
		}
		result += string(buf[:n])
	}
	return
}

func saveImg(index int, result string) {
	file, err := os.Create(strconv.Itoa(index) + ".jpg")
	if err != nil {
		fmt.Println("os.Create", err)
		return
	}
	_, err = file.Write([]byte(result))
}
