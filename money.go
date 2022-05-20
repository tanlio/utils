package utils

import (
	"github.com/shopspring/decimal"
	"math/big"
)

func Fen2Yuan(price int64) string {
	d := decimal.New(1, 2)
	result := decimal.NewFromFloat(float64(price)).DivRound(d, 2).String()
	return result
}

func Yuan2Fen(price float64) int64 {
	d := decimal.New(1, 2)
	df := decimal.NewFromFloat(price).Mul(d)

	return df.IntPart()
}

func HexToBigInt(hex string) *big.Int {
	n := new(big.Int)
	n, _ = n.SetString(hex, 16)

	return n
}
