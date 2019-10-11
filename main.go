package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten"
	"log"
	"randomclicker/data"
	"randomclicker/game"
)

var Game *game.Game

func init() {
	Game = game.NewGame()
	Game.PassiveIncome()

	ebiten.SetRunnableInBackground(true)

}

func update(screen *ebiten.Image) error {

	if !data.DidScale {
		ebiten.SetScreenSize(data.Width, data.Height)
		data.DidScale = true
		return nil

	}
	if ebiten.IsDrawingSkipped() {
		return nil
	}

	// if !data.DidScale {
	// 	ebiten.SetScreenSize(data.Width, data.Height)
	// 	data.DidScale = true
	// }

	// dunno if dis is needed
	// if Game.Screen == nil {
	fmt.Printf("%#v\n", screen)
	// }

	if data.DidScale {
		Game.Screen = screen
		return Game.Display()
	}
	return nil

}

func main() {
	ebiten.SetWindowDecorated(false)
	data.Width, data.Height = ebiten.ScreenSizeInFullscreen()
	if err := ebiten.Run(update, 320, 240, 2, "RandomClicker"); err != nil {
		log.Fatal(err)
	}
}
