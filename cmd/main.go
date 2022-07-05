package main

import (
	"fmt"
	"github.com/joshuaseligman/GoVM/internal/hardware"
)

func main() {
	hw := hardware.NewHardware("CPU", 0)
	fmt.Println(hw.GetName())
	fmt.Println(hw.GetId())
}