package cpu

import (
	"github.com/joshuaseligman/GoVM/pkg/hardware"
	"github.com/joshuaseligman/GoVM/pkg/hardware/memory"
)

type FetchUnit struct {
	hw *hardware.Hardware
	mmuInstr *memory.Mmu
}

func NewFetchUnit() *FetchUnit {
	fetchUnit := FetchUnit {
		hw: hardware.NewHardware("IFU", 0),
		mmuInstr: memory.NewMmu(),
	}
	return &fetchUnit
}