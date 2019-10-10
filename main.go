package main

import (
	"github.com/hajimehoshi/ebiten"
	"log"
	"rnggameshit/game"
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

	if Game.Screen == nil {
		Game.Screen = screen
	}

	if err := Game.Hook(); err != nil {
		return err
	}
	Game.Counter()
	Game.IncomePerSecond()

	return nil
}

func main() {

	if err := ebiten.Run(update, 320, 240, 2, "Hello, World!"); err != nil {
		log.Fatal(err)
	}
}