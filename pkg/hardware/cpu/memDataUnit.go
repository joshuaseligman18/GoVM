package cpu

import "github.com/joshuaseligman/GoVM/pkg/hardware"

// Struct for the memory data unit
type MemDataUnit struct {
	hw *hardware.Hardware // The hardware component
}

// Function that creates the memory data unit
func NewMemDataUnit() *MemDataUnit {
	mdu := MemDataUnit {
		hw: hardware.NewHardware("MDU", 0),
	}
	return &mdu
}

// Logs a message
func (mdu *MemDataUnit) Log(msg string) {
	mdu.hw.Log(msg)
}