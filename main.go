package main

import (
	"github.com/0xNathanW/chip-8/app"
)

func main() {
	emulator := app.NewApp()
	emulator.Run()
}
