package memoryboard

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type mouseState int

const (
	mouseStateNone mouseState = iota
	mouseStatePressing
	mouseStateSettled
)

type Input struct {
	mouseState    mouseState
	mouseInitPosX int
	mouseInitPosY int
}

func NewInput() *Input {
	return &Input{}
}

func (i *Input) Update() {
	switch i.mouseState {
	case mouseStateNone:
		if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
			i.mouseInitPosX, i.mouseInitPosY = ebiten.CursorPosition()
			i.mouseState = mouseStatePressing
		}
	case mouseStatePressing:
		if !ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
			i.mouseState = mouseStateSettled
		}
	case mouseStateSettled:
		i.mouseState = mouseStateNone
	}
}

func (i *Input) ClickDown() (int, int, bool) {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		return i.mouseInitPosX, i.mouseInitPosY, true
	}
	return 0, 0, false
}

func (i *Input) ClickUp() (int, int, bool) {
	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		return i.mouseInitPosX, i.mouseInitPosY, true
	}
	return 0, 0, false
}
