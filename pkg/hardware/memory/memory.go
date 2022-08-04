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

// Struct for the API
type MemoryAPI struct {
	Ram [0x10000]uint8 `json:"ram"` // The RAM
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

// Creates a new memory struct for the API
func NewEmptyMemory(size int) *Memory {
	mem := Memory {
		hw: hardware.NewHardware("RAM", 0), 
		ram: make([]uint8, size),
	}
	return &mem
}

// Flashes a program to the beginning of the memory array
func (mem *Memory) FlashProgram(program []uint32) {
	for i := 0; i < len(program); i++ {
		instr := program[i]
		for j := 0; j < 4; j++ {
			byteInstr := uint8(instr >> (8 * (3 - j)) & 0xFF)
			mem.ram[i * 4 + j] = byteInstr
		}
	}
}

// Resets the memory to 0x0 for all memory locations
func (mem *Memory) ResetMemory() {
	for i := 0; i < len(mem.ram); i++ {
		mem.ram[i] = 0x0
	}
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

// Function that converts the memory struct to an API friendly struct
func (mem *Memory) ConvertAPI() *MemoryAPI {
	newRam := [0x10000]uint8{}
	copy(newRam[:], mem.ram)
	return & MemoryAPI {
		Ram: newRam,
	}
}