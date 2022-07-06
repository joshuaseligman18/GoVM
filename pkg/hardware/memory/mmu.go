package memory

import (
	"github.com/joshuaseligman/GoVM/pkg/hardware"
)

// The struct for the memory management unit
type Mmu struct {
	hw *hardware.Hardware // The hardware struct
	mar uint // The memory address register
	mdr uint32 // The memory data register
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
	mmu.mdr = mmu.memory.Read(mmu.mar)
}

// Sends the signal to memory to write the value in the MDR to the address of the MAR
func (mmu *Mmu) CallWrite() {
	mmu.memory.Write(mmu.mar, mmu.mdr)
}

// Gets the MAR of the MMU
func (mmu *Mmu) GetMar() uint {
	return mmu.mar
}

// Sets the MAR of the MMU
func (mmu *Mmu) SetMar(newMar uint) {
	mmu.mar = newMar
}

// Gets the MDR of the MMU
func (mmu *Mmu) GetMdr() uint32 {
	return mmu.mdr
}

// Sets the MDR of the MMU
func (mmu *Mmu) SetMdr(newMdr uint32) {
	mmu.mdr = newMdr
}

// Logs a message
func (mmu *Mmu) Log(msg string) {
	mmu.hw.Log(msg)
}