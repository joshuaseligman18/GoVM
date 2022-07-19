package cpu

import (
	"github.com/joshuaseligman/GoVM/pkg/hardware"
	"github.com/joshuaseligman/GoVM/pkg/hardware/memory"
	"github.com/joshuaseligman/GoVM/pkg/util"
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

	case 0x1C2: // LDURB
		// Set the address and read the next 64 bits
		mdu.mmu.SetMar(exmemReg.writeVal)
		mdu.mmu.CallRead()
		// Return only the first 8 bits
		result := mdu.mmu.GetMdr() >> 56
		out <- &MEMWBReg {
			instr: exmemReg.instr,
			incrementedPC: exmemReg.incrementedPC,
			writeVal: result,
		}

	case 0x3C2: // LDURH
		// Set the address and read the next 64 bits
		mdu.mmu.SetMar(exmemReg.writeVal)
		mdu.mmu.CallRead()
		// Return only the first 16 bits
		result := mdu.mmu.GetMdr() >> 48
		out <- &MEMWBReg {
			instr: exmemReg.instr,
			incrementedPC: exmemReg.incrementedPC,
			writeVal: result,
		}

	case 0x5C4: // LDURSW
		// Set the address and read the next 64 bits
		mdu.mmu.SetMar(exmemReg.writeVal)
		mdu.mmu.CallRead()
		// Get first 32 bits
		orig := uint32(mdu.mmu.GetMdr() >> 32)
		result := util.SignExtend(orig)
		out <- &MEMWBReg {
			instr: exmemReg.instr,
			incrementedPC: exmemReg.incrementedPC,
			writeVal: result,
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