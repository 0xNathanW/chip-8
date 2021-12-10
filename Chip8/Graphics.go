package chip8

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
)

const (
//  Foreground and background colours
)

type Display struct {
	Screen     tcell.Screen
	PixelArray [64][32]int
	DrawFlag   bool
	Scale      int
}

func NewDisplay() *Display {
	screen, err := tcell.NewScreen()
	if err != nil {
		panic(fmt.Errorf("failed to initisialise screen %w", er))
	}

	display := &Display{
		Screen:     screen,
		PixelArray: [64][32]int{},
		DrawFlag:   false,
	}
	return display
}
