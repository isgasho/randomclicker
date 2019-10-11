package game

import (
	"errors"
	"fmt"
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

	g.buyItemHook()

	if err := g.defaultHooks(); err != nil {
		return err
	}
	return nil
}

func (g *Game) buyItemHook() {

	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		itm := getItemAtPoint(x, y)
		if itm != nil {
			itm.Level++
			g.CalcIncome()
		}
	}

}

func getItemAtPoint(x, y int) *data.Item {

	for i, v := range data.Items {
		if x > v.TopLeft.X && y > v.TopLeft.Y {
			if x < v.BotRight.X && y < v.BotRight.Y {
				fmt.Println(i)
				return v
			}
		}
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
