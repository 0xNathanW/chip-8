package Tests

import (
	"fmt"
	//"time"

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
		inst.Cycle()
		fmt.Println(inst.PC)
		fmt.Println(inst.V)
		// sleepTime := time.Second
		// time.Sleep(sleepTime)
	}
}
