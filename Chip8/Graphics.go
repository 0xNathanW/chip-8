package CHIP8

import (
	"github.com/nsf/termbox-go"
)

const (
	//  Foreground and background colours
	fg = termbox.ColorWhite
	bg = termbox.ColorBlack
)

type Screen struct {
	fg, bg termbox.Attribute
}

func TermboxSetup() error {
	err := termbox.Init()
	if err != nil {
		return err
	}
	defer termbox.Close()
	termbox.SetOutputMode(termbox.OutputNormal)
	termbox.HideCursor()
	err1 := termbox.Clear(bg, bg)
	if err1 != nil {
		return err1
	}
	return termbox.Flush()
}

func InitDisplay() *Screen {
	return &Screen{fg, bg}
}

func (c *Chip8) Draw() {
	for y := 0; y < len(c.PixelArray); y++ {
		for x := 0; x < len(c.PixelArray[y]); x++ {
			if c.PixelArray[y][x] == 0 {
				termbox.SetCell(x, y, rune(' '), c.Display.fg, c.Display.bg)
			} else {
				termbox.SetCell(x, y, rune(' '), c.Display.fg, c.Display.bg)
			}
		}
	}
	termbox.Flush()
}
