package Chip8

import (
	"fmt"
	"os"
)

type Chip8 struct {
	memory [4096]byte
	V      [16]byte // 16 CPU registers
	index  uint16   // 16 bit register for addresses
	stack  [16]uint16

	PC uint16 // Program counter
	SP uint8  // Pointer to top of stack

	delayTimer byte
	soundTimer byte

	drawFlag bool // Signals whether to draw on cycle

	display Display

	keymap [16]bool
}

func Initialise() *Chip8 {
	fmt.Println("Initialising CPU...")
	// Init
	inst := &Chip8{
		PC: 0x200, // 0x000 - 0x1FF reserved for interpreter
	}
	// Load fontSet into allocated memory
	inst.LoadFontSet()

	// Return memory address of instance
	return inst
}

func (c *Chip8) Cycle() {
	// Fetch opcode
	currentOpcode := c.fetchOpcode()
	// Execute opcode
	c.executeOpcode(currentOpcode)
	// Update timers
	c.UpdateTimers()
}

func (c Chip8) fetchOpcode() uint16 {
	// Combination of opcode at program counter and next, achieved through bitwise shift
	fmt.Println("Fetching Opcode..")
	opcode := uint16(c.memory[c.PC])<<8 | uint16(c.memory[c.PC+1])
	fmt.Println(opcode)
	return opcode
}

func (c *Chip8) executeOpcode(opcode uint16) {
	fmt.Println("Executing Opcode...")

	var addr uint16 = (opcode & 0x0FFF)
	var nibble byte = uint8((opcode & 0x000F))
	var x byte = uint8((opcode & 0x0F00) >> 8)
	var y byte = uint8((opcode & 0x00F0) >> 8)
	var kk byte = uint8((opcode & 0x00FF))

	switch opcode & 0xF000 {
	case 0:
		switch nibble {
		case 0x000:
			c.CLS()
		case 0x00E:
			c.RET()
		default:
			fmt.Println("c.SYS()")
		}
	case 1:
		c.JP_NNN(addr)
	case 2:
		c.CALL_NNN(addr)
	case 3:
		c.SE_VX_NN(x, kk)
	case 4:
		c.SNE_VX_NN(x, kk)
	case 5:
		c.SE_VX_VY(x, y)
	case 6:
		c.LD_VX_NN(x, kk)
	case 7:
		c.ADD_VX_NN(x, kk)
	case 8:
		switch nibble {
		case 0:
			c.LD_VX_VY(x, y)
		case 1:
			c.OR_VX_VY(x, y)
		case 2:
			c.AND_VX_VY(x, y)
		case 3:
			c.XOR_VX_NN(x, y)
		case 4:
			c.ADD_VX_VY(x, y)
		case 5:
			c.SUB_VX_VY(x, y)
		case 6:
			c.SHR_VX(x)
		case 7:
			c.SUBN_VX_VY(x, y)
		case 0x000E:
			c.SHL_VX(x)
		default:
			fmt.Println("Fuck knows pal")
		}
	case 9:
		c.SNE_VX_VY(x, y)
	case 0xA:
		c.LD_I_NNN(addr)
	case 0xB:
		c.JP_V0_NNN(addr)
	case 0xC:
		c.RND_VX_NN(x, kk)
	case 0xD:
		c.DRW_VX_VY_N(x, y, nibble)
	case 0xE:
		switch nibble {
		case 1:
			c.SKNP_VX(x)
		case 0x00E:
			c.SKP_VX(x)
		default:
			fmt.Println("Fuck knows pal")
		}

	case 0xF:
		switch nibble {
		case 7:
			c.LD_VX_DT(x)
		case 0x00A:
			c.LD_VX_K(x)
		case 8:
			c.LD_ST_VX(x)
		case 0x00E:
			c.ADD_I_VX(x)
		case 9:
			c.LD_F_VX(x)
		case 3:
			c.BCD_VX(x)
		case 5:
			switch y {
			case 5:
				c.LD_I_VX(x)
			case 1:
				c.LD_DT_VX(x)
			case 6:
				c.LD_VX_I(x)
			default:
				fmt.Println("Fuck knows pal")
			}
		default:
			fmt.Println("Fuck knows pal")
		}
	default:
		fmt.Println("Fuck knows pal")
	}
}

func (c *Chip8) UpdateTimers() {
	fmt.Println("Updating timer...")
	if c.delayTimer > 0 {
		c.delayTimer--
	}
	if c.soundTimer > 0 {
		c.soundTimer--
	}
}

func (c *Chip8) LoadFontSet() {

	var fontSet = []byte{
		0xF0, 0x90, 0x90, 0x90, 0xF0, //0
		0x20, 0x60, 0x20, 0x20, 0x70, //1
		0xF0, 0x10, 0xF0, 0x80, 0xF0, //2
		0xF0, 0x10, 0xF0, 0x10, 0xF0, //3
		0x90, 0x90, 0xF0, 0x10, 0x10, //4
		0xF0, 0x80, 0xF0, 0x10, 0xF0, //5
		0xF0, 0x80, 0xF0, 0x90, 0xF0, //6
		0xF0, 0x10, 0x20, 0x40, 0x40, //7
		0xF0, 0x90, 0xF0, 0x90, 0xF0, //8
		0xF0, 0x90, 0xF0, 0x10, 0xF0, //9
		0xF0, 0x90, 0xF0, 0x90, 0x90, //A
		0xE0, 0x90, 0xE0, 0x90, 0xE0, //B
		0xF0, 0x80, 0x80, 0x80, 0xF0, //C
		0xE0, 0x90, 0x90, 0x90, 0xE0, //D
		0xF0, 0x80, 0xF0, 0x80, 0xF0, //E
		0xF0, 0x80, 0xF0, 0x80, 0x80, //F
	}

	for i := range fontSet {
		c.memory[i] = fontSet[i]
	}
}

func (c *Chip8) LoadROM(ROMname string) {
	fmt.Println(ROMname)

	rom, readErr := os.Open(ROMname)
	if readErr != nil {
		fmt.Println("You done fucked up.")
	}

	defer rom.Close()
}
