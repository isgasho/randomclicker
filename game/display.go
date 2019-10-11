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

func (g *Game) Counter() {
	text.Draw(g.Screen, g.counter.String(), data.SmallFont, 100, 100, color.White)
}

func (g *Game) IncomePerSecond() {
	text.Draw(g.Screen, "Income/Sec: "+g.IncomeModifier.String(), data.SmallFont, 0, 30, color.White)

}

func (g *Game) Menu() {

	// width, height := g.Screen.Size()

	if menu == nil {
		fmt.Println("created menu")

		image, err := ebiten.NewImage(200, 200, ebiten.FilterDefault)
		if err != nil {
			panic(err)
		}

		text.Draw(image, "menu:", data.BigFont, 0, 24, color.Black)
		image.Fill(color.White)
		menu = image
	}

	if clicking == nil {
		clicking = g.Screen
	}
	// g.Screen.Clear()
	opts := ebiten.DrawImageOptions{
		// CompositeMode: ebiten.CompositeModeDestination,
	}

	fmt.Println("before:" + fmt.Sprintf("\t%s\t%v", g.WindowState, g.Screen))
	if g.WindowState == data.Menu {
		menu = g.Screen
		g.WindowState = data.Clicking
		g.Screen.DrawImage(clicking, &opts)
		// g.Screen = clicking
		// g.Screen.DrawImage(clicking)

	} else if g.WindowState == data.Clicking {
		g.WindowState = data.Menu
		clicking = g.Screen
		g.Screen.DrawImage(menu, &opts)
		// g.Screen = menu
		// ebitenutil.DebugPrint(g.Screen, "reee")
	}

	fmt.Println("after:" + fmt.Sprintf("\t%s\t%v", g.WindowState, g.Screen))

}
