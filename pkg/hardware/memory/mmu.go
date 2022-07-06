package memory

import (
	"github.com/joshuaseligman/GoVM/pkg/hardware"
)

// The struct for the memory management unit
type Mmu struct {
	hw *hardware.Hardware // The hardware struct
	mar int // The memory address register
	mdr uint32 // The memory data register
}

// Function that creates the MMU
func NewMmu() *Mmu {
	mmu := Mmu {
		hw: hardware.NewHardware("MMU", 0),
		mar: 0,
		mdr: 0,
	}
	return &mmu
}

// Logs a message
func (mmu *Mmu) Log(msg string) {
	mmu.hw.Log(msg)
}