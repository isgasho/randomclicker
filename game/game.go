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

var opts = &ebiten.DrawImageOptions{
	CompositeMode: ebiten.CompositeModeCopy,
}

func NewGame() *Game {
	return &Game{
		ModifierPerClick: *big.NewFloat(1.0),
		IncomeModifier:   *big.NewFloat(1.0),
		counter:          big.NewFloat(0.0),
		WindowState:      data.Clicking,
	}
}

func (g *Game) Menu() error {
	if err := g.MenuHook(); err != nil {
		return err
	}



	g.Screen.DrawImage(menu, opts)
	return nil
}

func (g *Game) Clicking() error {

	if g.Screen != Clicking {
		g.Screen.DrawImage(Clicking, opts)
	}
	g.Counter()
	g.IncomePerSecond()
	if err := g.ClickingHook(); err != nil {
		return err
	}
	return nil
}

func (g *Game) PassiveIncome() {

	g.Ticker = time.NewTicker(time.Second * 1)

	go func() {
		for {
			select {
			case <-g.TickerDone:
				return
			case <-g.Ticker.C:
				g.counter.Add(g.counter, &g.IncomeModifier)
			}
		}
	}()

}
