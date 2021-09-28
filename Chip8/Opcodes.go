package Chip8

import (
	"fmt"
	"math/rand"
)

//----------------------------------  OPCODE FUNCTIONS  -----------------------------//

func (c *Chip8) CLS() {
	// Clear the display
	// NOT DONE
	fmt.Println("lol")
}

func (c *Chip8) RET() {
	// Return from subroutine
	// Set program counter back to top of stack and decrement stack pointer
	c.PC = c.stack[c.SP]
	c.SP--
	c.PC += 2
	/// Check order
}

func (c *Chip8) JP_NNN(addr uint16) {
	// Jump to address
	c.PC = addr
}

func (c *Chip8) CALL_NNN(addr uint16) {
	// Call subroutine at address nnn
	// Store current PC address on stack
	c.stack[c.SP] = c.PC
	// Top of stack inceased by one
	c.SP++
	//	Set program counter to address
	c.PC = addr
}

func (c *Chip8) JP_V0_NNN(addr uint16) {
	// Jump to address + V0
	c.PC = addr + uint16(c.V[0])
}

func (c *Chip8) SE_VX_NN(x, kk byte) {
	// Skip nxt instruction if Vx=kk
	if c.V[x] == kk {
		c.PC += 2
	}
}

func (c *Chip8) SNE_VX_NN(x, kk byte) {
	// Skip nxt instruction if Vx!=kk
	if c.V[x] != kk {
		c.PC += 2
	}
}

func (c *Chip8) SE_VX_VY(x, y byte) {
	// Skip nxt instruction if Vx=Vy
	if c.V[x] == c.V[y] {
		c.PC += 2
	}
}

func (c *Chip8) SNE_VX_VY(x, y byte) {
	// Skip nxt instruction if Vx!=Vy
	if c.V[x] != c.V[y] {
		c.PC += 2
	}
}

func (c *Chip8) SKP_VX(x byte) {
	// Skip next instruction if key(Vx) is pressed
	fmt.Println("lol")
	c.PC += 2
}

func (c *Chip8) SKNP_VX(x byte) {
	// Skip next instruction if key(Vx) is not pressed
	fmt.Println("lol")
	c.PC += 2
}

func (c *Chip8) LD_VX_K(x byte) {
	// Wait for key press, store key pressed in Vx
	fmt.Println("lol")
}

func (c *Chip8) LD_VX_NN(x, kk byte) {
	// Set Vx to kk
	c.V[x] = kk
	c.PC += 2
}

func (c *Chip8) LD_VX_VY(x, y byte) {
	// Set Vx to Vy
	c.V[x] = c.V[y]
	c.PC += 2
}

func (c *Chip8) LD_VX_DT(x byte) {
	// Set Vx to delay timer value
	c.V[x] = c.delayTimer
	c.PC += 2
}

func (c *Chip8) LD_DT_VX(x byte) {
	// Set delay timer to Vx
	c.delayTimer = c.V[x]
	c.PC += 2
}

func (c *Chip8) LD_ST_VX(x byte) {
	// Set sound timer to Vx
	c.soundTimer = c.V[x]
	c.PC += 2
}

func (c *Chip8) LD_I_NNN(addr uint16) {
	// Set index to address
	c.index = addr
	c.PC += 2
}

func (c *Chip8) LD_F_VX(x byte) {
	// Set I = location of sprite for digit Vx
	// NOT DONE
	fmt.Println("lol")
}

func (c *Chip8) LD_I_VX(x byte) {
	// Store registers V0 through Vx in memory starting at location I
	startAddr := int(c.index)
	for i := 0; i < int(x); i++ {
		if startAddr+i < 4096 {
			c.memory[i+startAddr] = c.V[i]
		}
	}
	c.PC += 2
}

func (c *Chip8) LD_VX_I(x byte) {
	// Read registers V0 through Vx from memory starting at location I
	startAddr := int(c.index)
	for i := 0; i < int(x); i++ {
		if startAddr+i < 4096 {
			c.V[i] = c.memory[i+startAddr]
		}
	}
	c.PC += 2
}

func (c *Chip8) ADD_I_VX(x byte) {
	// Set I = I + Vx; Vf = 1 if I > 0xFFF else 0
	c.index += uint16(c.V[x])
	if c.index > 0xFFF {
		c.V[15] = 1
	} else {
		c.V[15] = 0
	}
	c.PC += 2
}

func (c *Chip8) ADD_VX_NN(x, kk byte) {
	// Set Vx = Vx + kk
	c.V[x] += kk
	c.PC += 2
}

func (c *Chip8) ADD_VX_VY(x, y byte) {
	// Flag register is 1 if Vx bigger than Vy else 0
	if c.V[x] > c.V[y] {
		c.V[15] = 1
	} else {
		c.V[15] = 0
	}
	// Add Vx to Vy
	c.V[x] += c.V[y]
	c.PC += 2
}

func (c *Chip8) SUB_VX_VY(x, y byte) {
	// Set Vx = Vx - Vy; VF=1 if not borrow else 0
	// Flag register is 1 if Vx greater than Vy else 0
	if c.V[x] > c.V[y] {
		c.V[15] = 1
	} else {
		c.V[15] = 0
	}
	// Subtract Vy from Vx
	c.V[x] -= c.V[y]
	c.PC += 2
}

func (c *Chip8) SUBN_VX_VY(x, y byte) {
	// Set Vx = Vy - Vx; VF=1 if not borrow else 0
	if c.V[y] > c.V[x] {
		c.V[15] = 1
	} else {
		c.V[15] = 0
	}
	// Subtract Vy from Vx
	c.V[x] = c.V[y] - c.V[x]
	c.PC += 2
}

func (c *Chip8) OR_VX_VY(x, y byte) {
	// Vx = Vx OR Vy
	c.V[x] = c.V[x] | c.V[y]
	c.PC += 2
}

func (c *Chip8) AND_VX_VY(x, y byte) {
	// Vx = Vx AND Vy
	c.V[x] = c.V[x] & c.V[y]
	c.PC += 2
}

func (c *Chip8) XOR_VX_NN(x, y byte) {
	// Vx = Vx XOR Vy
	c.V[x] = c.V[x] ^ c.V[y]
	c.PC += 2
}

func (c *Chip8) SHR_VX(x byte) {
	// Set Vf to Vx least significent bit
	c.V[15] = c.V[x] & 1
	// Bit shift Vx right (divide by 2)
	c.V[x] >>= 1
	c.PC += 2
}

func (c *Chip8) SHL_VX(x byte) {
	// Set Vf to Vx most significent bit
	c.V[15] = c.V[x] >> 7
	// Bit shift Vx left (multiply by 2)
	c.V[x] <<= 1
	c.PC += 2
}

func (c *Chip8) BCD_VX(x byte) {
	// Storing binary-coded decimal representation of Vx at memory addresses
	c.memory[c.index] = c.V[x] / 100
	c.memory[c.index+1] = (c.V[x] / 10) % 10
	c.memory[c.index+2] = (c.V[x] % 100) % 10
	c.PC += 2
}

func (c *Chip8) RND_VX_NN(x, kk byte) {
	// Vx = Random num AND kk
	c.V[x] = byte(rand.Intn(255)) & kk
	c.PC += 2
}

func (c *Chip8) DRW_VX_VY_N(x, y, nibble byte) {
	// NOT DONE
	fmt.Println("lol")
}
