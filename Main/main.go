package main

import (
	//"time"

	"image/color"
	"log"

	"github.com/0xNathanW/CHIP-8/CHIP8"

	//"github.com/0xNathanW/CHIP-8/Tests"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	width   = 64
	height  = 32
	scaling = 30
)

var (
	bg = color.Black
	fg = color.White
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
	if g.emulator.DrawFlag {
		frame := ebiten.NewImage(width, height)
		for y := 0; y < len(g.emulator.Display); y++ {
			for x := 0; x < len(g.emulator.Display[y]); x++ {
				ops := &ebiten.DrawImageOptions{}
				ops.GeoM.Translate(float64(x), float64(y))
				ops.GeoM.Scale(float64(screen.Bounds().Dx())/64, float64(screen.Bounds().Dy()/32))
				var colour color.Color
				if g.emulator.Display[y][x] == 1 {
					colour = fg
				} else {
					colour = bg
				}
				frame.Fill(colour)
				screen.DrawImage(frame, ops)

			}
		}
		op := &ebiten.DrawImageOptions{}
		screen.DrawImage(frame, op)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return width * scaling, height * scaling
}

func main() {
	ebiten.SetMaxTPS(60) // Setting ticks per second
	ebiten.SetWindowSize(width*scaling, height*scaling)
	ebiten.SetWindowResizable(true)
	ebiten.SetWindowTitle("Chip8")
	//Init game, load rom etc...
	game := &Game{
		CHIP8.Initialise(),
	}
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
