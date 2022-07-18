package memory

import (
	"fmt"

	"github.com/joshuaseligman/GoVM/pkg/hardware"
	"github.com/joshuaseligman/GoVM/pkg/util"
)

// Struct for memory
type Memory struct {
	hw *hardware.Hardware // The hardware struct
	ram []uint8 // The RAM
}

// Creates an empty memory struct
func NewMemory(addressableSpace uint) *Memory {
	mem := Memory { 
		hw: hardware.NewHardware("RAM", 0), 
		ram: make([]uint8, addressableSpace), 
	}
	return &mem
}

// Creates a memory struct with a loaded program
func NewFlashedMemory(program []uint32) *Memory {
	mem := Memory {
		hw: hardware.NewHardware("RAM", 0), 
		ram: make([]uint8, len(program) * 4),
	}
	for i := 0; i < len(program); i++ {
		instr := program[i]
		for j := 0; j < 4; j++ {
			byteInstr := uint8(instr >> (8 * (3 - j)) & 0xFF)
			mem.ram[i * 4 + j] = byteInstr
		}
	}
	return &mem
}

// Sets the MDR to the value stored in memory at the address MAR
func (mem *Memory) Read(addr uint64) uint8 {
	return mem.ram[addr] 
}

// Writes to RAM based on the current values of MAR and MDR
func (mem *Memory) Write(addr uint64, data uint8) {
	mem.ram[addr] = data
}

// Prints the value stored in the given address
func (mem *Memory) PrintMemory(addr uint64) {
	mem.Log(fmt.Sprintf("Addr: %s; Data: %s", util.ConvertToHexUint32(uint32(addr)), util.ConvertToHexUint8(mem.ram[addr])))
}

// Prints a range of memory addresses
func (mem *Memory) MemoryDump(start uint64, end uint64) {
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