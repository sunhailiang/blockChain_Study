package main

import (
	"net/http"
	"os"
	"fmt"
	"io"
)

func main() {

	//配置路由

	//默认页面
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		dealRequest(w, r)
	})
	//请求图片资源
	http.HandleFunc("/img/", func(w http.ResponseWriter, r *http.Request) {
		dealRequest(w, r)
	})
	//请求文档资源
	http.HandleFunc("/docfile/", func(w http.ResponseWriter, r *http.Request) {
		dealRequest(w, r)
	})
	//请求视频资源
	http.HandleFunc("/vedio/", func(w http.ResponseWriter, r *http.Request) {
		dealRequest(w, r)
	})

	//监听http请求
	http.ListenAndServe("127.0.0.1:8001", nil)
}
func dealRequest(w http.ResponseWriter, r *http.Request) {
	cond := r.URL.String()[len(r.URL.String())-1:]
	fmt.Println("山玩意儿", cond)
	basePath := "E:/go/src/软文笔记/就业班技术点/tcp_bs模拟服务器基础版本"
	if cond == "/" {
		//打开文件夹
		dir, err := os.OpenFile(basePath+r.URL.String(), os.O_RDONLY, os.ModeDir)
		errFun("os.OpenFile", err)
		//读取文件夹
		files, err := dir.Readdir(-1)
		errFun("os.OpenFile", err)
		//是否找到文件
		//遍历文件夹内的文件
		for _, file := range files {
			if !file.IsDir() {
				if ("/" + file.Name()) == r.URL.String() {
					readFile(w, r, basePath)
				}
			}
		}
	} else if cond == "o" {

	} else {
		readFile(w, r, basePath)
	}

}
func readFile(w http.ResponseWriter, r *http.Request, basePath string) {
	u := r.URL
	buf := make([]byte, 4096)
	file, err := os.Open(basePath + u.String())
	errFun("os.Open", err)
	for {
		n, err := file.Read(buf)
		if err != nil {
			if err == io.EOF {
				fmt.Println("read EOF")
			} else {
				fmt.Println("read err:", err)
			}
			return
		}
		w.Write(buf[:n])
	}
	//读完关闭文件
	defer file.Close()
}

func errFun(str string, err error) {
	if err != nil {
		fmt.Println(str+" err:", err)
		os.Exit(1)
	}
}
