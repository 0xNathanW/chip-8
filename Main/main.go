package main

import (
	//"fmt"
	"time"

	"github.com/0xNathanW/CHIP-8/Chip8"
)

func main() {
	for {
		Chip8.GetKeyPress()
		time.Sleep(1 * time.Second)
	}
	//  Initialising
	// fmt.Println("Initialising...")
	// emulator := Chip8.Initialise()
	// emulator.LoadROM(name)

	// //	Main loop
	// for {
	// 	emulator.Cycle()

	// 	if Chip8.DrawFlag() {
	// 		drawGFX()
	// 	}

	// 	Chip8.setKeys()
	// }
	// fmt.Println(Chip8.Test())
}
