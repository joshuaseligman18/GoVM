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
func (idu *DecodeUnit) DecodeInstruction(ifidReg *IFIDReg) *IDEXReg {
	first9Bits := ifidReg.instr >> 21
	idu.Log(fmt.Sprintf("%X", first9Bits))
	switch first9Bits {
	// IM instructions
	case 0x694, 0x695, 0x696, 0x697: // MOVZ
		// Register to write to
		regWrite := ifidReg.instr & 0x1F
		// Immediate to write
		immediate := ifidReg.instr & 0xFFFF >> 5
		
		if idu.cpu.GetRegisterLocks().Contains(regWrite) {
			return nil
		} else {
			idu.cpu.GetRegisterLocks().Enqueue(regWrite)
			return &IDEXReg {
				instr: ifidReg.instr,
				incrementedPC: ifidReg.incrementedPC,
				regReadData1: 0,
				regReadData2: 0,
				signExtendImm: signExtend(immediate),
			}
		}

	case 0x794, 0x795, 0x796, 0x797: // MOVK
		// Register to write to
		regWrite := ifidReg.instr & 0x1F
		// Register to read from
		regReadData1 := idu.cpu.GetRegisters()[regWrite]
		// Immediate to write
		immediate := ifidReg.instr & 0xFFFF >> 5

		if idu.cpu.GetRegisterLocks().Contains(regWrite) {
			return nil
		} else {
			idu.cpu.GetRegisterLocks().Enqueue(regWrite)
			return &IDEXReg {
				instr: ifidReg.instr,
				incrementedPC: ifidReg.incrementedPC,
				regReadData1: regReadData1,
				regReadData2: 0,
				signExtendImm: signExtend(immediate),
			}
		}
		
	}
	
	return nil
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