package game

import "cart/w4"

var ballSprite = [6]byte{
	0b00110000,
	0b01111000,
	0b11111100,
	0b11111100,
	0b01111000,
	0b00110000,
}

const (
	BallSize = 6
)

type Speed struct {
	X int
	Y int
}

type Ball struct {
	X     int
	Y     int
	Speed Speed
}

func (b *Ball) Draw() {
	*w4.DRAW_COLORS = 0x42
	w4.Oval(b.X, b.Y, BallSize, BallSize)
}

func (b *Ball) Update() {
	b.X += b.Speed.X
	b.Y += b.Speed.Y
}

func (b *Ball) CeckCollision(bricks []Brick) {
	if b.X <= 0 {
		b.BounceEW()
	}
	if b.X >= BoardWidth-BallSize {
		b.BounceEW()
	}
	if b.Y <= BoardOffset {
		b.Y = BoardOffset
		b.BounceNS()
	}
	if b.Y >= TotalHeight-BallSize {
		b.BounceNS()
	}
	bounced := false
	for i, brick := range bricks {
		if brick.Lives <= 0 {
			continue
		}
		if b.X+BallSize > brick.X && b.X < brick.X+BrickWidth && b.Y+BallSize >= brick.Y && b.Y <= brick.Y+BrickHeight {
			for _y := range BrickHeight {
				y := brick.Y - ball.Y + _y
				if y < 0 || y >= BallSize {
					continue
				}
				for _x := range BrickWidth {
					x := brick.X - ball.X + _x
					if x < 0 || x >= BallSize {
						continue
					}
					if ballSprite[y]&(1<<uint(x)) != 0 {
						if !bounced {
							if x > 0 && x < BallSize-1 {
								b.BounceNS()
							}
							if y > 0 && y < BallSize-1 {
								b.BounceEW()
							}
							bounced = true
						}
						bricks[i].Lives--
					}
				}
			}
		}
	}
}

func (b *Ball) BounceEW() {
	b.Speed.X = -b.Speed.X
}
func (b *Ball) BounceNS() {
	b.Speed.Y = -b.Speed.Y
}
