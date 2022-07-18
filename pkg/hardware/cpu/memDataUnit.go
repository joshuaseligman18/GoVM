package cpu

import (
	"github.com/joshuaseligman/GoVM/pkg/hardware"
	"github.com/joshuaseligman/GoVM/pkg/hardware/memory"
)

// Struct for the memory data unit
type MemDataUnit struct {
	hw *hardware.Hardware // The hardware component
	mmu *memory.Mmu
}

// Function that creates the memory data unit
func NewMemDataUnit(mem *memory.Memory) *MemDataUnit {
	mdu := MemDataUnit {
		hw: hardware.NewHardware("MDU", 0),
		mmu: memory.NewMmu(mem),
	}
	return &mdu
}

func (mdu *MemDataUnit) HandleMemoryAccess(out chan *MEMWBReg, exmemReg *EXMEMReg) {
	opcode := exmemReg.instr >> 21
	switch opcode {
	// Memory access instructions
	case 0x7C2: // LDUR
		// Set the address and read the next 64 bits
		mdu.mmu.SetMar(exmemReg.writeVal)
		mdu.mmu.CallRead()
		out <- &MEMWBReg {
			instr: exmemReg.instr,
			incrementedPC: exmemReg.incrementedPC,
			writeVal: mdu.mmu.GetMdr(),
		}

	// All other instructions
	default:
		out <- &MEMWBReg {
			instr: exmemReg.instr,
			incrementedPC: exmemReg.incrementedPC,
			writeVal: exmemReg.writeVal,
		}
	}
}

// Logs a message
func (mdu *MemDataUnit) Log(msg string) {
	mdu.hw.Log(msg)
}