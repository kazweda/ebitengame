package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth  = 300
	screenHeight = 300
)

type Game struct{}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	purpleCol := color.RGBA{255, 0, 255, 255}

	for x := 100; x < 200; x++ {
		for y := 100; y < 200; y++ {
			screen.Set(x, y, purpleCol)
		}
	}
}

func main() {
	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("Ebiten Test")

	if err := ebiten.RunGame(&Game{}); err != nil {
		panic(err)
	}
}
