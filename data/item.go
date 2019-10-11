package data

import "math"

func (i *Item) CurrentMultiplier() float64 {
	// 	BaseMultiplier ^ Level
	return math.Pow(i.BaseMultiplier, i.Level)
}

func (i *Item) CurrentPrice() float64 {
	// BasePrice * PriceMultiplier
	return i.BasePrice * i.PriceMultiplier
}
