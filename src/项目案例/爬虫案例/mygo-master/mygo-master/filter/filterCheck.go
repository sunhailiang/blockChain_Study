package mygo

import (
	"fmt"
	"strings"

	boom "github.com/BoomFilters"
	"github.com/bloom"
	cuckoo "github.com/goCuckoo"
)

func macuckooin() {
	// speicify capacity
	//cuckoo
	filter := cuckoo.NewFilter(10000)
	filter.Insert([]byte("zheng-ji,stupid"))
	//filter.Insert([]byte("stupid"))
	filter.Insert([]byte("coder"))
	if filter.Find([]byte("stupid")) {
		fmt.Println("exist")
	} else {
		fmt.Println("Not exist")
	}
	filter.Del([]byte("stupid"))
	fmt.Println(filter.Size())

}

func wilifiBloomFilterTest() {
	//wilifi/bloom
	n := uint(1000)
	filter := bloom.New(20*n, 5) // load of 20, 5 keys
	// filter.Add([]byte("Love"))
	// filter.Add([]byte("swxctx"))
	filter.AddString("love")
	if filter.Test([]byte("love")) {
		fmt.Println("exist")
	} else {
		fmt.Println("no exists")
	}
}

func bloomFiterTest2() {
	sbf := boom.NewDefaultStableBloomFilter(10000, 0.01)

	sbf.Add([]byte(`a`))
	if sbf.Test([]byte(`a`)) {
		fmt.Println("contains a")
	}

	if !sbf.TestAndAdd([]byte(`b`)) {
		fmt.Println("doesn't contain b")
	}

	if sbf.Test([]byte(`b`)) {
		fmt.Println("now it contains b!")
	}

	// Restore to initial state.
	sbf.Reset()
}

// 加载禁用词
func bloomFilterTest() {
	var BanWords string = "helloword"
	Title := "word"
	var adsTitle []string
	titleRune := []rune(Title)
	for m := 0; m < len(titleRune)-1; m++ {
		fmt.Println(string(titleRune[m : m+2]))
		adsTitle = append(adsTitle, string(titleRune[m:m+2]))
	}
	fmt.Println(len(adsTitle))

	s := strings.Split(BanWords, ",")
	fmt.Println(s)
	//cuckoo
	//filter := cuckoo.NewFilter(100000)
	n := uint(1000)
	filter := bloom.New(20*n, uint(len(s))) // load of 20, 5 keys
	for i := 0; i < len(s); i++ {
		filter.Add([]byte(s[i]))
	}

	for titleLen := 0; titleLen < len(adsTitle); titleLen++ {
		if filter.Test([]byte(adsTitle[titleLen])) {
			fmt.Println("exist")
		} else {
			fmt.Println("Not exist")
		}
	}

}
