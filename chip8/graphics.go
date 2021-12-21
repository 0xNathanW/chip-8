package chip8

import (
	"fmt"
	"time"

	"github.com/gdamore/tcell/v2"
)

const (
	//  Foreground and background colours.
	bg = tcell.ColorBlack
	fg = tcell.ColorWhite
	// Screen refreshes per second.
	refreshRate = time.Second / 60
)

type Display struct {
	Screen     tcell.Screen
	PixelArray [64][32]int
	DrawFlag   bool
	Scale      int
	Style      tcell.Style
}

func NewDisplay() *Display {
	screen, err := tcell.NewScreen()
	if err != nil {
		panic(fmt.Errorf("failed to initisialise screen %w", err))
	}

	display := &Display{
		Screen:     screen,
		PixelArray: [64][32]int{},
		Scale:      1,
		Style:      tcell.StyleDefault.Foreground(fg).Background(bg),
	}
	return display
}
