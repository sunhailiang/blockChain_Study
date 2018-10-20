package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"regexp"
	"strconv"

	"github.com/shopspring/decimal"
)

func main() {

}

func ConvertIntToFloat(price int64) float64 {
	result, _ := decimal.New(price, 0).DivRound(decimal.New(100, 0), 2).Float64()
	return result
}

func ConvertPriceStringToInt(price string) int32 {
	m, err := decimal.NewFromString(price)
	if err != nil {
		fmt.Println(err)
	}
	result := m.Mul(decimal.New(100, 0)).IntPart()
	return int32(result)
}

func ConvertCommission(price, rate string) int32 {
	priceF, err := decimal.NewFromString(price)
	if err != nil {
		fmt.Errorf("convert to price error: price -> %s", price)
	}
	rateF, err := decimal.NewFromString(rate)
	if err != nil {
		fmt.Errorf("convert to price error: price -> %s", price)
	}
	result := priceF.Mul(rateF).IntPart()
	return int32(result)
}
func Converturl(uEnc string) []byte {
	uDec, err := base64.StdEncoding.DecodeString(regexp.MustCompile(" ").ReplaceAllString(uEnc, "+"))
	if err != nil {
		log.Fatalln(err)
	}
	return uDec
}

// ConvertStringToInt 字段类型转换
func ConvertStringToInt(data string) int32 {
	f, err := strconv.ParseFloat(data, 32)
	if err != nil {
		fmt.Printf("convert err ->%s", data)
	}
	return int32(f)
}

// GetDiscountRate 返回折扣率
func GetDiscountRate(couponPrice, originPrice int32) int32 {
	if couponPrice <= 0 || originPrice <= 0 {
		return 0
	}
	cPrice := decimal.NewFromFloat(float64(couponPrice))
	oPrice := decimal.NewFromFloat(float64(originPrice))
	n := decimal.NewFromFloat(float64(10000))
	return int32(cPrice.DivRound(oPrice, 4).Mul(n).IntPart())
}

// FloatStringToInt32 小数转整数,返回格式：小数*100
func FloatStringToInt32(s string) int32 {
	d, err := decimal.NewFromString(s)
	if err != nil {
		return 0
	}
	return int32(d.Mul(decimal.New(100, 0)).IntPart())
}

// FloatStringToFloat64 字符串转float
func FloatStringToFloat64(s string) float64 {
	d, err := decimal.NewFromString(s)
	if err != nil {
		return 0
	}
	num, _ := d.Mul(decimal.New(100, 0)).Float64()
	return num
}
