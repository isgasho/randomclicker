package game

import (
	"fmt"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/text"
	"image/color"
	"randomclicker/data"
)

var menu *ebiten.Image
var Clicking *ebiten.Image

func (g *Game) Display() error {

	if Clicking == nil {
		Clicking = g.Screen
	}

	if menu == nil {
		otherInit()
	}

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

func otherInit() {
	image, err := ebiten.NewImage(data.Width, data.Height, ebiten.FilterDefault)
	if err != nil {
		panic(err)
	}
	image.Fill(color.White)

	nextPixel := 30

	for i, v := range data.Items {
		text.Draw(image, fmt.Sprintf("%s :", i), data.SmallFont, 0, nextPixel+30, color.Black)
		text.Draw(image, fmt.Sprintf("Price: %.3g ", v.CurrentPrice()), data.SmallFont, 20, nextPixel+60, color.Black)
		text.Draw(image, fmt.Sprintf("Multiplier: %.3g ", v.CurrentMultiplier()), data.SmallFont, 20, nextPixel+90, color.Black)

		topleft := &data.Point{
			X: 0,
			Y: nextPixel,
		}

		nextPixel += 90

		botright := &data.Point{
			X: 300,
			Y: nextPixel,
		}

		v.TopLeft = topleft
		v.BotRight = botright

		ebitenutil.DrawLine(image, 0, float64(nextPixel+5), 300, float64(nextPixel+5), color.Black)
	}

	ebitenutil.DrawLine(image, 300, 0, 300, float64(data.Height), color.Black)
	text.Draw(image, "menu:", data.NormalFont, 0, 24, color.Black)
	menu = image

}
