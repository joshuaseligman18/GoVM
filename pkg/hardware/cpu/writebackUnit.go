package cpu

import "github.com/joshuaseligman/GoVM/pkg/hardware"

// Struct for the writeback unit
type WritebackUnit struct {
	hw hardware.Hardware // The hardware component
	cpu *Cpu
}

// Function that creates the writeback unit
func NewWritebackUnit(parentCpu *Cpu) *WritebackUnit {
	wbu := WritebackUnit {
		hw: *hardware.NewHardware("WBU", 0),
		cpu: parentCpu,
	}
	return &wbu
}

// Function for handling the writeback
func (wbu *WritebackUnit) HandleWriteback(out chan bool, memwbReg *MEMWBReg) {
	opcode := memwbReg.instr >> 21

	switch opcode {
	case 0x694, 0x695, 0x696, 0x697, // MOVZ
		 0x794, 0x795, 0x796, 0x797, // MOVK
		 0x458, 0x558, // ADD, ADDS
		 0x488, 0x489, // ADDI
		 0x588, 0x589, // ADDIS
		 0x658, 0x758, // SUB, SUBS
		 0x688, 0x689, // SUBI
		 0x788, 0x789: // SUBIS
		reg := wbu.cpu.GetRegisterLocks().Dequeue()
		wbu.cpu.GetRegisters()[reg] = memwbReg.writeVal
	}

	out <- true
}