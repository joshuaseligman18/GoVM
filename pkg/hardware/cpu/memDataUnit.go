package cpu

import (
	"github.com/joshuaseligman/GoVM/pkg/hardware"
	"github.com/joshuaseligman/GoVM/pkg/hardware/memory"
	"github.com/joshuaseligman/GoVM/pkg/util"
)

// Struct for the memory data unit
type MemDataUnit struct {
	hw  *hardware.Hardware // The hardware component
	mmu *memory.Mmu
}

// Function that creates the memory data unit
func NewMemDataUnit(mem *memory.Memory) *MemDataUnit {
	mdu := MemDataUnit{
		hw:  hardware.NewHardware("MDU", 0),
		mmu: memory.NewMmu(mem),
	}
	return &mdu
}

// Function that handles the memory access for an instruction
func (mdu *MemDataUnit) HandleMemoryAccess(out chan *MEMWBReg, exmemReg *EXMEMReg) {
	opcode := exmemReg.Instr >> 21
	switch opcode {
	// Memory access instructions
	case 0x7C2: // LDUR
		// Set the address and read the next 64 bits
		mdu.mmu.SetMar(exmemReg.WorkingAddr)
		mdu.mmu.CallRead()
		out <- &MEMWBReg{
			Instr:         exmemReg.Instr,
			IncrementedPC: exmemReg.IncrementedPC,
			WriteVal:      mdu.mmu.GetMdr(),
		}

	case 0x1C2: // LDURB
		// Set the address and read the next 64 bits
		mdu.mmu.SetMar(exmemReg.WorkingAddr)
		mdu.mmu.CallRead()
		// Return only the first 8 bits
		result := mdu.mmu.GetMdr() >> 56
		out <- &MEMWBReg{
			Instr:         exmemReg.Instr,
			IncrementedPC: exmemReg.IncrementedPC,
			WriteVal:      result,
		}

	case 0x3C2: // LDURH
		// Set the address and read the next 64 bits
		mdu.mmu.SetMar(exmemReg.WorkingAddr)
		mdu.mmu.CallRead()
		// Return only the first 16 bits
		result := mdu.mmu.GetMdr() >> 48
		out <- &MEMWBReg{
			Instr:         exmemReg.Instr,
			IncrementedPC: exmemReg.IncrementedPC,
			WriteVal:      result,
		}

	case 0x5C4: // LDURSW
		// Set the address and read the next 64 bits
		mdu.mmu.SetMar(exmemReg.WorkingAddr)
		mdu.mmu.CallRead()
		// Get first 32 bits
		orig := uint32(mdu.mmu.GetMdr() >> 32)
		result := util.SignExtend(orig, 32)
		out <- &MEMWBReg{
			Instr:         exmemReg.Instr,
			IncrementedPC: exmemReg.IncrementedPC,
			WriteVal:      result,
		}

	case 0x7C0: // STUR
		// Set the address and read the next 64 bits
		mdu.mmu.SetMar(exmemReg.WorkingAddr)
		mdu.mmu.SetMdr(exmemReg.WriteVal)
		mdu.mmu.CallWrite(64)
		out <- &MEMWBReg{
			Instr:         exmemReg.Instr,
			IncrementedPC: exmemReg.IncrementedPC,
		}

	case 0x1C0: // STURB
		// Set the address and read the next 64 bits
		mdu.mmu.SetMar(exmemReg.WorkingAddr)
		mdu.mmu.SetMdr(exmemReg.WriteVal)
		mdu.mmu.CallWrite(8)
		out <- &MEMWBReg{
			Instr:         exmemReg.Instr,
			IncrementedPC: exmemReg.IncrementedPC,
		}

	case 0x3C0: // STURH
		// Set the address and read the next 64 bits
		mdu.mmu.SetMar(exmemReg.WorkingAddr)
		mdu.mmu.SetMdr(exmemReg.WriteVal)
		mdu.mmu.CallWrite(16)
		out <- &MEMWBReg{
			Instr:         exmemReg.Instr,
			IncrementedPC: exmemReg.IncrementedPC,
		}

	case 0x5C0: // STURW
		// Set the address and read the next 64 bits
		mdu.mmu.SetMar(exmemReg.WorkingAddr)
		mdu.mmu.SetMdr(exmemReg.WriteVal)
		mdu.mmu.CallWrite(32)
		out <- &MEMWBReg{
			Instr:         exmemReg.Instr,
			IncrementedPC: exmemReg.IncrementedPC,
		}

	// All other instructions
	default:
		out <- &MEMWBReg{
			Instr:         exmemReg.Instr,
			IncrementedPC: exmemReg.IncrementedPC,
			WriteVal:      exmemReg.WriteVal,
		}
	}
}

// Logs a message
func (mdu *MemDataUnit) Log(msg string) {
	mdu.hw.Log(msg)
}

// Resets the memory data unit
func (mdu *MemDataUnit) Reset() {
	mdu.mmu.Reset()
}
