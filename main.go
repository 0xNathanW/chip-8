package main

import (
	"os"
	"time"

	"github.com/0xNathanW/CHIP-8/CHIP8"
	"github.com/nsf/termbox-go"
)

func main() {
	// Initalisation
	emulator := CHIP8.Initialise()
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()
	termbox.SetOutputMode(termbox.OutputNormal)
	termbox.HideCursor()
	display := CHIP8.InitDisplay()
	// Clock is a channel that is sent ticks
	clock := time.Tick(emulator.ClockSpeed)

	// Channel for keypresses
	eventQ := make(chan termbox.Event)
	// Goroutine to recieve keypress events
	go func() {
		for {
			eventQ <- termbox.PollEvent()
		}
	}()

	// Mainloop
	for {
		select {
		case event := <-eventQ:
			if event.Type == termbox.EventKey {
				if event.Key == termbox.KeyEsc {
					os.Exit(0)
				}

				if k, ok := CHIP8.Keymap[event.Ch]; ok {
					emulator.Keypad[k] = true
				}
			}
		case <-clock:
			emulator.Cycle()
			if emulator.DrawFlag {
				display.Draw(emulator)
				emulator.DrawFlag = false
			}
		}
	}
}
