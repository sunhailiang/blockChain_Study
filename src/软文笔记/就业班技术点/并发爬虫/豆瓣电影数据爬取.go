package main

import (
	"net/http"
	"fmt"
	"strconv"
	"io"
	"os"
	"regexp"
)

func main() {
	var start, end int
	fmt.Printf("请输入爬取的起始页( >= 1 )：")
	fmt.Scan(&start)
	fmt.Printf("请输入爬取的终止页( >= 起始页)：")
	fmt.Scan(&end)
	worker(start, end)
}

func worker(start, end int) {
	page := make(chan int)
	for i := start; i < end; i++ {
		go SpiderPage(i, page)
	}
	for i := start; i < end; i++ {
		<-page
	}

}
func SpiderPage(index int, page chan<- int) {
	url := "https://movie.douban.com/top250?start=" + strconv.Itoa((index-1)*25) + "&filter="
	fmt.Printf("正在爬取第%d页信息\n", index)
	err, result := getHttp(url)
	if err != nil && err != io.EOF {
		fmt.Println("getHttp err:", err)
		return
	}
	saveFile(result, index)
	page <- index
}

//保存文件
func saveFile(result string, index int) {
	repName := regexp.MustCompile(`<img width="100" alt="(?s:(.*?))"`)
	filmName := repName.FindAllStringSubmatch(result, -1) //电影名称

	repScore := regexp.MustCompile(`<span class="rating_num" property="v:average">(?s:(.*?))</span>`)
	filmScore := repScore.FindAllStringSubmatch(result, -1)

	repNum := regexp.MustCompile(`<span>(?s:(\d*?))人评价</span>`)
	pNum := repNum.FindAllStringSubmatch(result, -1)

	file, err := os.Create(strconv.Itoa(index) + "页" + ".txt")
	if err != nil {
		fmt.Println("os.Create err:", err)
		return
	}
	file.WriteString("电影名称" + "\t\t\t\t" + "评分" + "\t\t\t\t" + "评分人数" + "\n")
	defer file.Close()
	for i := 0; i < len(filmName); i++ {
		file.WriteString(filmName[i][1] + "\t\t\t\t" + filmScore[i][1] + "\t\t\t\t" + pNum[i][1] + "\n")
	}
}

func getHttp(url string) (err error, result string) {

	res, _err := http.Get(url)
	defer res.Body.Close()

	if err != nil {
		err = _err
		return
	}
	buf := make([]byte, 4096)

	for {
		n, err := res.Body.Read(buf)
		if n == 0 {
			fmt.Println("读取完成", err)
			break
		}
		result += string(buf[:n])
	}
	return

}
