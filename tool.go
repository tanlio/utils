package utils

import (
	"github.com/shopspring/decimal"
	"math"
	"math/rand"
	"path/filepath"
	"reflect"
	"runtime"
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

type RandType int

const (
	Lowercase RandType = iota + 1
	CapitalLetter
	LowercaseAndCapitalLetter
	Number
	NumberAndLowercase
	NumberAndCapitalLetter
	NumberAndLowercaseAndCapitalLetter
)

func RandString(strType RandType, args ...interface{}) string {
	n := Rander.Intn(10) + 5
	if len(args) > 0 && reflect.TypeOf(args[0]).String() == "int" {
		n = args[0].(int)
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

func GenerateRedPacket(totalAmount int64, num int64, maxAmount int64, minAmount int64, rng *rand.Rand) []int64 {
	result := make([]int64, num)
	var averageAmount = totalAmount / num
	if maxAmount < averageAmount || minAmount > averageAmount {
		return result
	}

	if rng == nil {
		rng = Rander
	}

	remainAmount := totalAmount - minAmount*num
	for i := 0; i < int(num); i++ {
		var amount int64
		if i == int(num-1) {
			amount = remainAmount
		} else {
			if remainAmount != 0 {
				tempRandAmount := IfInt64(remainAmount > maxAmount-minAmount, maxAmount-minAmount, remainAmount)
				amount = rng.Int63n(tempRandAmount)
			}
		}

		result[i] = minAmount + amount
		remainAmount -= amount
	}

	rng.Shuffle(len(result), func(i, j int) {
		result[i], result[j] = result[j], result[i]
	})

	return result
}

func GetProjectPath() string {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("Failed to get current file path")
	}

	// 获取当前文件所在目录
	dir := filepath.Dir(filename)

	// 逐级向上遍历，直到找到项目根目录（假设项目根目录为包含"go.mod"文件的目录）
	for {
		files, err := filepath.Glob(filepath.Join(dir, "go.mod"))
		if err == nil && len(files) > 0 {
			return dir
		}

		// 到达文件系统根目录时停止遍历
		if dir == "/" || dir == "." {
			break
		}

		dir = filepath.Dir(dir)
	}

	panic("Failed to find project root directory")
}
