package main

import (
	"github.com/joshuaseligman/GoVM/pkg/hardware/memory"
	"github.com/joshuaseligman/GoVM/pkg/util"
	"github.com/joshuaseligman/GoVM/pkg/assembler"
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

	assembler.AssembleProgram("test.goas", 0x10000)
}