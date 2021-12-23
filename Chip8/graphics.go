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
	//display.setScale(screen.Size())
	return display
}

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

// func (d *display) setScale(w, h int) {
// 	a, b := w/128, h/32
// 	if a < b {
// 		d.scale = a
// 	} else {
// 		d.scale = b
// 	}
// }