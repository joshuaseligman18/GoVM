package cpu

import (
	"fmt"

	"github.com/joshuaseligman/GoVM/pkg/hardware"
)

// Struct for the decode unit
type DecodeUnit struct {
	hw *hardware.Hardware
	cpu *Cpu
}

// Function that creates the decode unit
func NewDecodeUnit(parentCpu *Cpu) *DecodeUnit {
	decodeUnit := DecodeUnit {
		hw: hardware.NewHardware("IDU", 0),
		cpu: parentCpu,
	}
	return &decodeUnit
}

// Function that decodes an instruction into its operands
func (idu *DecodeUnit) DecodeInstruction(out chan *IDEXReg, ifidReg *IFIDReg) {
	opcode := ifidReg.instr >> 21
	idu.Log(fmt.Sprintf("%X", opcode))
	switch opcode {
	// IM instructions
	case 0x694, 0x695, 0x696, 0x697: // MOVZ
		// Register to write to
		regWrite := ifidReg.instr & 0x1F
		// Immediate to write
		immediate := ifidReg.instr & 0x1FFFFF >> 5
		
		// Wait until register opens up
		for idu.cpu.GetRegisterLocks().Contains(regWrite) {
			continue
		}

		idu.cpu.GetRegisterLocks().Enqueue(regWrite)
		out <- &IDEXReg {
			instr: ifidReg.instr,
			incrementedPC: ifidReg.incrementedPC,
			regReadData1: 0,
			regReadData2: 0,
			signExtendImm: signExtend(immediate), // Should always be positive
		}

	case 0x794, 0x795, 0x796, 0x797: // MOVK
		// Register to write to
		regWrite := ifidReg.instr & 0x1F

		// Wait until the updated value is written
		for idu.cpu.GetRegisterLocks().Contains(regWrite) {
			continue
		}
		// Register to read from
		regReadData1 := idu.cpu.GetRegisters()[regWrite]
		// Immediate to write
		immediate := ifidReg.instr & 0x1FFFFF >> 5
		
		idu.cpu.GetRegisterLocks().Enqueue(regWrite)
		out <- &IDEXReg {
			instr: ifidReg.instr,
			incrementedPC: ifidReg.incrementedPC,
			regReadData1: regReadData1,
			regReadData2: 0,
			signExtendImm: signExtend(immediate), // Should always be positive
		}

	case 0x458: // ADD
		// Registers to read from
		reg1 := ifidReg.instr & 0x1FFFFF >> 16
		reg2 := ifidReg.instr & 0x3FF >> 5
		
		// Wait until the registers have the most up-to-date data
		for idu.cpu.GetRegisterLocks().Contains(reg1) || idu.cpu.GetRegisterLocks().Contains(reg2){
			continue
		}

		regData1 := idu.cpu.GetRegisters()[reg1]
		regData2 := idu.cpu.GetRegisters()[reg2]

		// Add the write register to the queue
		regWrite := ifidReg.instr & 0x1F
		idu.cpu.GetRegisterLocks().Enqueue(regWrite)

		out <- &IDEXReg {
			instr: ifidReg.instr,
			incrementedPC: ifidReg.incrementedPC,
			regReadData1: regData1,
			regReadData2: regData2,
			signExtendImm: 0,
		}
	}
}

// Logs a message
func (idu *DecodeUnit) Log(msg string) {
	idu.hw.Log(msg)
}

// Sign extends a uint32 to a uint64
func signExtend(val uint32) uint64 {
	// Get the sign
	sign := uint64(val >> 31)
	longSign := uint64(0)
	// Repeat it 32 times
	for i := 0; i < 32; i++ {
		longSign = longSign << 1 | sign
	}
	// Combine the original value with the long sign
	result := longSign << 32 | uint64(val)
	return result
}