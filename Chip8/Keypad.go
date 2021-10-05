package CHIP8

import (
	"github.com/hajimehoshi/ebiten/v2"
)

var keys = map[ebiten.Key]byte{
	ebiten.Key1: 0x1,
	ebiten.Key2: 0x2,
	ebiten.Key3: 0x3,
	ebiten.Key4: 0xC,
	ebiten.KeyQ: 0x4,
	ebiten.KeyW: 0x5,
	ebiten.KeyE: 0x6,
	ebiten.KeyR: 0xD,
	ebiten.KeyA: 0x7,
	ebiten.KeyS: 0x8,
	ebiten.KeyD: 0x9,
	ebiten.KeyF: 0xE,
	ebiten.KeyZ: 0xA,
	ebiten.KeyX: 0x0,
	ebiten.KeyC: 0xB,
	ebiten.KeyV: 0xF,
}

func (c *Chip8) PressedKeys() {
	for key, val := range keys {
		if ebiten.IsKeyPressed(key) {
			c.keymap[val] = true
		}
	}
}

func (c *Chip8) ResetKeys() {
	for k := range c.keymap {
		c.keymap[k] = false
	}
}
