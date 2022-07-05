package memory

import "github.com/joshuaseligman/GoVM/pkg/hardware"

// Struct for memory
type Memory struct {
	hw *hardware.Hardware
	ram []uint32
	mar uint
	mdr uint32
}

// Creates a new memory struct
func NewMemory(addressableSpace uint) *Memory {
	mem := Memory { 
		hw: hardware.NewHardware("RAM", 0), 
		ram: make([]uint32, addressableSpace), 
		mar: 0,
		mdr: 0,
	}
	return &mem
}

// Sets the MDR to the value stored in memory at the address MAR
func (mem *Memory) Read() {
	mem.mdr = mem.ram[mem.mar] 
}

// Writes to RAM based on the current values of MAR and MDR
func (mem *Memory) Write() {
	mem.ram[mem.mar] = mem.mdr
}

// Gets the MAR
func (mem *Memory) GetMar() uint {
	return mem.mar
}

// Gets the MDR
func (mem *Memory) GetMdr() uint32 {
	return mem.mdr
}

// Sets the MAR
func (mem *Memory) SetMar(newMar uint) {
	mem.mar = newMar
}

// Sets the MDR
func (mem *Memory) SetMdr(newMdr uint32) {
	mem.mdr = newMdr
}

// Logs a message
func (mem *Memory) Log(msg string) {
	mem.hw.Log(msg)
}