package main

// Author:	DragonChen https://github.com/dragonchen-tw/
// Title:	Memory Board
// Date:	2023/09/18

import (
	"log"

	"github.com/dragonchen-tw/go-games/memoryboard/memoryboard"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	game, err := memoryboard.NewGame()
	if err != nil {
		log.Fatal(err)
	}

	ebiten.SetWindowSize(memoryboard.ScreenWidth, memoryboard.ScreenHeight)
	ebiten.SetWindowTitle("Memory Board")

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
