package Chip8

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	displayWidth  = 64
	displayHeight = 32
)

type Display struct {
	pixels [displayWidth][displayHeight]bool // Graphics display, 64 by 32 pixels
}

func newDisplay() Display {
	display := &Display{}
	return *display
}

func (d *Display) Frame(screen *ebiten.Image) {

	var newFrame = ebiten.NewImage(displayWidth, displayHeight)

	for x := range d.pixels[0] {
		for y := range d.pixels[1] {

			var pixelColour color.Gray16

			if d.pixels[x][y] {
				pixelColour = color.White
			} else {
				pixelColour = color.Black
			}

			newFrame.Set(x, y, pixelColour)

		}
	}

}
