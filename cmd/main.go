package main

import (
	"fmt"

	"github.com/joshuaseligman/GoVM/pkg/assembler"
	"github.com/joshuaseligman/GoVM/pkg/hardware/memory"
	"github.com/joshuaseligman/GoVM/pkg/util"
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
}