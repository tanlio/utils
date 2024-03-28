package utils

import (
	"github.com/shopspring/decimal"
	"math/big"
	"reflect"
)

func Fen2Yuan(price int64, args ...interface{}) string {
	var num int32
	if len(args) > 0 && reflect.TypeOf(args[0]).String() == "int" {
		num = int32(args[0].(int))
	} else {
		num = 2
	}
	d := decimal.New(1, num)
	result := decimal.NewFromFloat(float64(price)).DivRound(d, num).String()
	return result
}

func Yuan2Fen(price float64, args ...interface{}) int64 {
	var num int32
	if len(args) > 0 && reflect.TypeOf(args[0]).String() == "int" {
		num = int32(args[0].(int))
	} else {
		num = 2
	}
	d := decimal.New(1, num)
	df := decimal.NewFromFloat(price).Mul(d)

	return df.IntPart()
}

func HexToBigInt(hex string) *big.Int {
	n := new(big.Int)
	n, _ = n.SetString(hex, 16)

	return n
}
