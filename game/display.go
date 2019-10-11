package game

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/text"
	"image/color"
	"randomclicker/data"
)

var menu *ebiten.Image
var clicking *ebiten.Image

func (g *Game) Display() {

	if clicking == nil {
		clicking = g.Screen
	}
	opts := &ebiten.DrawImageOptions{
		CompositeMode: ebiten.CompositeModeCopy,
	}

	g.Screen.Clear()

	switch g.WindowState {
	case data.Clicking:
		if g.Screen != clicking {
			g.Screen.DrawImage(clicking, opts)

		}
		g.Counter()
		g.IncomePerSecond()
	case data.Menu:
		g.Screen.DrawImage(menu, opts)
	}
}

func (g *Game) Counter() {
	text.Draw(g.Screen, g.counter.String(), data.SmallFont, 100, 100, color.White)
}

func (g *Game) IncomePerSecond() {
	text.Draw(g.Screen, "Income/Sec: "+g.IncomeModifier.String(), data.SmallFont, 0, 30, color.White)
}

func init() {
	image, err := ebiten.NewImage(200, 200, ebiten.FilterDefault)
	if err != nil {
		panic(err)
	}

	image.Fill(color.White)

	text.Draw(image, "menu:", data.NormalFont, 0, 24, color.Black)
	menu = image

}

func (g *Game) Menu() {
	switch g.WindowState {
	case data.Menu:
		g.WindowState = data.Clicking
	case data.Clicking:
		g.WindowState = data.Menu
	}
}
