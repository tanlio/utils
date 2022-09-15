package utils

import (
	"github.com/shopspring/decimal"
	"math"
	"math/rand"
	"reflect"
	"time"
)

var Rander = rand.New(rand.NewSource(time.Now().UnixNano()))

//三目运算

func IfInt(cond bool, trueInt int, falseInt int) int {
	if cond {
		return trueInt
	} else {
		return falseInt
	}
}

func IfBool(cond bool, trueBool bool, falseBool bool) bool {
	if cond {
		return trueBool
	} else {
		return falseBool
	}
}

func IfInt64(cond bool, trueInt64 int64, falseInt64 int64) int64 {
	if cond {
		return trueInt64
	} else {
		return falseInt64
	}
}

func IfString(cond bool, trueString string, falseString string) string {
	if cond {
		return trueString
	} else {
		return falseString
	}
}

func IfFloat64(cond bool, trueFloat64 float64, falseFloat64 float64) float64 {
	if cond {
		return trueFloat64
	} else {
		return falseFloat64
	}
}

// 加法

func FloatAdd(f1, f2 float64) float64 {
	ff, _ := decimal.NewFromFloat(f1).Add(decimal.NewFromFloat(f2)).Float64()
	return ff
}

// 减法

func FloatSub(f1, f2 float64) float64 {
	ff, _ := decimal.NewFromFloat(f1).Sub(decimal.NewFromFloat(f2)).Float64()
	return ff
}

// 乘法

func FloatMul(f1, f2 float64) float64 {
	ff, _ := decimal.NewFromFloat(f1).Mul(decimal.NewFromFloat(f2)).Float64()
	return ff
}

// 除法

func FloatDiv(f1, f2 float64) float64 {
	ff, _ := decimal.NewFromFloat(f1).Div(decimal.NewFromFloat(f2)).Float64()
	return ff
}

func Timestamp2Str(timestamp int64) string {
	if timestamp == 0 {
		return ""
	}
	return time.Unix(timestamp, 0).Format("2006-01-02 15:04:05")
}

func Str2Timestamp(str string) int64 {
	locationTime, _ := time.ParseInLocation("2006-01-02 15:04:05", str, time.Local)
	return locationTime.Unix()
}

//计算两点经纬度之间距离

func EarthDistance(lat1, lng1, lat2, lng2 float64) float64 {
	radius := 6378.137 //km
	rad := math.Pi / 180.0
	lat1 = lat1 * rad
	lng1 = lng1 * rad
	lat2 = lat2 * rad
	lng2 = lng2 * rad
	theta := lng2 - lng1
	dist := math.Acos(math.Sin(lat1)*math.Sin(lat2) + math.Cos(lat1)*math.Cos(lat2)*math.Cos(theta))
	return dist * radius
}

const (
	Lowercase int = iota + 1
	CapitalLetter
	LowercaseAndCapitalLetter
	Number
	NumberAndLowercase
	NumberAndCapitalLetter
	NumberAndLowercaseAndCapitalLetter
)

func RandString(strType int, args ...interface{}) string {
	n := 0
	if len(args) > 0 && reflect.TypeOf(args).String() == "int" {
		n = args[0].(int)
	}
	if n == 0 {
		n = Rander.Intn(10) + 5
	}
	var sourceStr, targetStr string
	switch strType {
	case Lowercase:
		sourceStr = "abcdefghijklmnopqrstuvwxyz"
	case CapitalLetter:
		sourceStr = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	case LowercaseAndCapitalLetter:
		sourceStr = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	case Number:
		sourceStr = "0123456789"
	case NumberAndLowercase:
		sourceStr = "0123456789abcdefghijklmnopqrstuvwxyz"
	case NumberAndCapitalLetter:
		sourceStr = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	case NumberAndLowercaseAndCapitalLetter:
		sourceStr = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	}

	for i := 0; i < n && len(sourceStr) > 0; i++ {
		randNum := Rander.Intn(len(sourceStr))
		targetStr += string(sourceStr[randNum])
	}

	return targetStr
}
