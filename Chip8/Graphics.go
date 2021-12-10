package CHIP8

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
)

const (
//  Foreground and background colours
)

type Display struct {
	Screen     	tcell.Screen
	PixelArray	[64][32]int
	DrawFlag   	bool
	Scale 		int
}

func NewDisplay() *Display {
	screen, err := tcell.NewScreen()
	if err != nil {
		panic(fmt.Error("failed to initisialise screen ", err))
	}

	display := &Display{
		Screen:     screen,
		PixelArray: [64][32]{},
		DrawFlag:   false,
	}
	return display
}

func (c *Chip8) Draw() {
	for y := 0; y < len(c.PixelArray); y++ {
		for x := 0; x < len(c.PixelArray[y]); x++ {
			if c.PixelArray[y][x] == 0 {

			} else {

			}
		}
	}

}
