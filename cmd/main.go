package main

import (
	"log"

	"github.com/joshuaseligman/GoVM/pkg/assembler"
	// "github.com/joshuaseligman/GoVM/internal/gui"
	"github.com/joshuaseligman/GoVM/pkg/hardware/clock"
	"github.com/joshuaseligman/GoVM/pkg/hardware/cpu"
	"github.com/joshuaseligman/GoVM/pkg/hardware/memory"
)

func main() {

	assembledProgram, err := assembler.AssembleProgramFile("test.goas", 0x1000)
	if err != nil {
		log.Fatal(err)
	}
	
	mem := memory.NewFlashedMemory(assembledProgram)
	
	mem.MemoryDump(0, 30)
	
	clk := clock.NewClock()
	
	cpu := cpu.NewCpu(mem, clk)
	// guiData := gui.NewGuiData(cpu)

	// clk.AddClockListener(guiData)
	clk.AddClockListener(cpu)
	
	clk.StartClock(500)
	// gui.CreateGui(guiData)
}