package main

import (
	"github.com/joshuaseligman/GoVM/pkg/hardware/memory"
)

func main() {
	ram := memory.NewMemory(0x10000)
	ram.Log("Hello memory")
}