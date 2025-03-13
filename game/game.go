package game

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
		X: 1,
		Y: 1,
	},
}

func Update() {
	board.Update()
	for _, b := range bricks {
		b.Update()
	}
	ball.Update()
	ball.CeckCollision(bricks[:])
}

func Draw() {
	board.Draw()
	for _, b := range bricks {
		b.Draw()
	}
	ball.Draw()
}

func Initialize() {
	board.Initialize()

	for y := range BrickLines {
		for x := range BricksPerLine {
			bricks[y*10+x] = Brick{
				X:     x * BrickWidth,
				Y:     BoardOffset + y*BrickHeight + 16,
				Lives: 4 - y,
			}
		}
	}
}
