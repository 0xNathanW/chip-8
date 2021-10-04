package Tests

import (
	"fmt"

	"github.com/0xNathanW/CHIP-8/CHIP8"
)

func Test() *CHIP8.Chip8 {
	inst := CHIP8.Initialise()
	// fmt.Println("------- Memory Before Load--------")
	// fmt.Println(inst.Memory)
	// fmt.Println("-----------------------------------")
	inst.LoadROM()
	// fmt.Println("------- Memory After Load--------")
	// fmt.Println(inst.Memory)
	// fmt.Println("-----------------------------------")

	for {
		fmt.Println("----------- New Cycle ----------------")
		inst.Cycle()
		fmt.Println("Program Counter: ", inst.PC)
		fmt.Println("Registers: ", inst.V)
		fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>> GFX <<<<<<<<<<<<<<<<<<<<<<<<<<<<<")
		TestDisplayOutput(inst)
	}
}

func TestDisplayOutput(c *CHIP8.Chip8) {
	for y := 0; y < int(len(c.Display)); y++ {
		for x := 0; x < int(len(c.Display[0])); x++ {
			pixel := c.Display[y][x]
			if pixel == 1 {
				fmt.Print("██")
			} else {
				fmt.Print("--")
			}
		}
		fmt.Print("\n")
	}
}
