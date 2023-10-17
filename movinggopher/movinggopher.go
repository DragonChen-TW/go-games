package main

// Author:	DragonChen https://github.com/dragonchen-tw/
// Title:	Moving Gopher
// Date:	2023/09/18

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var img *ebiten.Image
var imgW, imgH int
var imgX, imgY float64
var direction string = "right"

func init() {
	var err error
	img, _, err = ebitenutil.NewImageFromFile("imgs/gopher.png")
	if err != nil {
		log.Fatal(err)
	}

	// get imgW imgH, and half them
	imgW, imgH = img.Bounds().Dx()/2, img.Bounds().Dy()/2
	imgX, imgY = 0, 0
}

type Game struct{}

func (g *Game) Update() error {
	// depending on direction, change imgX imgY
	var movePixels float64 = 3

	switch direction {
	case "right":
		if imgX > float64(640-imgW) {
			direction = "down"
			break
		}
		imgX += movePixels
	case "down":
		if imgY > float64(480-imgH) {
			direction = "left"
			break
		}
		imgY += movePixels
	case "left":
		if imgX < 0 {
			direction = "up"
			break
		}
		imgX -= movePixels
	case "up":
		if imgY < 0 {
			direction = "right"
			break
		}
		imgY -= movePixels
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(0.5, 0.5)
	op.GeoM.Translate(imgX, imgY)
	screen.DrawImage(img, op)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 480
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")

	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
