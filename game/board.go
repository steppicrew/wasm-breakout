package game

import "cart/w4"

const (
	TotalWidth  = 160
	TotalHeight = 160

	BoardWidth  = TotalHeight
	BoardHeight = 144

	BoardOffset = TotalHeight - BoardHeight
)

type Board struct {
}

func (b *Board) Draw() {
	*w4.DRAW_COLORS = 0x41
	w4.Rect(0, BoardOffset, BoardWidth, BoardHeight)
}

func (b *Board) Initialize() {
}

func (b *Board) Update() {
}
