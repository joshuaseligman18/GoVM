package cpu

import (
	"fmt"

	"github.com/joshuaseligman/GoVM/pkg/hardware"
)

// Struct for the execute unit
type ExecuteUnit struct {
	hw *hardware.Hardware // The hardware component
}

// Function that creates the execute unit
func NewExecuteUnit() *ExecuteUnit {
	exu := ExecuteUnit { hw: hardware.NewHardware("EXU", 0) }
	return &exu
}

// Function to execute the given instruction
func (exu *ExecuteUnit) ExecuteInstruction(idexReg *IDEXReg) {
	opcode := idexReg.instr >> 21

	switch opcode {
	case 0x694, 0x695, 0x696, 0x697: // MOVZ
		shiftAmt := idexReg.instr & 0b11
		exu.Log(fmt.Sprintf("%d", shiftAmt))
	}
}

// Logs a message
func (exu *ExecuteUnit) Log(msg string) {
	exu.hw.Log(msg)
}