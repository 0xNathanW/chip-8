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

	go func() {
		for {
			<-screen.C
			c.Display.Screen.Show()
		}
	}()

	go func() {
		for {
			<-timers.C
			c.UpdateTimers()
		}
	}()

	fmt.Println("Running...")
	for {
		select {
		case <-clock.C:
			c.Cycle()

		case event := <-eventQ:
			if key, ok := event.(*tcell.EventKey); ok {

				if key.Key() == tcell.KeyEsc {
					os.Exit(0)
				}

				if k, ok := KeyMap[key.Rune()]; ok {
					fmt.Println(key.Name())
					c.keypad[k] = true
				}
			}
		}

	}

}
