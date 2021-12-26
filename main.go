package main

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kazweda/ebitengame/circle"
)

const (
	screenWidth  = 300
	screenHeight = 300

	ballRadius = 15
)

type Game struct{}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

var (
	ballPositionX = float64(screenWidth) / 2
	ballPositionY = float64(screenHeight) / 2
	ballMovementX = float64(0.5)
	ballMovementY = float64(0.75)
)

func (g *Game) Update() error {
	ballPositionX += ballMovementX
	ballPositionY += ballMovementY

	if ballPositionX >= screenWidth-ballRadius || ballPositionX <= ballRadius {
		ballMovementX *= -1
	}

	if ballPositionY >= screenHeight-ballRadius || ballPositionY <= ballRadius {
		ballMovementY *= -1
	}

	return nil
}

func (g *Game) drawCircle(screen *ebiten.Image, x, y, radius int, clr color.Color) {
	circle.DrawCircle(screen, x, y, radius, clr)
}

func (g *Game) Draw(screen *ebiten.Image) {
	purpleClr := color.RGBA{255, 0, 255, 255}

	x := int(math.Round(ballPositionX))
	y := int(math.Round(ballPositionY))

	g.drawCircle(screen, x, y, ballRadius, purpleClr)
}

func main() {
	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("The Moving Ball")

	if err := ebiten.RunGame(&Game{}); err != nil {
		panic(err)
	}
}
