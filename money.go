package utils

import (
	"math/big"

	"github.com/shopspring/decimal"
)

func getScale(defaultScale int32, args ...int32) int32 {
	scale := defaultScale
	if len(args) > 0 && args[0] >= 0 {
		scale = args[0]
	}
	return scale
}

func decimalFromAny(price any) (decimal.Decimal, bool) {
	switch v := price.(type) {
	case nil:
		return decimal.Decimal{}, false

	case decimal.Decimal:
		return v, true
	case *decimal.Decimal:
		if v == nil {
			return decimal.Decimal{}, false
		}
		return *v, true
	case string:
		d, err := decimal.NewFromString(v)
		if err != nil {
			return decimal.Decimal{}, false
		}
		return d, true
	case float32:
		return decimal.NewFromFloat(float64(v)), true
	case float64:
		return decimal.NewFromFloat(v), true
	case int:
		return decimal.NewFromInt(int64(v)), true
	case int8:
		return decimal.NewFromInt(int64(v)), true
	case int16:
		return decimal.NewFromInt(int64(v)), true
	case int32:
		return decimal.NewFromInt(int64(v)), true
	case int64:
		return decimal.NewFromInt(v), true
	case *big.Int:
		if v == nil {
			return decimal.Decimal{}, false
		}
		return decimal.NewFromBigInt(v, 0), true
	default:
		return decimal.Decimal{}, false
	}
}

func Fen2Yuan(price any, args ...int32) string {
	scale := getScale(2, args...)

	dPrice, ok := decimalFromAny(price)
	if !ok {
		return "0"
	}
	if dPrice.IsZero() {
		return "0"
	}

	d := decimal.New(1, scale)
	return dPrice.DivRound(d, scale).String()
}

func Yuan2Fen(price any, args ...int32) int64 {
	scale := getScale(2, args...)

	dPrice, ok := decimalFromAny(price)
	if !ok {
		return 0
	}
	if dPrice.IsZero() {
		return 0
	}

	d := decimal.New(1, scale)
	result := dPrice.Mul(d).Round(0)

	return result.IntPart()
}

func Fen2YuanBig(price any, args ...int32) string {
	scale := getScale(8, args...)

	dPrice, ok := decimalFromAny(price)
	if !ok {
		return "0"
	}
	if dPrice.IsZero() {
		return "0"
	}

	d := decimal.New(1, scale)
	result := dPrice.DivRound(d, scale)

	return result.String()
}

func Yuan2FenBig(price any, args ...int32) *big.Int {
	scale := getScale(8, args...)

	dPrice, ok := decimalFromAny(price)
	if !ok {
		return big.NewInt(0)
	}
	if dPrice.IsZero() {
		return big.NewInt(0)
	}

	d := decimal.New(1, scale)
	result := dPrice.Mul(d).Round(0)

	return result.BigInt()
}
