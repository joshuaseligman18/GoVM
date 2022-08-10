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

	clk := clock.NewClock()
	
	mem := memory.NewFlashedMemory(assembledProgram, clk)
		
	cpu := cpu.NewCpu(mem, clk)
	// guiData := gui.NewGuiData(cpu)

	// clk.AddClockListener(guiData)
	clk.AddClockListener(cpu)
	
	clk.StartClock(1000)

	// gui.CreateGui(guiData)
}