package game

import (
	"cart/w4"
	"math"
)

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
	X     float64
	Y     float64
	Value float64
}

type Ball struct {
	X     float64
	Y     float64
	iX    int
	iY    int
	Speed Speed
}

func (b *Ball) Draw() {
	*w4.DRAW_COLORS = 0x42
	w4.Oval(b.iX, b.iY, BallSize, BallSize)
}

func (b *Ball) Update() {
	b.X += b.Speed.X
	b.Y += b.Speed.Y
	b.iX = int(b.X)
	b.iY = int(b.Y)
}

func (b *Ball) Initialize() {
	b.normalizeSpeed()
}

func (b *Ball) CeckCollision(bricks []Brick, bouncer Bouncer) {
	bouncerX0, bouncerY0, bouncerX1, _ := bouncer.Border()
	if b.iY+BallSize > bouncerY0 && b.iX+BallSize > bouncerX0 && b.iX < bouncerX1 {
		b.BounceUp()
		ballCenter := b.iX + BallSize/2
		bouncerCenter := bouncer.X + BouncerWidth/2
		switch {
		case ballCenter < bouncerCenter-BouncerWidth/4:
			b.Speed.X -= .5
		case ballCenter > bouncerCenter+BouncerWidth/4:
			b.Speed.X += .5
		}
	}

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

	for i, brick := range bricks {
		if brick.Lives <= 0 {
			continue
		}
		if b.iX+BallSize > brick.X && b.iX < brick.X+BrickWidth && b.iY+BallSize >= brick.Y && b.iY <= brick.Y+BrickHeight {
			mask := byte(0b11111100)
			if brick.X > b.iX {
				diff := uint(brick.X - b.iX)
				mask = (mask << diff) >> diff
			}
			if brick.X+BrickWidth < b.iX+BallSize {
				diff := uint(b.iX + BallSize - (brick.X + BrickWidth))
				mask = (mask >> diff) << diff
			}
			bouncedNS := false
			bouncedEW := false
			for _y := range BrickHeight {
				y := brick.Y - b.iY + _y
				if y < 0 || y >= BallSize {
					continue
				}
				overlap := ballSprite[y] & mask
				if overlap != 0 {
					w4.Blit(&overlap, 0, _y, 6, 1, w4.BLIT_1BPP)
					if !bouncedEW {
						if overlap&0b11000000 > 0 {
							b.BounceRight()
							bouncedEW = true
						}
						if overlap&0b00001100 > 0 {
							b.BounceLeft()
							bouncedEW = true
						}
					}
					if !bouncedNS {
						bouncedNS = true
						if y < BallSize/2 {
							b.BounceDown()
						} else {
							b.BounceUp()
						}
					}
					bricks[i].Lives--
				}
			}
		}
	}
	b.normalizeSpeed()
}

func (b *Ball) normalizeSpeed() {
	abs := math.Sqrt(b.Speed.X*b.Speed.X + b.Speed.Y*b.Speed.Y)
	if b.Speed.Y > abs {
		abs = b.Speed.Y
	}
	b.Speed.X /= abs / b.Speed.Value
	b.Speed.Y /= abs / b.Speed.Value
}

func (b *Ball) BounceEW() {
	b.Speed.X = -b.Speed.X
}
func (b *Ball) BounceNS() {
	b.Speed.Y = -b.Speed.Y
}
func (b *Ball) BounceLeft() {
	if b.Speed.X > 0 {
		b.Speed.X = -b.Speed.X
	}
}
func (b *Ball) BounceRight() {
	if b.Speed.X < 0 {
		b.Speed.X = -b.Speed.X
	}
}
func (b *Ball) BounceUp() {
	if b.Speed.Y > 0 {
		b.Speed.Y = -b.Speed.Y
	}
}
func (b *Ball) BounceDown() {
	if b.Speed.Y < 0 {
		b.Speed.Y = -b.Speed.Y
	}
}
