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
	ifidReg *IFIDReg // The register between the fetch and decode units
}

// Creates the CPU
func NewCpu(mem *memory.Memory) *Cpu {
	cpu := Cpu {
		hw: hardware.NewHardware("CPU", 0),
		acc: 0x0,
		reg: make([]uint64, 32),
		programCounter: 0,
		fetchUnit: NewFetchUnit(mem),
		ifidReg: NewIFIDReg(0, 0),
	}
	return &cpu
}

// Function that gets called every clock cycle
func (cpu *Cpu) Pulse() {
	fetchOut := cpu.fetchUnit.FetchInstruction(&cpu.programCounter)
	cpu.ifidReg = fetchOut

	fmt.Println(cpu.ifidReg)
	fmt.Println(cpu.programCounter)
}

// Logs a message
func (cpu *Cpu) Log(msg string) {
	cpu.hw.Log(msg)
}

// Gets the program counter
func (cpu *Cpu) GetProgramCounter() uint{
	return cpu.programCounter
}

// Gets the accumulator
func (cpu *Cpu) GetAcc() uint64 {
	return cpu.acc
}

// Gets the registers
func (cpu *Cpu) GetRegisters() []uint64 {
	return cpu.reg
}