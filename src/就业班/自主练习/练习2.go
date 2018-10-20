package main

import "fmt"

func main() {
	//{"red", "black", "red", "pink", "blue", "pink", "blue"}
	//——>	{"red", "black", "pink", "blue"}

	var strSlice = []string{"red", "black", "red", "pink", "blue", "pink", "blue"}
	//strSlice = getNoRepeatSlice(strSlice)
	getNoRepeatSlice(strSlice)
	//fmt.Println(newSlice)

}

//方法一
//func getNoRepeatSlice(slice []string) []string {
//	var newSlice []string
//	var mp = make(map[string]string)
//
//	for i := 0; i < len(slice); i++ {
//		mp[slice[i]] = "a"
//	}
//	for k, _ := range mp {
//		newSlice = append(newSlice, k)
//	}
//	return newSlice
//}

//}
func getNoRepeatSlice(slice []string) {
	newSlice:=slice[:1]
	for i := 1; i < len(slice); i++ {
		for j := 0; j< len(newSlice); j++ {
			if newSlice[j] != slice[i] {
				newSlice = append(newSlice, slice[i])
			}
		}
	}
	fmt.Println(newSlice)
}
