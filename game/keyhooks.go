package game

import (
	"errors"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
	"math/big"
	"time"
)

var one = big.NewFloat(1.0)

func (g *Game) Hook() error {

	g.hookMouseClick()
	g.hookA()
	if err := g.hookEsc(); err != nil {
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
				// fmt.Println("Tick at", t)
				// fmt.Printf("%#v\n", g.Screen)
				g.counter.Add(g.counter, &g.IncomeModifier)
			}
		}
	}()

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
		g.Menu()
	}
}
