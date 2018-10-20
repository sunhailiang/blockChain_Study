package main

import (
	"fmt"
	"net/http"
	"io"
)

func main()  {
	// 获取服务器 应答包内容
	//resp, err := http.Get("http://www.itcast.cn/")
	resp, err := http.Get("http://www.baidu.com/")
	//resp, err := http.Get("http://127.0.0.1:8000/hello.c")
	//resp, err := http.Get("http://127.0.0.1:8000/lf.txt")
	if err != nil {
		fmt.Println("http.Get err:", err)
		return
	}
	defer resp.Body.Close()

	// 简单查看应答包
	fmt.Println("Header:", resp.Header)
	fmt.Println("Status:", resp.Status)
	fmt.Println("Header:", resp.StatusCode)
	fmt.Println("Proto:", resp.Proto)

	buf := make([]byte, 4096)

	var result string
	for {
		n, err := resp.Body.Read(buf)
		if n == 0 {
			fmt.Println("--Read finish!")
			break
		}
		if err != nil && err != io.EOF {
			fmt.Println("resp.Body.Read err:", err)
			return
		}
		result += string(buf[:n])
	}
	fmt.Printf("result:|%v|\n", result)
}
