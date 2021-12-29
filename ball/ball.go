package ball

import (
	"image/color"
	"math"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	ballRadius = 15
)

var (
	screenW        float64
	screenH        float64
	posX           float64
	posY           float64
	moveX          float64
	moveY          float64
	prevUpdateTime time.Time
)

func Init(screenWidth int, screenHeight int) {
	screenW = float64(screenWidth)
	screenH = float64(screenHeight)
	posX = screenW / 2
	posY = screenH / 2
	moveX = float64(0.00000006)
	moveY = float64(0.00000004)
	prevUpdateTime = time.Now()
}

func UpdatePosition() {
	timeDelta := float64(time.Since(prevUpdateTime))
	prevUpdateTime = time.Now()

	posX += moveX * timeDelta
	posY += moveY * timeDelta

	const minX = ballRadius
	const minY = ballRadius
	var maxX = screenW - ballRadius
	var maxY = screenH - ballRadius

	if posX >= maxX || posX <= minX {
		if posX > maxX {
			posX = maxX
		} else if posX < minX {
			posX = minX
		}

		moveX *= -1
	}

	if posY >= maxY || posY <= minY {
		if posY > maxY {
			posY = maxY
		} else if posY < minY {
			posY = minY
		}

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
