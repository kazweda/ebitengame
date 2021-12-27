package ball

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	ballRadius = 15
)

var (
	screenW float64
	screenH float64
	posX    float64
	posY    float64
	moveX   float64
	moveY   float64
)

func Init(screenWidth int, screenHeight int) {
	screenW = float64(screenWidth)
	screenH = float64(screenHeight)
	posX = screenW / 2
	posY = screenH / 2
	moveX = float64(0.5)
	moveY = float64(0.75)
}

func UpdatePosition() {
	posX += moveX
	posY += moveY

	if posX >= screenW-ballRadius || posX <= ballRadius {
		moveX *= -1
	}

	if posY >= screenH-ballRadius || posY <= ballRadius {
		moveY *= -1
	}
}

func DrawCircle(screen *ebiten.Image) {
	purpleClr := color.RGBA{255, 0, 255, 255}

	x := int(math.Round(posX))
	y := int(math.Round(posY))
	radius64 := float64(ballRadius)
	minAngle := math.Acos(1 - 1/radius64)

	for angle := float64(0); angle <= 360; angle += minAngle {
		xDelta := radius64 * math.Cos(angle)
		yDelta := radius64 * math.Sin(angle)

		x1 := int(math.Round(float64(x) + xDelta))
		y1 := int(math.Round(float64(y) + yDelta))

		screen.Set(x1, y1, purpleClr)
	}
}
