package main

import (
	"github.com/joshuaseligman/GoVM/pkg/hardware/memory"
	"github.com/joshuaseligman/GoVM/pkg/util"
)

func main() {
	ram := memory.NewMemory(0x10000)
	ram.SetMar(0x0)
	ram.SetMdr(0x0000000000000005)
	ram.Write()
	ram.SetMar(0x10000)
	ram.SetMdr(0x0)
	ram.SetMar(0x0)
	ram.Read()
	ram.Log(util.ConvertToHex(ram.GetMdr(), 16))
}