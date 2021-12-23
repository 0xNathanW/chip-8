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

type display struct {
	screen     tcell.Screen
	pixelArray [64][32]int
	scale      int
	style      tcell.Style
}

func newDisplay() *display {
	screen, err := tcell.NewScreen()
	if err != nil {
		panic(fmt.Errorf("failed to initisialise screen %w", err))
	}

	display := &display{
		screen:     screen,
		pixelArray: [64][32]int{},
		style:      tcell.StyleDefault.Foreground(fg).Background(bg),
	}
	display.setScale(screen.Size())
	return display
}

func (d *display) setScale(w, h int) {
	a, b := w/128, h/32
	if a < b {
		d.scale = a
	} else {
		d.scale = b
	}
}