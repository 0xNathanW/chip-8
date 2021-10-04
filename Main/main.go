package main

import (
	//"time"

	"image/color"

	"github.com/0xNathanW/CHIP-8/CHIP8"
	"github.com/0xNathanW/CHIP-8/Tests"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	width   = 640
	height  = 320
	scaling = 10
)

type Game struct {
	emulator *CHIP8.Chip8
}

func (g *Game) Update() error {
	g.emulator.PressedKeys()
	g.emulator.Cycle()
	g.emulator.ResetKeys()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	window := ebiten.NewImage(width, height)
	window.Fill(color.Black)
	img := ebiten.NewImage(width/2, height/2)
	img.Fill(color.White)
	window.DrawImage(img, op)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return width * scaling, height * scaling
}

func main() {
	// ebiten.SetMaxTPS(60)
	// ebiten.SetWindowSize(width*scaling, height*scaling)
	// ebiten.SetWindowResizable(true)
	// ebiten.SetWindowTitle("Chip8")
	// //Init game, load rom etc...
	// game := &Game{
	// 	CHIP8.Initialise(),
	// }
	// if err := ebiten.RunGame(game); err != nil {
	// 	log.Fatal(err)
	// }
	Tests.Test()
}
