package cpu

import (
	"fmt"

	"github.com/joshuaseligman/GoVM/pkg/hardware"
)

// Struct for the decode unit
type DecodeUnit struct {
	hw *hardware.Hardware
}

// Function that creates the decode unit
func NewDecodeUnit() *DecodeUnit {
	decodeUnit := DecodeUnit { hw: hardware.NewHardware("IDU", 0) }
	return &decodeUnit
}

// Function that decodes an instruction into its operands
func (idu *DecodeUnit) DecodeInstruction(ifidReg *IFIDReg) {
	first9Bits := ifidReg.instr >> 21
	idu.Log(fmt.Sprintf("%X", first9Bits))
	switch first9Bits {
	case 0x694, 0x695, 0x696, 0x697, // MOVZ
		 0x794, 0x795, 0x796, 0x797: // MOVK
		// Do something
	}
}

// Logs a message
func (idu *DecodeUnit) Log(msg string) {
	idu.hw.Log(msg)
}