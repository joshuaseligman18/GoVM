package cpu

import "github.com/joshuaseligman/GoVM/pkg/hardware"

// Struct for the CPU
type Cpu struct {
	hw *hardware.Hardware
}

// Creates the CPU
func NewCpu() *Cpu {
	cpu := Cpu { hw: hardware.NewHardware("CPU", 0) }
	return &cpu
}

// Logs a message
func (cpu *Cpu) Log(msg string) {
	cpu.hw.Log(msg)
}