package game

import (
	"fmt"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/text"
	"image/color"
	"randomclicker/data"
)

var menu *ebiten.Image
var clicking *ebiten.Image

func (g *Game) Display() error {

	if clicking == nil {
		clicking = g.Screen
	}

	// g.Screen.Clear()

	switch g.WindowState {
	case data.Clicking:
		return g.Clicking()
	case data.Menu:
		return g.Menu()
	}
	return nil
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

	counter := 1

	for i, v := range data.Items {
		text.Draw(image, fmt.Sprintf("%s :", i), data.SmallFont, 0, counter*30, color.Black)
		text.Draw(image, fmt.Sprintf("Price: %.3g ", v.CurrentPrice()), data.SmallFont, 20, counter*60, color.Black)
		text.Draw(image, fmt.Sprintf("Multiplier: %.3g ", v.CurrentMultiplier()), data.SmallFont, 20, counter*90, color.Black)
		counter += 3
	}

	text.Draw(image, "menu:", data.NormalFont, 0, 24, color.Black)
	menu = image

}
