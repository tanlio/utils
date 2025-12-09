package utils

import (
	"fmt"

	"github.com/shopspring/decimal"
)

func Fen2Yuan(price int64, args ...int32) string {
	scale := int32(2)
	if len(args) > 0 && args[0] >= 0 {
		scale = args[0]
	}

	d := decimal.New(1, scale)
	return decimal.NewFromInt(price).DivRound(d, scale).String()
}

func Yuan2Fen(price any, args ...int32) int64 {
	scale := int32(2)
	if len(args) > 0 && args[0] >= 0 {
		scale = args[0]
	}

	var priceStr string
	switch v := price.(type) {
	case string:
		priceStr = v
	case float32, float64:
		priceStr = fmt.Sprintf("%v", v)
	case int, int64, int32:
		priceStr = fmt.Sprintf("%v", v)
	default:
		return 0
	}

	dPrice, err := decimal.NewFromString(priceStr)
	if err != nil {
		return 0
	}

	d := decimal.New(1, scale)
	result := dPrice.Mul(d).Round(0)

	return result.IntPart()
}
