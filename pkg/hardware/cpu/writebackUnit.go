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
	case 0x7C0, 0x1C0, 0x3C0, 0x5C0: // STUR, STURB, STURH, STURW
		break
	default:
		reg := wbu.cpu.GetRegisterLocks().Dequeue()
		wbu.cpu.GetRegisters()[reg] = memwbReg.writeVal
	}

	out <- true
}