package main

import (
	"log"

	"github.com/joshuaseligman/GoVM/pkg/assembler"
	"github.com/joshuaseligman/GoVM/internal/gui"
	"github.com/joshuaseligman/GoVM/pkg/hardware/clock"
	"github.com/joshuaseligman/GoVM/pkg/hardware/cpu"
	"github.com/joshuaseligman/GoVM/pkg/hardware/memory"
)

func main() {

	prog := 
`ADD X0, X1, X2
MOVZ X5, #0x5, LSL 0
MOVZ X0, #0xABCD, LSL 0
MOVK X0, #0x8456, LSL 16
B #0x6
MOVK X1, #0x1234, LSL 0
MOVK X2, #0x1234, LSL 0
MOVK X3, #0x1234, LSL 0
MOVK X4, #0x1234, LSL 0
MOVK X5, #0x1234, LSL 0
MOVK X6, #0x1234, LSL 0
MOVK X0, #0x1234, LSL 16
SUBI X5, X5, #0x1
CBNZ X5, #0x7FFF5
HLT
MOVK X7, #0x1234, LSL 0
MOVK X8, #0x1234, LSL 0
MOVK X9, #0x1234, LSL 0
MOVK X10, #0x1234, LSL 0`

	assembledProgram, err := assembler.AssembleProgramAPI(prog)
	if err != nil {
		log.Fatal(err)
	}
	
	mem := memory.NewFlashedMemory(assembledProgram)
	
	mem.MemoryDump(0, 30)
	
	cpu := cpu.NewCpu(mem)
	guiData := gui.NewGuiData(cpu)

	clk := clock.NewClock()

	clk.AddClockListener(guiData)
	clk.AddClockListener(cpu)
	
	go clk.StartClock(500)
	gui.CreateGui(guiData)
}