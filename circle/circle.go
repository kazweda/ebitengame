package circle

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

func DrawCircle(screen *ebiten.Image, ballX float64, ballY float64, radius int) {
	purpleClr := color.RGBA{255, 0, 255, 255}

	x := int(math.Round(ballX))
	y := int(math.Round(ballY))
	radius64 := float64(radius)
	minAngle := math.Acos(1 - 1/radius64)

	for angle := float64(0); angle <= 360; angle += minAngle {
		xDelta := radius64 * math.Cos(angle)
		yDelta := radius64 * math.Sin(angle)

		x1 := int(math.Round(float64(x) + xDelta))
		y1 := int(math.Round(float64(y) + yDelta))

		screen.Set(x1, y1, purpleClr)
	}
}
