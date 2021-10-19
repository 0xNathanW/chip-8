package CHIP8

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"runtime"
	str "strings"
	"time"
)

type Chip8 struct {
	//=====  CPU  =====//
	Memory [4096]byte
	V      [16]byte // 16 CPU registers
	index  uint16   // 16 bit register for addresses
	stack  [16]uint16
	PC     uint16 // Program counter
	SP     uint8  // Pointer to top of stack
	//=====  Timers  =====//
	delayTimer byte
	soundTimer byte
	ClockSpeed time.Duration
	//=====  Output  =====//
	DrawFlag   bool         // Signals whether to draw on cycle
	PixelArray [32][64]byte // Graphics display, 64 by 32 pixels
	//=====  Input  =====//
	Keypad [16]bool
}

func Initialise(clocksPerSecond int) *Chip8 {
	// Init pointer to chip8 instance.
	inst := &Chip8{
		PC:         0x200, // 0x000 - 0x1FF reserved for interpreter.
		ClockSpeed: (time.Second / time.Duration(clocksPerSecond)),
	}
	// Load fontSet into allocated memory.
	inst.LoadFontSet()

	// Load program into program.
	inst.LoadROM()

	return inst
}

// Cycle will be called on each clock of CPU.
func (c *Chip8) Cycle() {
	// Fetch opcode
	currentOpcode := c.fetchOpcode()
	// Execute opcode
	c.executeOpcode(currentOpcode)
	// Update timers
	c.UpdateTimers()
}

func (c Chip8) fetchOpcode() uint16 {
	opcode := uint16(c.Memory[c.PC])<<8 | uint16(c.Memory[c.PC+1])
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
			fmt.Println("Err1")
			c.UnknownOpcode()
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
			fmt.Println("Err2")
			c.UnknownOpcode()
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
			fmt.Println("Err3")
			c.UnknownOpcode()
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
				fmt.Println("Err4")
				c.UnknownOpcode()
			}
		default:
			fmt.Println("Err5")
			c.UnknownOpcode()
		}
	default:
		fmt.Println("Err6")
		c.UnknownOpcode()
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
	// Preset sprites for numbers/letters
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
		c.Memory[i] = fontSet[i]
	}
}

func (c *Chip8) LoadROM() error {

	// Delimiter for reading input based on os.
	// End of line in windows \r\n, linux and mac just \n.
	var delim byte
	if runtime.GOOS == "windows" {
		delim = '\r'
	} else {
		delim = '\n'
	}

	// List available ROMs.
	fmt.Println("Programs available:")
	files, err := ioutil.ReadDir("./ROMs")
	if err != nil {
		log.Fatal("ROM folder does not exist")
	}

	for _, f := range files {
		if path.Ext("./ROMs/"+f.Name()) == ".ch8" {
			fmt.Println(f.Name()[:len(f.Name())-4])
		}
	}

	// Collecting name of ROM.
	fmt.Println("Which program would you like to load:")
	in := bufio.NewReader(os.Stdin)
	input, err := in.ReadString(delim)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return err
	}
	strippedInput := str.TrimSpace(input)
	path := "./ROMs/" + strippedInput + ".ch8"
	// Attempt to open program.
	// If it doesnt exist list available and prompt user to try again.
	rom, openErr := os.Open(path)
	if openErr != nil {
		fmt.Println("Invalid ROM!")
		c.LoadROM()
	}
	defer rom.Close()

	// Check that the program can fit into memory.
	info, err := rom.Stat()
	if err != nil {
		return err
	}
	_size := info.Size()
	fmt.Println("Size of program: ", _size)
	if int(_size) >= len(c.Memory) {
		log.Fatal("Program you're trying to load is too large.")
	}

	// Temp array to load hold program info.
	tempAlloc := make([]byte, _size)
	n, err := rom.Read(tempAlloc)
	if err != nil {
		return err
	}
	if n != int(_size) {
		log.Fatal("Error reading the program.")
	}

	// Move data from tempAlloc to CPU memory.
	for b := 0; b < int(_size); b++ {
		c.Memory[0x200+b] = tempAlloc[b]
	}

	return nil
}
