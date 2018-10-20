package base

import (
	"strings"
	"time"
)

// 计算年龄
func GetYearDiffer(start_time, end_time string) int {
	var Age int64
	var pslTime string
	if strings.Index(start_time, ".") != -1 {
		pslTime = "2006.01.02"
	} else if strings.Index(start_time, "-") != -1 {
		pslTime = "2006-01-02"
	} else {
		pslTime = "2006/01/02"
	}
	t1, err := time.ParseInLocation(pslTime, start_time, time.Local)
	t2, err := time.ParseInLocation("2006-01-02", end_time, time.Local)

	if err == nil && t1.Before(t2) {
		diff := t2.Unix() - t1.Unix()
		Age = diff / (3600 * 365 * 24)
		return int(Age)
	} else {
		return int(Age)
	}
}
