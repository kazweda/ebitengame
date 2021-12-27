package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kazweda/ebitengame/ball"
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
	ball.UpdatePosition()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ball.DrawCircle(screen)
}

func init() {
	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("The Moving Ball")
	ball.Init(screenWidth, screenHeight)
}

func main() {
	if err := ebiten.RunGame(&Game{}); err != nil {
		panic(err)
	}
}
