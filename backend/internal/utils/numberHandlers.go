package utils

import "github.com/shopspring/decimal"

func ExtractFloat(d decimal.Decimal) float64 {
	f, _ := d.Float64()
	return f
}

func ToDecimal(v float64) decimal.Decimal {
	return decimal.NewFromFloat(v)
}
