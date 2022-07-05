package main

import (
	"fmt"
	"github.com/joshuaseligman/GoVM/pkg/hardware"
)

func main() {
	hw := hardware.NewHardware("CPU", 0)
	fmt.Println(hw.GetName())
	fmt.Println(hw.GetId())
}