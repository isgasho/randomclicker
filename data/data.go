package data

import (
	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten/examples/resources/fonts"
	"golang.org/x/image/font"
	"log"
	"math/big"
)

var Counter = 0

var (
	SmallFont  font.Face
	NormalFont font.Face
	BigFont    font.Face
)

var (
	Height int
	Width  int
)

type WindowState int

const (
	Clicking WindowState = iota
	Menu
)

type Item struct {
	Level             float64
	BaseValue         float64
	BaseMultiplier    float64    // mod for inc/sec
	currentMultiplier float64    // BaseMultiplier ^ Level
	BasePrice         float64    // base buy price (modded by PriceMultiplier)
	currentPrice      *big.Float // BasePrice * PriceMultiplier
	PriceMultiplier   float64
}

var Items = map[string]Item{
	"nr1": Item{
		Level:           0,
		BaseValue:       2,
		BaseMultiplier:  1.03,
		BasePrice:       10,
		PriceMultiplier: 1.2,
	},
}

func init() {

	tt, err := truetype.Parse(fonts.MPlus1pRegular_ttf)
	if err != nil {
		log.Fatal(err)
	}

	const dpi = 72
	SmallFont = truetype.NewFace(tt, &truetype.Options{
		Size:    24,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	NormalFont = truetype.NewFace(tt, &truetype.Options{
		Size:    32,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	BigFont = truetype.NewFace(tt, &truetype.Options{
		Size:    48,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
}
