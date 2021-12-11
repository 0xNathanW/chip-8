package main

import (
	"github.com/0xNathanW/chip-8/chip8"
)

func main() {
	emulator := chip8.NewSystem()
	emulator.Run()
}
