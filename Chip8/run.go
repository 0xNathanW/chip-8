package chip8

import (
	"fmt"
	"os"
	"time"

	"github.com/gdamore/tcell/v2"
)

func (c *Chip8) Run() {

	err := c.Display.Screen.Init()
	if err != nil {
		panic(fmt.Errorf("error initializing screen: %v", err))
	}
	c.Display.Screen.HideCursor()
	defer c.Display.Screen.Fini()

	clock := time.NewTicker(c.ClockSpeed)
	screen := time.NewTicker(refreshRate)
	// Timers run at a set 60Hz rate.
	timers := time.NewTicker(time.Second / 60)

	eventQ := make(chan tcell.Event)

	go func() {
		for {
			eventQ <- c.Display.Screen.PollEvent()
		}
	}()

	fmt.Println("Running...")
	for {
		select {
		case event := <-eventQ:
			if key, ok := event.(*tcell.EventKey); ok {

				if key.Key() == tcell.KeyEsc {
					os.Exit(0)
				}

				if k, ok := KeyMap[key.Rune()]; ok {
					// Debug: is printing correctly.
					c.Keypad[k] = true
				}
			}

		case <-clock.C:
			c.Cycle()

		case <-timers.C:
			c.UpdateTimers()

		case <-screen.C:
			c.Display.Screen.Show()
		}
	}

}
