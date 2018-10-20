package mygo

import (
	"bytes"
	"math"
	"reflect"
	"strconv"
	"time"
)

/*[]string for string*/
func ByteString(p []string) string {
	var resu string
	for _, q := range p {
		resu = q
	}
	return resu
}

//结构体转为map
func StructToMap(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		data[t.Field(i).Name] = v.Field(i).Interface()
	}
	return data
}

//float小数位
func GetAdsStatisCount(num1 int64, num2 int64) float64 {
	var resu float64
	resu = float64(num1) / float64(num2)
	resu = math.Trunc(resu*1e2+0.5) * 1e-2
	return resu
}

func MyGoDateFormate(startTs int64, endTs int64) []string {
	timeLayout := "2006-01-02"           //转化所需模板
	loc, _ := time.LoadLocation("Local") //重要：获取时区
	var buf bytes.Buffer

	startTsTime := time.Unix(startTs/1e3, 0).Format(timeLayout)
	endTsTime := time.Unix(endTs/1e3, 0).Format(timeLayout)

	var adsDateList []string
	startTime, _ := time.ParseInLocation(timeLayout, startTsTime, loc) //使用模板在对应时区转化为time.time类型
	endTime, _ := time.ParseInLocation(timeLayout, endTsTime, loc)

	subDate := time.Time.Sub(endTime, startTime).Hours()

	subDay := int(subDate / 24)

	for i := 0; i <= subDay+1; i++ {
		hour := strconv.Itoa(24 * i)
		buf.WriteString(hour)
		buf.WriteString("h")
		hh, _ := time.ParseDuration(buf.String())
		buf.Reset()
		startSum := startTime.Add(hh)
		startSr := startSum.Unix()                              //转化为时间戳 类型是int64
		startTimeSr := time.Unix(startSr, 0).Format(timeLayout) //设置时间戳 使用模板格式化为日期字符串
		adsDateList = append(adsDateList, startTimeSr)
	}
	return adsDateList
}

//时间戳转换为日期输出
func MyGoDateFormateUnix(timestamp int64) string {
	timeLayout := "2006-01-02" //转化所需模板
	tampTimeSr := time.Unix(timestamp/1e3, 0).Format(timeLayout)
	return tampTimeSr
}
