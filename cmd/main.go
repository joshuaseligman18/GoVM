package main

import (
	"fmt"

	"github.com/joshuaseligman/GoVM/pkg/assembler"
	"github.com/joshuaseligman/GoVM/pkg/hardware/clock"
	"github.com/joshuaseligman/GoVM/pkg/hardware/cpu"
	"github.com/joshuaseligman/GoVM/pkg/gui"
)

func main() {
	fmt.Printf("%b\n", assembler.AssembleProgram("test.goas", 0x10000)[:10])

	cpu := cpu.NewCpu()

	clk := clock.NewClock()

	clk.AddClockListener(cpu)
	
	guiData := gui.NewGuiData(cpu)
	clk.AddClockListener(guiData)
	
	go clk.StartClock(500)
	gui.CreateGui(guiData)
}