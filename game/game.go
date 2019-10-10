package game

import (
	"github.com/hajimehoshi/ebiten"
	"math/big"
	"randomclicker/data"
	"time"
)

type Game struct {
	Screen           *ebiten.Image
	counter          *big.Float
	Ticker           *time.Ticker
	TickerDone       chan bool
	ModifierPerClick big.Float
	IncomeModifier   big.Float
	WindowState      data.WindowState
}

func NewGame() *Game {
	return &Game{
		ModifierPerClick: *big.NewFloat(1.0),
		IncomeModifier:   *big.NewFloat(1.0),
		counter:          big.NewFloat(0.0),
		WindowState:      data.Clicking,
	}
}
