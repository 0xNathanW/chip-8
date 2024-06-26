package main

import (
	"flag"

	chip8 "github.com/0xNathanW/chip-8/Chip8"
)

func main() {
	// flags determine colours.
	flag.StringVar(&chip8.Bg, "bg", "black", "Background colour")
	flag.StringVar(&chip8.Fg, "fg", "white", "Foreground colour")
	flag.Parse()
	emulator := chip8.NewSystem()
	emulator.Run()
}
