package cpu

import "github.com/joshuaseligman/GoVM/pkg/hardware"

// Struct for the CPU
type Cpu struct {
	hw *hardware.Hardware
	acc uint64
	reg []uint64
}

// Creates the CPU
func NewCpu() *Cpu {
	cpu := Cpu {
		hw: hardware.NewHardware("CPU", 0),
		acc: 0x0,
		reg: make([]uint64, 32),
	}
	return &cpu
}

// Function that gets called every clock cycle
func (cpu *Cpu) Pulse() {
	cpu.Log("pulse received")
}

// Logs a message
func (cpu *Cpu) Log(msg string) {
	cpu.hw.Log(msg)
}

// Gets the accumulator
func (cpu *Cpu) GetAcc() uint64 {
	return cpu.acc
}

// Gets the registers
func (cpu *Cpu) GetRegisters() []uint64 {
	return cpu.reg
}