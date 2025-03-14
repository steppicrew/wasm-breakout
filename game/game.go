package game

import "math"

const (
	BricksPerLine = 10
	BrickLines    = 5
)

var bricks = [BricksPerLine * BrickLines]Brick{}

var board = Board{}

var ball = Ball{
	X: 80,
	Y: 100,
	Speed: Speed{
		X:     1,
		Y:     1,
		Value: math.Sqrt2,
	},
}

var bouncer = Bouncer{
	X: 72,
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

/*
var ball = Ball{
	X: 20,
	Y: BoardOffset + 57,
	Speed: Speed{
		X: 0,
		Y: -1,
	},
}
*/

func Update() {
	board.Update()
	for _, b := range bricks {
		b.Update()
	}
	bouncer.Update()
	ball.Update()
	ball.CeckCollision(bricks[:], bouncer)
}

func Draw() {
	board.Draw()
	for _, b := range bricks {
		b.Draw()
	}
	bouncer.Draw()
	ball.Draw()
}

func Initialize() {
	board.Initialize()
	ball.Initialize()
	bouncer.Initialize()

	for y := range BrickLines {
		for x := range BricksPerLine {
			bricks[y*10+x] = Brick{
				X:     x * BrickWidth,
				Y:     BoardOffset + y*BrickHeight + 16,
				Lives: Max(4-y, 1),
			}
		}
	}
}
