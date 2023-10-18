package memoryboard

import (
	"image/color"
	"log"
	"strconv"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text"
)

type TileData struct {
	value int
	i     int
	j     int
}

type Tile struct {
	data  TileData
	color color.Color
}

var (
	mplusNormalFont font.Face
)

func init() {
	tt, err := opentype.Parse(fonts.MPlus1pRegular_ttf)
	if err != nil {
		log.Fatal(err)
	}

	const dpi = 72
	mplusNormalFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    32,
		DPI:     dpi,
		Hinting: font.HintingVertical,
	})
	if err != nil {
		log.Fatal(err)
	}
}

const (
	tileSize   = 80
	tileMargin = 4
)

var (
	tileImage = ebiten.NewImage(tileSize, tileSize)
)

func init() {
	tileImage.Fill(color.RGBA{0xff, 0x55, 0x55, 0xff})
}

func NewTile(value, i, j int) *Tile {
	return &Tile{
		data: TileData{
			value: value,
			i:     i,
			j:     j,
		},
	}
}

// func tileAt(tiles map[*Tile]struct{}, x, y int) *Tile {
// 	// for t := range tiles {

// 	// 	if x >= t.data.x && x <= t.data.x+tileSize && y >= t.data.y && y <= t.data.y+tileSize {
// 	// 		return t
// 	// 	}
// 	// }
// 	// return nil
// 	return nil
// }

func (t *Tile) Draw(boardImage *ebiten.Image) {
	i, j := t.data.i, t.data.j
	op := &ebiten.DrawImageOptions{}
	x := i*tileSize + (i+1)*tileMargin
	y := j*tileSize + (j+1)*tileMargin

	op.GeoM.Translate(float64(x), float64(y))
	op.ColorScale.ScaleWithColor(t.color)
	boardImage.DrawImage(tileImage, op)

	mplusf := mplusNormalFont

	str := strconv.Itoa(t.data.value)
	w := font.MeasureString(mplusf, str).Floor()
	h := (mplusf.Metrics().Ascent + mplusf.Metrics().Descent).Floor()
	x += (tileSize - w) / 2
	y += (tileSize-h)/2 + mplusf.Metrics().Ascent.Floor()
	text.Draw(boardImage, str, mplusf, x, y, color.Black)
}
