package cpu

import "github.com/joshuaseligman/GoVM/pkg/hardware"

// Struct for the ALU
type Alu struct {
	hw *hardware.Hardware // Hardware component
	negativeFlag bool // Flag for a negative output
	zeroFlag bool // Flag for a zero output
	overflowFlag bool // Flag for an overflow
	carryFlag bool // Flag for a carry
}

// Function that creates a new ALU
func NewAlu() *Alu {
	alu := Alu { 
		hw: hardware.NewHardware("ALU", 0),
		negativeFlag: false,
		zeroFlag: false,
		overflowFlag: false,
		carryFlag: false, 
	}
	return &alu
}