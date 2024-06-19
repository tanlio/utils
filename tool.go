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

func IfInt(condition bool, trueInt int, falseInt int) int {
	if condition {
		return trueInt
	} else {
		return falseInt
	}
}

func IfBool(condition bool, trueBool bool, falseBool bool) bool {
	if condition {
		return trueBool
	} else {
		return falseBool
	}
}

func IfInt64(condition bool, trueInt64 int64, falseInt64 int64) int64 {
	if condition {
		return trueInt64
	} else {
		return falseInt64
	}
}

func IfString(condition bool, trueString string, falseString string) string {
	if condition {
		return trueString
	} else {
		return falseString
	}
}

func IfFloat64(condition bool, trueFloat64 float64, falseFloat64 float64) float64 {
	if condition {
		return trueFloat64
	} else {
		return falseFloat64
	}
}

func If(condition bool, trueVal, falseVal interface{}) interface{} {
	if condition {
		return trueVal
	}
	return falseVal
}

func IfTernary(a bool, b, c interface{}) interface{} {
	if a {
		return b
	}
	return c
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
	if f1 == 0 || f2 == 0 {
		return 0
	}
	ff, _ := decimal.NewFromFloat(f1).Div(decimal.NewFromFloat(f2)).Float64()
	return ff
}

func Timestamp2Str(timestamp int64, args ...interface{}) string {
	if timestamp == 0 {
		return ""
	}

	format := time.DateTime
	if len(args) > 0 && reflect.TypeOf(args[0]).String() == "string" {
		format = args[0].(string)
	}
	return time.Unix(timestamp, 0).Format(format)
}

func Str2Timestamp(str string, args ...interface{}) int64 {
	format := time.DateTime
	if len(args) > 0 && reflect.TypeOf(args[0]).String() == "string" {
		format = args[0].(string)
	}
	locationTime, _ := time.ParseInLocation(format, str, time.Local)
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

func HaversineDistance(lat1, lon1, lat2, lon2 float64) float64 {
	const EarthRadius = 6371.0 // Radius of the Earth in kilometers
	rad := math.Pi / 180.0
	dLat := (lat2 - lat1) * rad
	dLon := (lon2 - lon1) * rad

	lat1 = lat1 * rad
	lat2 = lat2 * rad

	a := math.Sin(dLat/2)*math.Sin(dLat/2) + math.Cos(lat1)*math.Cos(lat2)*math.Sin(dLon/2)*math.Sin(dLon/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	return EarthRadius * c
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

func SliceToSlice[T any, V any](array []T, iteratee func(T) V) []V {
	result := make([]V, 0)
	for _, item := range array {
		k := iteratee(item)
		result = append(result, k)
	}

	return result
}
