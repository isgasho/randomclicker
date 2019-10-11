package game

import (
	"errors"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
	"math/big"
	"randomclicker/data"
)

var one = big.NewFloat(1.0)

func (g *Game) ClickingHook() error {
	g.hookMouseClick()
	if err := g.defaultHooks(); err != nil {
		return err
	}
	return nil
}

func (g *Game) MenuHook() error {
	if err := g.defaultHooks(); err != nil {
		return err
	}
	return nil
}

func (g *Game) defaultHooks() error {
	g.hookA()
	if err := g.hookEsc(); err != nil {
		return err
	}
	return nil
}

func (g *Game) hookMouseClick() {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		g.counter.Add(g.counter, one)
		g.IncomeModifier = *g.IncomeModifier.Add(&g.IncomeModifier, one)
	}
}

func (g *Game) hookEsc() error {

	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		// g.TickerDone <- true
		g.Ticker.Stop()
		return errors.New("closing")
	}
	return nil
}

func (g *Game) hookA() {
	if inpututil.IsKeyJustPressed(ebiten.KeyA) {
		switch g.WindowState {
		case data.Menu:
			g.WindowState = data.Clicking
		case data.Clicking:
			g.WindowState = data.Menu
		}
	}
}
