package main

import (
	"math/rand"
	"time"
	"strconv"
	"fmt"
)
func main() {
	//双色球：6个数，不重复，不为0，1~33，蓝色1~16
	res := getNum()
	fmt.Println(res)
}
func getNum() string {
	var resNum string
	//设置随机数的种子
	rand.Seed(time.Now().UnixNano())
	//定义一个放红球的map
	intMap := make(map[int]string, 7)
	//定义一个蓝球随机数
	bNum := rand.Intn(16)
	for {
		//定义一个红球随机数
		rNum := rand.Intn(33) + 1
		//不用纠结，此处的"a"没意义，打酱油的value
		intMap[rNum] = "a"
		if len(intMap) ==6 {
			break
		}
	}
	for k, _ := range intMap {
		//转成字符串
		resNum += strconv.Itoa(k) + ","
	}
	return resNum + strconv.Itoa(bNum)
}
