package game

import (
	"github.com/hajimehoshi/ebiten"
	"math/big"
	"time"
)

type Game struct {
	Screen           *ebiten.Image
	counter          *big.Float
	Ticker           *time.Ticker
	TickerDone       chan bool
	ModifierPerClick big.Float
	IncomeModifier   big.Float
}

func NewGame() *Game {
	return &Game{
		ModifierPerClick: *big.NewFloat(1.0),
		IncomeModifier:   *big.NewFloat(1.0),
		counter:          big.NewFloat(0.0),
	}
}
