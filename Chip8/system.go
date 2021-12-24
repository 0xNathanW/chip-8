package chip8

import (
	"fmt"
	"log"
	"os"
	"time"
)

const clockSpeed = (time.Second / 500)

type Chip8 struct {
	//=====  CPU  =====//
	memory [4096]byte
	V      [16]byte // 16 CPU registers
	index  uint16   // 16 bit register for addresses
	stack  [16]uint16
	PC     uint16 // Program counter
	SP     uint8  // Pointer to top of stack
	//=====  Timers  =====//
	delayTimer byte
	soundTimer byte
	ClockSpeed time.Duration
	//=====  GFX  =====//
	display *display
	//=====  Input  =====//
	keypad [16]bool
	//=====  Misc  =====//
	isPaused bool
}

func NewSystem() *Chip8 {
	// Init pointer to chip8 instance
	inst := &Chip8{
		PC:         0x200, // 0x000 - 0x1FF reserved for interpreter
		ClockSpeed: clockSpeed,
		display:    newDisplay(),
	}
	// Load fontSet into allocated memory
	inst.LoadFontSet()

	return inst
}

func (c *Chip8) Cycle() {
	// Fetch opcode
	currentOpcode := c.fetchOpcode()
	// Execute opcode
	c.executeOpcode(currentOpcode)
}

func (c Chip8) fetchOpcode() uint16 {
	opcode := uint16(c.memory[c.PC])<<8 | uint16(c.memory[c.PC+1])
	return opcode
}

func (c *Chip8) executeOpcode(opcode uint16) {

	var addr uint16 = (opcode & 0x0FFF)
	var nibble byte = uint8((opcode & 0x000F))
	var x byte = uint8((opcode & 0x0F00) >> 8)
	var y byte = uint8((opcode & 0x00F0) >> 4)
	var kk byte = uint8((opcode & 0x00FF))

	switch opcode & 0xF000 {
	case 0x0000:
		switch nibble {
		case 0x000:
			c.CLS()
		case 0x00E:
			c.RET()
		default:
			c.unknownOpcode(opcode)
		}
	case 0x1000:
		c.JP_NNN(addr)
	case 0x2000:
		c.CALL_NNN(addr)
	case 0x3000:
		c.SE_VX_NN(x, kk)
	case 0x4000:
		c.SNE_VX_NN(x, kk)
	case 0x5000:
		c.SE_VX_VY(x, y)
	case 0x6000:
		c.LD_VX_NN(x, kk)
	case 0x7000:
		c.ADD_VX_NN(x, kk)
	case 0x8000:
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
			c.unknownOpcode(opcode)
		}
	case 0x9000:
		c.SNE_VX_VY(x, y)
	case 0xA000:
		c.LD_I_NNN(addr)
	case 0xB000:
		c.JP_V0_NNN(addr)
	case 0xC000:
		c.RND_VX_NN(x, kk)
	case 0xD000:
		c.DRW_VX_VY_N(x, y, nibble)
	case 0xE000:
		switch nibble {
		case 1:
			c.SKNP_VX(x)
		case 0x00E:
			c.SKP_VX(x)
		default:
			c.unknownOpcode(opcode)
		}

	case 0xF000:
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
				c.unknownOpcode(opcode)
			}
		default:
			c.unknownOpcode(opcode)
		}
	default:
		c.unknownOpcode(opcode)
	}
}

func (c *Chip8) UpdateTimers() {
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

func (c *Chip8) LoadROM(file string) {
	path := "./ROMs/" + file
	// Attempt to open program
	// If it doesnt exist list available and prompt user to try again
	rom, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer rom.Close()
	// Check that the program can fit into memory
	info, err := rom.Stat()
	if err != nil {
		log.Fatal(err)
	}
	_size := info.Size()
	fmt.Println("Size of program: ", _size)
	if int(_size) >= len(c.memory) {
		log.Fatal("Program you're trying to load is too large.")
	}
	// Temp array to load hold program info
	tempAlloc := make([]byte, _size)
	n, err := rom.Read(tempAlloc)
	if err != nil {
		log.Fatal(err)
	}
	if n != int(_size) {
		log.Fatal("Error reading the program.")
	}
	// Move data from tempAlloc to CPU memory
	for b := 0; b < int(_size); b++ {
		c.memory[0x200+b] = tempAlloc[b]
	}

}

func (c *Chip8) reset() {
	c.PC = 0x200
	c.index = 0
	c.SP = 0
	c.delayTimer = 0
	c.soundTimer = 0
	c.V = [16]byte{}
	c.stack = [16]uint16{}
	c.display.pixelArray = [64][32]int{}
}
