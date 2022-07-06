package main

import (
	"fmt"

	"github.com/joshuaseligman/GoVM/pkg/assembler"
	"github.com/joshuaseligman/GoVM/pkg/hardware/clock"
	"github.com/joshuaseligman/GoVM/pkg/hardware/cpu"
	"github.com/joshuaseligman/GoVM/pkg/hardware/memory"
	"github.com/joshuaseligman/GoVM/pkg/util"
	"github.com/joshuaseligman/GoVM/pkg/gui"
)

func main() {
	ram := memory.NewMemory(0x10000)
	ram.SetMar(0x0)
	ram.SetMdr(0x00000005)
	ram.Write()
	ram.SetMar(0x10000)
	ram.SetMdr(0x0)
	ram.SetMar(0x0)
	ram.Read()
	ram.Log(util.ConvertToHex(ram.GetMdr(), 8))

	fmt.Printf("%b", assembler.AssembleProgram("test.goas", 0x10000)[:10])

	cpu := cpu.NewCpu()

	clk := clock.NewClock()

	clk.AddClockListener(cpu)
	go clk.StartClock(500)

	guiData := gui.NewGuiData(cpu)
	
	gui.CreateGui(guiData)
}