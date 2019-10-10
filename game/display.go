package game

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/text"
	"image/color"
	"randomclicker/data"
)

var oldImg *ebiten.Image
var menu *ebiten.Image
var clicking *ebiten.Image

func (g *Game) Counter() {
	text.Draw(g.Screen, g.counter.String(), data.SmallFont, 100, 100, color.White)
}

func (g *Game) IncomePerSecond() {
	text.Draw(g.Screen, "Income/Sec: "+g.IncomeModifier.String(), data.SmallFont, 0, 30, color.White)
}

func (g *Game) Menu() {

	width, height := g.Screen.Size()

	if menu == nil {
		a, err := ebiten.NewImage(width, height, ebiten.FilterDefault)
		if err != nil {
			panic(err)
		}
		menu = a
	}

	if clicking == nil {
		clicking = g.Screen
	}

	if g.WindowState == data.Menu {
		g.Screen = clicking
	} else if g.WindowState == data.Clicking {
		g.Screen = menu
	}
}
