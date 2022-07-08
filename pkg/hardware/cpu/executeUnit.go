package cpu

import (
	"github.com/joshuaseligman/GoVM/pkg/hardware"
)

// Struct for the execute unit
type ExecuteUnit struct {
	hw *hardware.Hardware // The hardware component
	alu *Alu
}

// Function that creates the execute unit
func NewExecuteUnit() *ExecuteUnit {
	exu := ExecuteUnit {
		hw: hardware.NewHardware("EXU", 0),
		alu: NewAlu(),
	}
	return &exu
}

// Function to execute the given instruction
func (exu *ExecuteUnit) ExecuteInstruction(idexReg *IDEXReg) *EXMEMReg {
	opcode := idexReg.instr >> 21

	switch opcode {
	case 0x694, 0x695, 0x696, 0x697: // MOVZ
		// Get the shift amount
		shiftAmt := idexReg.instr & 0b11
		actualShiftAmt := exu.alu.Multiply(uint64(shiftAmt), 0x10)[1]
		
		// Compute the new register value
		newReg := idexReg.signExtendImm << actualShiftAmt

		return &EXMEMReg {
			instr: idexReg.instr,
			writeVal: newReg,
		}
	}
	
	return nil
}

// Logs a message
func (exu *ExecuteUnit) Log(msg string) {
	exu.hw.Log(msg)
}