package chip8

import (
	"time"
)

func (c *Chip8) Run(gfx chan<- [64][32]int) {
	clock := time.NewTicker(c.ClockSpeed)
	// Timers run at a set 60Hz rate.
	timers := time.NewTicker(time.Second / 60)

	for {
		select {
		case <-clock.C:
			c.Cycle()
			if c.DrawFlag {
				gfx <- c.Display
				c.DrawFlag = false
			}
		case <-timers.C:
			c.UpdateTimers()
		}
	}
}
