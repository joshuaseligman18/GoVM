package cpu

import (
	"fmt"

	"github.com/joshuaseligman/GoVM/pkg/hardware"
	"github.com/joshuaseligman/GoVM/pkg/hardware/memory"
)

// Struct for the CPU
type Cpu struct {
	hw *hardware.Hardware // The hardware struct
	acc uint64 // The accumulator
	reg []uint64 // Other registers
	programCounter uint // The address of the current instruction being fetched
	fetchUnit *FetchUnit // The fetch unit
}

// Creates the CPU
func NewCpu(mem *memory.Memory) *Cpu {
	cpu := Cpu {
		hw: hardware.NewHardware("CPU", 0),
		acc: 0x0,
		reg: make([]uint64, 32),
		programCounter: 0,
		fetchUnit: NewFetchUnit(mem),
	}
	return &cpu
}

// Function that gets called every clock cycle
func (cpu *Cpu) Pulse() {
	instr := cpu.fetchUnit.FetchInstruction(cpu.programCounter)
	cpu.programCounter++
	fmt.Println(cpu.programCounter)
	fmt.Println(instr)
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