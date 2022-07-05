package main

import (
	"github.com/joshuaseligman/GoVM/pkg/hardware"
)

func main() {
	hw := hardware.NewHardware("CPU", 0)
	hw.Log("hello")
}