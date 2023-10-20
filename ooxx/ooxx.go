package main

import (
	"fmt"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

const (
	screenWidth  = 300
	screenHeight = 300
	boardSize    = 3
	cellSize     = screenWidth / boardSize
)

var (
	board       [boardSize][boardSize]string
	currentTurn = "X"
	gameOver    = false
)

var mplusNormalFont font.Face

func init() {
	tt, err := opentype.Parse(fonts.MPlus1pRegular_ttf)
	if err != nil {
		log.Fatal(err)
	}

	const dpi = 72
	mplusNormalFont, _ = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    60,
		DPI:     dpi,
		Hinting: font.HintingVertical,
	})
}

type Game struct{}

func (g *Game) Update() error {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) && !gameOver {
		x, y := ebiten.CursorPosition()
		row, col := y/cellSize, x/cellSize
		if row >= 0 && row < boardSize && col >= 0 && col < boardSize && board[row][col] == "" {
			board[row][col] = currentTurn
			if checkWin(row, col) {
				fmt.Printf("%s wins!\n", currentTurn)
				gameOver = true
			} else {
				switchTurn()
			}
		}
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for i := 1; i < boardSize; i++ {
		ebitenutil.DrawLine(screen, float64(i*cellSize), 0, float64(i*cellSize), screenHeight, color.White)
		ebitenutil.DrawLine(screen, 0, float64(i*cellSize), screenWidth, float64(i*cellSize), color.White)
	}

	for row := 0; row < boardSize; row++ {
		for col := 0; col < boardSize; col++ {
			if board[row][col] == "X" {
				// ebitenutil.DrawRect(screen, float64(col*cellSize), float64(row*cellSize), cellSize, cellSize, color.White)
				text.Draw(screen, "X", mplusNormalFont, col*cellSize+cellSize/2-5, row*cellSize+cellSize/2, color.White)
			} else if board[row][col] == "O" {
				// ebitenutil.DrawRect(screen, float64(col*cellSize), float64(row*cellSize), cellSize, cellSize, color.White)
				text.Draw(screen, "O", mplusNormalFont, col*cellSize+cellSize/2, row*cellSize+cellSize/2, color.White)
			}
		}
	}

	ebitenutil.DebugPrint(screen, fmt.Sprintf("Current Turn: %s", currentTurn))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 300, 300
}

func main() {
	for row := 0; row < boardSize; row++ {
		for col := 0; col < boardSize; col++ {
			board[row][col] = ""
		}
	}

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Tic-Tac-Toe")

	game := &Game{}
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}

func switchTurn() {
	if currentTurn == "X" {
		currentTurn = "O"
	} else {
		currentTurn = "X"
	}
}

func checkWin(row, col int) bool {
	player := board[row][col]

	// 检查每一行
	for i := 0; i < 3; i++ {
		if board[i][0] == player && board[i][1] == player && board[i][2] == player {
			return true
		}
	}

	// 检查每一列
	for i := 0; i < 3; i++ {
		if board[0][i] == player && board[1][i] == player && board[2][i] == player {
			return true
		}
	}

	// 检查对角线
	if board[0][0] == player && board[1][1] == player && board[2][2] == player {
		return true
	}

	// 检查反对角线
	if board[0][2] == player && board[1][1] == player && board[2][0] == player {
		return true
	}

	return false
}
