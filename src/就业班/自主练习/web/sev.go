package main

import (
	"net/http"
	"os"
	"fmt"
	"io"
)

func main() {

	http.HandleFunc("/", testJson)
	http.ListenAndServe("192.168.35.79:8008", nil)
}

func testJson(w http.ResponseWriter, r *http.Request) {

	u := r.URL
	dir, err := os.OpenFile("E:/go/src/就业班/自主练习/web/", os.O_RDONLY, os.ModeDir)
	if err != nil {
		fmt.Println("os.OpenFile", err)
		return
	}
	fileInfos, err := dir.Readdir(-1)
	if err != nil {
		fmt.Println("dir.Readdir", err)
		return
	}
	for _, fileInfo := range fileInfos {
		if !fileInfo.IsDir() {
			if ("/" + fileInfo.Name()) == u.String() {
				file, err := os.Open("E:/go/src/就业班/自主练习/web/" + u.String())
				if err != nil {
					fmt.Println("os.Open", err)
					return
				}
				buf := make([]byte, 4096)
				for {
					n, err := file.Read(buf)
					if err != nil {
						if err == io.EOF {
							fmt.Println("read conplete")
						} else {
							fmt.Println("read err:", err)
						}
						return
					}
					w.Write(buf[:n])
				}
			}
		}
	}
}
