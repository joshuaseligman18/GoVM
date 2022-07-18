package memory

import (
	"github.com/joshuaseligman/GoVM/pkg/hardware"
)

// The struct for the memory management unit
type Mmu struct {
	hw *hardware.Hardware // The hardware struct
	mar uint64 // The memory address register
	mdr uint64 // The memory data register
	memory *Memory
}

// Function that creates the MMU
func NewMmu(mem *Memory) *Mmu {
	mmu := Mmu {
		hw: hardware.NewHardware("MMU", 0),
		mar: 0,
		mdr: 0,
		memory: mem,
	}
	return &mmu
}

// Sends the signal to memory to read the value in the address of the MAR
func (mmu *Mmu) CallRead() {
	var newMdr uint64 = 0x0
	for i := 0; i < 8; i++ {
		newMdr = newMdr << 8 | uint64(mmu.memory.Read(mmu.mar + uint64(i)))
	}
	mmu.mdr = newMdr
}

// Sends the signal to memory to write the value in the MDR to the address of the MAR
func (mmu *Mmu) CallWrite() {
	for i := 0; i < 8; i++ {
		data := uint8((mmu.mdr >> (i * 8)) & 0xFF)
		mmu.memory.Write(mmu.mar + uint64(i), data)
	}
}

// Gets the MAR of the MMU
func (mmu *Mmu) GetMar() uint64 {
	return mmu.mar
}

// Sets the MAR of the MMU
func (mmu *Mmu) SetMar(newMar uint64) {
	mmu.mar = newMar
}

// Gets the MDR of the MMU
func (mmu *Mmu) GetMdr() uint64 {
	return mmu.mdr
}

// Sets the MDR of the MMU
func (mmu *Mmu) SetMdr(newMdr uint64) {
	mmu.mdr = newMdr
}

// Logs a message
func (mmu *Mmu) Log(msg string) {
	mmu.hw.Log(msg)
}