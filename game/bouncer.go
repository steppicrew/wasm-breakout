package game

import "cart/w4"

const (
	BouncerWidth  = 32
	BouncerHeight = 8
	BouncerY      = 140
)

type Bouncer struct {
	X int
}

func (b *Bouncer) Draw() {
	*w4.DRAW_COLORS = 0x22
	w4.Rect(b.X, BouncerY, BouncerWidth, BouncerHeight)
}

func (b *Bouncer) Update() {
	var gamepad = *w4.GAMEPAD1

	// Only the buttons that were pressed down this frame
	var pressedThisFrame = gamepad

	if pressedThisFrame&w4.BUTTON_RIGHT != 0 {
		b.X = Min(b.X+1, BoardWidth-BouncerWidth)
	}
	if pressedThisFrame&w4.BUTTON_LEFT != 0 {
		b.X = Max(b.X-1, 0)
	}
}

func (b *Bouncer) Initialize() {
}

func (b *Bouncer) Border() (int, int, int, int) {
	return b.X, BouncerY, b.X + BouncerWidth - 1, BouncerY + BouncerHeight - 1
}
