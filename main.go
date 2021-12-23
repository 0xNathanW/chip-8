package main

import (
	"flag"
	"github.com/0xNathanW/chip-8/chip8"
)

func main() {
	flag.StringVar(&chip8.Bg, "bg", "black", "Background colour")
	flag.StringVar(&chip8.Fg, "fg", "white", "Foreground colour")
	flag.Parse()
	emulator := chip8.NewSystem()
	emulator.Run()
}
