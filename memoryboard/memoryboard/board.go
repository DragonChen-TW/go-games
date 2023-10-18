package memoryboard

import (
	"image/color"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

type Board struct {
	nrow, ncol int
	tiles      [][]Tile
	guessed    [][]bool
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

func NewBoard(nrow, ncol int) (*Board, error) {
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
	return b, nil
}

func (b *Board) Size() (int, int) {
	x := b.nrow*tileSize + (b.nrow+1)*tileMargin
	y := b.ncol*tileSize + (b.ncol+1)*tileMargin
	return x, y
}

func (b *Board) Update(input *Input) error {
	return nil
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

			if b.guessed[i][j] {
				b.tiles[i][j].Draw(boardImage)
			}
		}
	}

	// b.tiles[0][0].Draw(boardImage)
}
