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

	if ebiten.IsDrawingSkipped() {
		return nil
	}

	// dunno if dis is needed
	// if Game.Screen == nil {
	fmt.Printf("%#v\n", screen)
	// }

	Game.Screen = screen
	return Game.Display()

}

func main() {
	ebiten.SetWindowDecorated(false)
	data.Width, data.Height = ebiten.ScreenSizeInFullscreen()

	if err := ebiten.Run(update, data.Width, data.Height, 1, "RandomClicker"); err != nil {
		log.Fatal(err)
	}
}
