package main

import ( //"github.com/0xNathanW/CHIP-8/CHIP8"
	//"github.com/0xNathanW/CHIP-8/Tests"
	"github.com/0xNathanW/CHIP-8/CHIP8"
	//"github.com/nsf/termbox-go"
	"time"
)

func main() {
	CHIP8.TermboxSetup()
	emulator := CHIP8.Initialise()

	clock := time.Tick(emulator.ClockSpeed)

	for {
		select {
		case <-clock:
			emulator.Cycle()
		}
	}
}
