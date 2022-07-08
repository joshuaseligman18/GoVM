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

func (mdu *MemDataUnit) HandleMemoryAccess(exmemReg *EXMEMReg) *MEMWBReg {
	opcode := exmemReg.instr >> 21
	switch opcode {
	// Memory access instructions

	// All other instructions
	default:
		return & MEMWBReg{
			instr: exmemReg.instr,
			writeVal: exmemReg.writeVal,
		}
	}
}

// Logs a message
func (mdu *MemDataUnit) Log(msg string) {
	mdu.hw.Log(msg)
}