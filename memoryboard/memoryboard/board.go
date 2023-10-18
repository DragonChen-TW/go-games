package memoryboard

import (
	"fmt"
	"image/color"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type Board struct {
	posX, posY int

	nrow, ncol int
	tiles      [][]Tile
	guessed    [][]bool

	isLastFlip bool
	lastFlipI  int
	lastFlipJ  int
	isCurFlip  bool
	curFlipI   int
	curFlipJ   int

	guessTimer int64
}

func GenTiles(nrow, ncol int) [][]Tile {
	var pairs []int = make([]int, nrow*ncol)
	for i := 0; i < nrow*ncol; i++ {
		pairs[i] = i / 2
	}
	rand.Shuffle(nrow*ncol, func(a, b int) {
		pairs[a], pairs[b] = pairs[b], pairs[a]
	})

	var tiles [][]Tile = make([][]Tile, nrow)
	var tileColors = make([]color.Color, (nrow*ncol)/2)
	for i := range tileColors {
		tileColors[i] = color.RGBA{
			128 + uint8(rand.Intn(128)),
			128 + uint8(rand.Intn(128)),
			128 + uint8(rand.Intn(128)),
			0xff,
		}
	}
	for i := 0; i < nrow; i++ {
		tiles[i] = make([]Tile, ncol)
		for j := 0; j < ncol; j++ {
			value := pairs[i*ncol+j]
			tiles[i][j] = Tile{
				data: TileData{
					value: value,
					i:     i,
					j:     j,
				},
				color: tileColors[value],
			}
		}
	}
	return tiles
}

func NewBoard(nrow, ncol int, screenW, screenH int) (*Board, error) {
	var guessed = make([][]bool, nrow)
	for i := 0; i < nrow; i++ {
		guessed[i] = make([]bool, ncol)
	}

	tiles := GenTiles(nrow, ncol)

	b := &Board{
		nrow:    nrow,
		ncol:    ncol,
		tiles:   tiles,
		guessed: guessed,
	}
	bx, by := b.Size()
	bx = (screenW - bx) / 2
	by = (screenH - by) / 2
	fmt.Println("BoardXY", bx, by)
	b.posX = bx
	b.posY = by
	return b, nil
}

func (b *Board) Size() (int, int) {
	x := b.nrow*tileSize + (b.nrow+1)*tileMargin
	y := b.ncol*tileSize + (b.ncol+1)*tileMargin
	return x, y
}

func (b *Board) WhichTileClicked(clickX, clickY int) (*Tile, bool) {
	clickX -= b.posX
	clickY -= b.posY
	fmt.Println("Relative XY", clickX, clickY)

	// out of the bound of board
	tileSizeWithMargin := tileSize + tileMargin
	if clickX < 0 || clickX/tileSizeWithMargin >= b.nrow ||
		clickY < 0 || clickY/tileSizeWithMargin >= b.ncol {
		fmt.Println("Out")
	} else if clickX%tileSizeWithMargin >= tileMargin && clickY%tileSizeWithMargin >= tileMargin {
		fmt.Println("In")
		tileX := clickX / tileSizeWithMargin
		tileY := clickY / tileSizeWithMargin
		return &b.tiles[tileX][tileY], true
	} else {
		fmt.Println("Border")
	}
	return nil, false
}

func (b *Board) Update(input *Input) error {
	if b.guessTimer > 0 && b.guessTimer <= time.Now().UnixMilli() {
		b.isCurFlip = false
		b.isLastFlip = false
		b.guessTimer = 0
	}

	if b.guessTimer == 0 {
		if clickX, clickY, ok := input.ClickUp(); ok {
			fmt.Println("Click", clickX, clickY)
			if clickedTile, ok := b.WhichTileClicked(clickX, clickY); ok {
				b.Guess(clickedTile)
			}
		}
	}

	return nil
}

func (b *Board) Guess(guessTile *Tile) bool {
	guessI := guessTile.data.i
	guessJ := guessTile.data.j
	if b.isLastFlip {
		lastFlipTile := b.tiles[b.lastFlipI][b.lastFlipJ]
		if b.guessed[guessI][guessJ] {
			fmt.Println("This tile has been guessed")
		} else if *guessTile == lastFlipTile {
			fmt.Println("Click on the same tiles")
		} else if guessTile.data.value == lastFlipTile.data.value {
			// Correct guess
			fmt.Println("Correct Guess")
			b.guessed[guessI][guessJ] = true
			b.guessed[b.lastFlipI][b.lastFlipJ] = true
			return true
		} else {
			// Wrong guess
			fmt.Println("Wrong Guess")
			b.isCurFlip = true
			b.curFlipI = guessI
			b.curFlipJ = guessJ

			b.guessTimer = time.Now().UnixMilli() + 1000
		}
	} else {
		b.isLastFlip = true
		b.lastFlipI = guessI
		b.lastFlipJ = guessJ
	}
	return false
}

func (b *Board) Draw(boardImage *ebiten.Image) {
	boardImage.Fill(color.White)
	for i := 0; i < b.nrow; i++ {
		for j := 0; j < b.ncol; j++ {
			op := &ebiten.DrawImageOptions{}
			x := i*tileSize + (i+1)*tileMargin
			y := j*tileSize + (j+1)*tileMargin
			op.GeoM.Translate(float64(x), float64(y))
			boardImage.DrawImage(tileImage, op)

			if b.guessed[i][j] || (b.isLastFlip && i == b.lastFlipI && j == b.lastFlipJ) ||
				(b.isCurFlip && i == b.curFlipI && j == b.curFlipJ) {
				b.tiles[i][j].Draw(boardImage)
			}
		}
	}
}
