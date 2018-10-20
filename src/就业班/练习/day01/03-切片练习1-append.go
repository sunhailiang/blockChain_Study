package main

import "fmt"

func noEmpty(data []string) []string{
	out := data[:0]				// 在原切片上截取一个长度为 0 的切片 == make([]string, 0)
	for _, str := range data {
		if str != "" {
			out = append(out, str)
		}
		// 取到空字符串，不作为。
	}
	return out
}

// 直接在原串上操作
func noEmpty2(data []string) []string {
	i := 0
	for _, str := range data {
		if str != "" {
			data[i] = str
			i++
		}
		// 取到空字符串，不作为。
	}
	return data[:i]
}

func main()  {
	// {"red" "black", "pink", "", "", "pink", "blue"}
	data := []string{"red", "", "black", "", "", "pink", "blue"}
	afterData := noEmpty2(data)
	fmt.Println("afterData:", afterData)
}
