package douban

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"time"
)

type Swoop struct {
	url    string
	header map[string]string
}

// get_html_header
func (swoop Swoop) get_html_header() string {
	client := &http.Client{}
	req, err := http.NewRequest("GET", swoop.url, nil)
	if err != nil {
		log.Fatalf("new request err->%v", err)
	}
	for key, value := range swoop.header {
		req.Header.Add(key, value)
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("do client err->%v", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("read resp err->%v", err)
	}
	return string(body)

}

func Convert() {
	header := map[string]string{
		"Host":                      "movie.douban.com",
		"Connection":                "keep-alive",
		"Cache-Control":             "max-age=0",
		"Upgrade-Insecure-Requests": "1",
		"Accept":                    "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8",
		"Referer":                   "https://movie.douban.com/top250",
	}
	// 随机获取user_agent，避免被封
	header["User-Agent"] = GetRandomUserAgent()

	//创建csv文件
	f, err := os.Create("./douban.csv")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	//写入
	f.WriteString("电影名称" + "\t" + "评分" + "\t" + "评价数量" + "\t" + "\r\n")

	// 读取每页数据并写入
	for i := 0; i < 10; i++ {
		log.Println("正在爬取第" + strconv.Itoa(i) + "页数据...")
		url := "https://movie.douban.com/top250?start=" + strconv.Itoa(i*25)
		swoop := &Swoop{url, header}
		html := swoop.get_html_header()

		//评价人数
		commentCount := `<span>(.*?)评价</span>`
		rp2 := regexp.MustCompile(commentCount)
		txt2 := rp2.FindAllStringSubmatch(html, -1)

		//评分
		pattern3 := `property="v:average">(.*?)</span>`
		rp3 := regexp.MustCompile(pattern3)
		txt3 := rp3.FindAllStringSubmatch(html, -1)

		//电影名称
		pattern4 := `img width="(.*?)" alt="(.*?)" src=`
		rp4 := regexp.MustCompile(pattern4)
		txt4 := rp4.FindAllStringSubmatch(html, -1)

		f.WriteString("\xEF\xBB\xBF")

		for i := 0; i < len(txt2); i++ {
			fmt.Printf("%s %s %s\n", txt4[i][1], txt3[i][1], txt2[i][1])
			f.WriteString(txt4[i][2] + "\t" + txt3[i][1] + "\t" + txt2[i][1] + "\t" + "\r\n")
		}
	}
}

var userAgent = [...]string{"Mozilla/5.0 (compatible, MSIE 10.0, Windows NT, DigExt)",
	"Mozilla/4.0 (compatible, MSIE 7.0, Windows NT 5.1, 360SE)",
	"Mozilla/4.0 (compatible, MSIE 8.0, Windows NT 6.0, Trident/4.0)",
	"Mozilla/5.0 (compatible, MSIE 9.0, Windows NT 6.1, Trident/5.0,",
	"Opera/9.80 (Windows NT 6.1, U, en) Presto/2.8.131 Version/11.11",
	"Mozilla/4.0 (compatible, MSIE 7.0, Windows NT 5.1, TencentTraveler 4.0)",
	"Mozilla/5.0 (Windows, U, Windows NT 6.1, en-us) AppleWebKit/534.50 (KHTML, like Gecko) Version/5.1 Safari/534.50",
	"Mozilla/5.0 (Macintosh, Intel Mac OS X 10_7_0) AppleWebKit/535.11 (KHTML, like Gecko) Chrome/17.0.963.56 Safari/535.11",
	"Mozilla/5.0 (Macintosh, U, Intel Mac OS X 10_6_8, en-us) AppleWebKit/534.50 (KHTML, like Gecko) Version/5.1 Safari/534.50",
	"Mozilla/5.0 (Linux, U, Android 3.0, en-us, Xoom Build/HRI39) AppleWebKit/534.13 (KHTML, like Gecko) Version/4.0 Safari/534.13",
	"Mozilla/5.0 (iPad, U, CPU OS 4_3_3 like Mac OS X, en-us) AppleWebKit/533.17.9 (KHTML, like Gecko) Version/5.0.2 Mobile/8J2 Safari/6533.18.5",
	"Mozilla/4.0 (compatible, MSIE 7.0, Windows NT 5.1, Trident/4.0, SE 2.X MetaSr 1.0, SE 2.X MetaSr 1.0, .NET CLR 2.0.50727, SE 2.X MetaSr 1.0)",
	"Mozilla/5.0 (iPhone, U, CPU iPhone OS 4_3_3 like Mac OS X, en-us) AppleWebKit/533.17.9 (KHTML, like Gecko) Version/5.0.2 Mobile/8J2 Safari/6533.18.5",
	"MQQBrowser/26 Mozilla/5.0 (Linux, U, Android 2.3.7, zh-cn, MB200 Build/GRJ22, CyanogenMod-7) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1"}

var r = rand.New(rand.NewSource(time.Now().UnixNano()))

func GetRandomUserAgent() string {
	return userAgent[r.Intn(len(userAgent))]
}
