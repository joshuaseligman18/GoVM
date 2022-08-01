package main

import (
	"log"

	"github.com/joshuaseligman/GoVM/pkg/assembler"
	"github.com/joshuaseligman/GoVM/pkg/gui"
	"github.com/joshuaseligman/GoVM/pkg/hardware/clock"
	"github.com/joshuaseligman/GoVM/pkg/hardware/cpu"
	"github.com/joshuaseligman/GoVM/pkg/hardware/memory"
)

func main() {
	assembledProgram, err := assembler.AssembleProgram("test.goas", 0x10000)
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