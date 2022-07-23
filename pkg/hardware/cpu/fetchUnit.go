package cpu

import (
	"sync"

	"github.com/joshuaseligman/GoVM/pkg/hardware"
	"github.com/joshuaseligman/GoVM/pkg/hardware/memory"
	"github.com/joshuaseligman/GoVM/pkg/util"
)

type FetchUnit struct {
	hw *hardware.Hardware
	mmu *memory.Mmu
}

func NewFetchUnit(mem *memory.Memory) *FetchUnit {
	fetchUnit := FetchUnit {
		hw: hardware.NewHardware("IFU", 0),
		mmu: memory.NewMmu(mem),
	}
	return &fetchUnit
}

// Fetches the instruction from memory
func (ifu *FetchUnit) FetchInstruction(out chan *IFIDReg, addr *uint64, flushChan chan bool, flushWg *sync.WaitGroup) {
	ifu.mmu.SetMar(*addr)
	ifu.mmu.CallRead()
	ifu.Log(util.ConvertToHexUint32(uint32(ifu.mmu.GetMdr() >> 32)))
	*addr += 4

	// Flush if there is a signal
	if len(flushChan) > 0 {
		ifu.Log("Flushing")
		flushWg.Done()
	} else {
		out <- &IFIDReg {
			instr: uint32(ifu.mmu.GetMdr() >> 32),
			incrementedPC: *addr,
		}
	}
}

// Logs a message
func (ifu *FetchUnit) Log(msg string) {
	ifu.hw.Log(msg)
}