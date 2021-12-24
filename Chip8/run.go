package chip8

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gdamore/tcell/v2"
)

func (c *Chip8) Run() {
	err := c.display.screen.Init()
	if err != nil {
		panic(fmt.Errorf("error initializing screen: %v", err))
	}
	c.display.screen.HideCursor()
	defer c.display.screen.Fini()

	c.chooseROM()

	clock := time.NewTicker(c.ClockSpeed)      // the cpu clock.
	screen := time.NewTicker(refreshRate)      // the screen refresh.
	timers := time.NewTicker(time.Second / 60) // timers run at a set 60Hz rate.

	eventQ := make(chan tcell.Event) // q for tcell events.

	// goroutine polls for events.
	go func() {
		for {
			eventQ <- c.display.screen.PollEvent()
		}
	}()

	for {
		select {
		case <-clock.C:
			c.Cycle()

		case event := <-eventQ:
			if key, ok := event.(*tcell.EventKey); ok {
				// chip8 keypad press.
				if k, ok := keyMap[key.Rune()]; ok {
					c.keypad[k] = true
				}
				// p for pause.
				if key.Rune() == 'p' {
					c.isPaused = true
					c.paused()
				}
			}

		case <-timers.C:
			c.UpdateTimers()

		case <-screen.C:
			c.display.screen.Show()
		}
	}

}

// loop that runs when emulator is in paused state.
func (c *Chip8) paused() {

	eventQ := make(chan tcell.Event)

	go func() {
		for {
			eventQ <- c.display.screen.PollEvent()
		}
	}()

	for event := range eventQ {
		if key, ok := event.(*tcell.EventKey); ok {
			// escape key quits program
			if key.Key() == tcell.KeyEsc {
				c.display.screen.Clear()
				c.display.screen.Fini()
				os.Exit(0)
			}
			// pressing p again unpauses the emulator.
			if key.Rune() == 'p' {
				c.isPaused = false
				return
			}
			// r key restarts the emulator.
			// allowing a new ROM to be loaded.
			if key.Rune() == 'r' {
				c.reset()
				c.display.screen.Clear()
				c.chooseROM()
				return
			}
			// right arrow key steps forward one cycle..
			if key.Key() == tcell.KeyRight {
				c.Cycle()
				c.display.screen.Show()
			}

		}
		if _, ok := event.(*tcell.EventResize); ok {
			c.display.screen.Sync()
		}

	}
}

func (c *Chip8) chooseROM() {

	c.display.fill()

	offset := 10

	txt := "Select a ROM to run:"
	div := "---------------------"
	c.display.drawLine(offset, 5, txt, false)
	c.display.drawLine(offset, 6, div, false)

	roms := make(map[int]string)

	files, err := os.ReadDir("./ROMs")
	if err != nil {
		log.Fatal("ROM folder does not exist")
	}

	// add ROM's to display.
	for i, file := range files {
		roms[i] = file.Name()
		if i == 0 {
			c.display.drawLine(offset, i+7, file.Name(), true)
		} else {
			c.display.drawLine(offset, i+7, file.Name(), false)
		}
	}

	currentIdx := 0 // tracks current selected ROM.
	c.display.screen.Show()
	eventQ := make(chan tcell.Event)

	go func() {
		for {
			eventQ <- c.display.screen.PollEvent()
		}
	}()

	for event := range eventQ {
		if key, ok := event.(*tcell.EventKey); ok {
			if key.Key() == tcell.KeyUp {
				if currentIdx > 0 {
					c.display.drawLine(offset, currentIdx+7, roms[currentIdx], false)
					currentIdx--
					c.display.drawLine(offset, currentIdx+7, roms[currentIdx], true)
					c.display.screen.Show()
				}
			}

			if key.Key() == tcell.KeyDown {
				if currentIdx < len(roms)-1 {
					c.display.drawLine(offset, currentIdx+7, roms[currentIdx], false)
					currentIdx++
					c.display.drawLine(offset, currentIdx+7, roms[currentIdx], true)
					c.display.screen.Show()
				}
			}

			if key.Key() == tcell.KeyEnter {
				c.LoadROM(roms[currentIdx])
				c.display.screen.Clear()
				c.display.screen.Show()
				return
			}

			if key.Key() == tcell.KeyEsc {
				c.display.screen.Fini()
				c.display.screen.Clear()
				os.Exit(0)
			}

		}

	}
}
