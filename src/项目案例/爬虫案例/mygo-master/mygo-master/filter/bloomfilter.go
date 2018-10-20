package pool

import (
	"fmt"
	"strings"

	"github.com/bloom"
)

func PoolFilter(title string) bool {

	var filterTwo []string
	var filterThree []string
	var filterFour []string

	strFilterWords := strings.Split(filterWords, ",")

	//bloom filter
	n := uint(1000)
	filter := bloom.New(20*n, uint(len(strFilterWords))) // load of 20, 5 keys
	//add filter
	for i := 0; i < len(strFilterWords); i++ {
		filter.Add([]byte(strFilterWords[i]))
	}

	filterRune := []rune(title)
	// fmt.Println(len(titleRune))
	if len(filterRune) >= 2 {
		for m := 0; m < len(filterRune)-1; m++ {
			filterTwo = append(filterTwo, string(filterRune[m:m+2]))
		}
		//check filter
		for titleLen := 0; filterLen < len(filterTwo); filterLen++ {
			if filter.Test([]byte(filterTwo[titleLen])) {
				//fmt.Println("exist")
				return true
			}
		}
	}

	if len(filterRune) >= 3 {
		for m := 0; m < len(filterRune)-2; m++ {
			filterThree = append(filterThree, string(filterRune[m:m+3]))
		}
		//check filter
		for filterLen := 0; filterLen < len(filterThree); filterLen++ {
			if filter.Test([]byte(filterThree[filterLen])) {
				fmt.Println("exist")
				return true
			}
		}
	}

	if len(filterRune) >= 4 {
		for m := 0; m < len(filterRune)-3; m++ {
			filterFour = append(filterFour, string(filterRune[m:m+3]))
		}
		//check filter
		for filterLen := 0; filterLen < len(filterFour); filterLen++ {
			if filter.Test([]byte(filterFour[filterLen])) {
				fmt.Println("exist")
				return true
			}
		}
	}
	return false
}
