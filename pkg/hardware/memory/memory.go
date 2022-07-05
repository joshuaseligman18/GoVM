package memory

import "github.com/joshuaseligman/GoVM/pkg/hardware"

// Struct for memory
type Memory struct {
	hw *hardware.Hardware
	ram []int64
	mar int
	mdr int64
}

// Creates a new memory struct
func NewMemory(addressableSpace int) *Memory {
	mem := Memory { 
		hw: hardware.NewHardware("RAM", 0), 
		ram: make([]int64, addressableSpace), 
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
func (mem *Memory) GetMar() int {
	return mem.mar
}

// Gets the MDR
func (mem *Memory) GetMdr() int64 {
	return mem.mdr
}

// Sets the MAR
func (mem *Memory) SetMar(newMar int) {
	mem.mar = newMar
}

// Sets the MDR
func (mem *Memory) SetMdr(newMdr int64) {
	mem.mdr = newMdr
}

// Logs a message
func (mem *Memory) Log(msg string) {
	mem.hw.Log(msg)
}