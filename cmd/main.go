package main

import (
	"fmt"
	"github.com/joshuaseligman/GoVM/pkg/hardware"
)

func main() {
	fmt.Println("Hello world")
	hw := hardware.NewHardware("CPU")
	fmt.Println(hw.GetName())
}