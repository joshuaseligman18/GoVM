package cpu

import (
	"github.com/joshuaseligman/GoVM/pkg/hardware"
	"github.com/joshuaseligman/GoVM/pkg/hardware/memory"
	"github.com/joshuaseligman/GoVM/pkg/util"
)

// The struct for the fetch unit
type FetchUnit struct {
	hw  *hardware.Hardware // The hardware component
	mmu *memory.Mmu // A MMU to interface with memory to fetch the instructions
}

// Creates a new fetch unit
func NewFetchUnit(mem *memory.Memory) *FetchUnit {
	fetchUnit := FetchUnit{
		hw:  hardware.NewHardware("IFU", 0),
		mmu: memory.NewMmu(mem),
	}
	return &fetchUnit
}

// Fetches the instruction from memory
func (ifu *FetchUnit) FetchInstruction(out chan *IFIDReg, addr *uint64) {
	ifu.mmu.SetMar(*addr)
	ifu.mmu.CallRead()
	ifu.Log(util.ConvertToHexUint32(uint32(ifu.mmu.GetMdr() >> 32)))
	*addr += 4

	out <- &IFIDReg{
		Instr:         uint32(ifu.mmu.GetMdr() >> 32),
		IncrementedPC: *addr,
	}
}

// Logs a message
func (ifu *FetchUnit) Log(msg string) {
	ifu.hw.Log(msg)
}

// Resets the fetch unit
func (ifu *FetchUnit) Reset() {
	ifu.mmu.Reset()
}
