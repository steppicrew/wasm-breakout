package game

import "cart/w4"

const (
	BrickWidth  = 16
	BrickHeight = 8
)

type Brick struct {
	// The brick's position
	X     int
	Y     int
	Lives int
}

func (b *Brick) Draw() {
	if b.Lives <= 0 {
		return
	}
	*w4.DRAW_COLORS = uint16(0x40 + b.Lives)
	w4.Rect(b.X, b.Y, BrickWidth, BrickHeight)
}

func (b *Brick) Update() {
}
