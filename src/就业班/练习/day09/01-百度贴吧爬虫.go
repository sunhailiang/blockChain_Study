package main

import (
	"fmt"
	"strconv"
	"net/http"
	"io"
	"os"
)

func HttpGet(url string) (result string, err error) {
	resp, err1 := http.Get(url)
	if err1 != nil {
		err = err1				// 将封装函数内部的错误，传出给调用者。
		return
	}
	defer resp.Body.Close()

	// 循环读取 网页数据， 传出给调用者
	buf := make([]byte, 4096)
	for {
		n, err2 := resp.Body.Read(buf)
		if n == 0 {
			fmt.Println("读取网页完成")
			break
		}
		if err2 != nil && err2 != io.EOF {
			err = err2
			return
		}
		// 累加每一次循环读到的 buf 数据，存入result 一次性返回。
		result += string(buf[:n])
	}
	return
}

// 爬取页面操作。
func working(start, end int)  {
	fmt.Printf("正在爬取第%d页到%d页....\n", start, end)

	// 循环爬取每一页数据
	for i:=start; i<=end; i++ {
		url := "https://tieba.baidu.com/f?kw=%E7%BB%9D%E5%9C%B0%E6%B1%82%E7%94%9F&ie=utf-8&pn="+ strconv.Itoa((i-1)*50)
		result, err := HttpGet(url)
		if err != nil {
			fmt.Println("HttpGet err:", err)
			continue
		}
		// fmt.Println("result=", result)
		// 将读到的整网页数据，保存成一个文件
		f, err := os.Create("第 " + strconv.Itoa(i) + " 页" + ".html")
		if err != nil {
			fmt.Println("Create err:", err)
			continue
		}
		f.WriteString(result)
		f.Close()				// 保存好一个文件，关闭一个文件。
	}
}

func main()  {
	// 指定爬取起始、终止页
	var start, end int
	fmt.Print("请输入爬取的起始页（>=1）:")
	fmt.Scan(&start)
	fmt.Print("请输入爬取的终止页（>=start）:")
	fmt.Scan(&end)

	working(start, end)
}
