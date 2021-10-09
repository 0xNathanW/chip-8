package CHIP8

import (
	"github.com/nsf/termbox-go"
)

const (
	//  Foreground and background colours
	bg = termbox.ColorDefault
	fg = termbox.ColorBlue
)

type Display struct {
	fg, bg termbox.Attribute
}

func InitDisplay() *Display {
	return &Display{fg, bg}
}

func (d *Display) Draw(c *Chip8) {
	for y := 0; y < len(c.PixelArray); y++ {
		for x := 0; x < len(c.PixelArray[y]); x++ {
			var pixColour termbox.Attribute
			if c.PixelArray[y][x] == 0 {
				pixColour = bg
			} else {
				pixColour = fg
			}
			termbox.SetCell(x*2, y, rune(' '), pixColour, pixColour)
			termbox.SetCell(x*2+1, y, rune(' '), pixColour, pixColour)
		}
	}
	termbox.Flush()
}
