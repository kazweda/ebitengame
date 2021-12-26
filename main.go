package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kazweda/ebitengame/circle"
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

func (g *Game) drawCircle(screen *ebiten.Image, x, y, radius int, clr color.Color) {
	circle.DrawCircle(screen, x, y, radius, clr)
}

func (g *Game) Draw(screen *ebiten.Image) {
	purpleClr := color.RGBA{255, 0, 255, 255}

	g.drawCircle(screen, 150, 150, 50, purpleClr)
}

func main() {
	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("Ebiten Test")

	if err := ebiten.RunGame(&Game{}); err != nil {
		panic(err)
	}
}
