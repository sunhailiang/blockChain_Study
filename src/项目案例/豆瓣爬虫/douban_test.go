package douban

import (
	"fmt"
	"testing"
	"time"
)

func TestDouban(t *testing.T) {
	ti := time.Now()
	Convert()
	elapsed := time.Since(ti)
	fmt.Println("爬取结束,总共耗时: ", elapsed)
}
