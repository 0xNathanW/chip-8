package chip8

import (
	"fmt"
	"time"

	"github.com/gdamore/tcell/v2"
)

var (
	//  Foreground and background colours.
	Bg string
	Fg string
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

	background := tcell.GetColor(Bg)
	foreground := tcell.GetColor(Fg)

	display := &display{
		screen:     screen,
		pixelArray: [64][32]int{},
		scale:      2,
		style:      tcell.StyleDefault.
						Foreground(foreground).
						Background(background),
	}
	return display
}

// Draws pixel array to screen.
func (d *display) draw() {
	for y := 0; y < 32; y++ {
		for x := 0; x < 64; x++ {
			var char rune
			if d.pixelArray[x][y] == 1 {
				char = '█'
			} else {
				char = ' '
			}
			d.screen.SetContent(x*d.scale, y*(d.scale-1), char, nil, d.style)
			d.screen.SetContent(x*d.scale+1, y*(d.scale-1), char, nil, d.style)
		}
	}
}

// Draw a singular line of text to screen.
func (d *display) drawLine(x int, y int, text string, highlight bool) {
	if highlight {
		for i:=0; i < len(text); i++ {
			d.screen.SetContent(x+i, y, rune(text[i]), nil, d.style.Reverse(true))
		}
	} else {
		for i:=0; i < len(text); i++ {
			d.screen.SetContent(x+i, y, rune(text[i]), nil, d.style)
		}
	}
}

// Fills 128x32 pixel array with background colour.
func (d *display) fill() {
	for y := 0; y < 32; y++ {
		for x := 0; x < 128; x++ {
			d.screen.SetContent(x, y, ' ', nil, d.style)
		}
	}
}