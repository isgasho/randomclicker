package main

import (
	"github.com/hajimehoshi/ebiten"
	"log"
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
	Game.Screen = screen
	// }

	Game.Display()

	if err := Game.Hook(); err != nil {
		return err
	}

	return nil
}

func main() {

	if err := ebiten.Run(update, 320, 240, 2, "RandomClicker"); err != nil {
		log.Fatal(err)
	}
}
