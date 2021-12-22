package app

import (
	"fmt"
	//"github.com/rivo/tview"
)

func (a *App) Run() {
	
	
	if err := a.App.SetRoot(a.Layout, true).SetFocus(a.SelectROM).Run(); err != nil {
		panic(err)
	}
	fmt.Println("Done...")
}


func (a *App) emulate() {
	a.App.SetFocus(a.Display)
	chipGFX := make(chan [64][32]int)

	go a.Chip8.Run(chipGFX)

	go func() {
		
	}

	for {
		select {
		case gfx := <-chipGFX:
			a.drawToDisplay(gfx)
		}
	}

}

func (a *App) drawToDisplay(px [64][32]int) {
	var txt string
	for y := 0; y < 32; y++ {
		for x := 0; x < 64; x++ {
			if px[x][y] == 1 {
				txt += "█"
			} else {
				txt += " "
			}
		}
		txt += "\n"
	}
	a.Display.SetText(txt)
}

