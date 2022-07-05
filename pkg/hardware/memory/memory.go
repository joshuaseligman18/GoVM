package memory

import "github.com/joshuaseligman/GoVM/pkg/hardware"

// Struct for memory
type Memory struct {
	hw *hardware.Hardware
	ram []int64
}

// Creates a new memory struct
func NewMemory(addressableSpace int) *Memory {
	mem := Memory { hw: hardware.NewHardware("RAM", 0), ram: make([]int64, addressableSpace)}
	return &mem
}

// Logs a message
func (mem *Memory) Log(msg string) {
	mem.hw.Log(msg)
}