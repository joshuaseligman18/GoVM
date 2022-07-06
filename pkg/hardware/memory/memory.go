package memory

import (
	"fmt"

	"github.com/joshuaseligman/GoVM/pkg/hardware"
	"github.com/joshuaseligman/GoVM/pkg/util"
)

// Struct for memory
type Memory struct {
	hw *hardware.Hardware // The hardware struct
	ram []uint32 // The RAM
	mar uint // The memory address register
	mdr uint32 // The memory data register
}

// Creates an empty memory struct
func NewMemory(addressableSpace uint) *Memory {
	mem := Memory { 
		hw: hardware.NewHardware("RAM", 0), 
		ram: make([]uint32, addressableSpace), 
		mar: 0,
		mdr: 0,
	}
	return &mem
}

// Creates a memory struct with a loaded program
func NewFlashedMemory(program []uint32) *Memory {
	mem := Memory {
		hw: hardware.NewHardware("RAM", 0), 
		ram: program,
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

// Prints the value stored in the given address
func (mem *Memory) PrintMemory(addr uint) {
	mem.Log(fmt.Sprintf("Addr: %s; Data: %s", util.ConvertToHexUint32(uint32(addr), 8), util.ConvertToHexUint32(mem.ram[addr], 8)))
}

// Prints a range of memory addresses
func (mem *Memory) MemoryDump(start uint, end uint) {
	mem.Log("***** Start of Memory Dump *****")
	for i := start; i <= end; i++ {
		mem.PrintMemory(i)
	}
	mem.Log("***** End of Memory Dump *****")
}

// Logs a message
func (mem *Memory) Log(msg string) {
	mem.hw.Log(msg)
}