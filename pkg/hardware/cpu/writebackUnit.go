package cpu

import (
	"fmt"

	"github.com/joshuaseligman/GoVM/pkg/hardware"
)

// Struct for the writeback unit
type WritebackUnit struct {
	hw  hardware.Hardware // The hardware component
	cpu *Cpu // The corresponding CPU
}

// Function that creates the writeback unit
func NewWritebackUnit(parentCpu *Cpu) *WritebackUnit {
	wbu := WritebackUnit{
		hw:  *hardware.NewHardware("WBU", 0),
		cpu: parentCpu,
	}
	return &wbu
}

// Function for handling the writeback
func (wbu *WritebackUnit) HandleWriteback(out chan bool, memwbReg *MEMWBReg) {
	opcode := memwbReg.Instr >> 21

	if opcode >= 0x0A0 && opcode <= 0x0BF || // B
		opcode >= 0x5A0 && opcode <= 0x5A7 || // CBZ
		opcode >= 0x5A8 && opcode <= 0x5AF { // CBNZ
		out <- true
		return
	}

	switch opcode {
	case 0x7C0, 0x1C0, 0x3C0, 0x5C0: // STUR, STURB, STURH, STURW
		break
	default:
		reg := wbu.cpu.GetRegisterLocks().Dequeue()
		wbu.cpu.GetRegisters()[reg] = memwbReg.WriteVal
		wbu.Log(fmt.Sprintf("Unlocked %d", reg))
		wbu.Log(wbu.cpu.GetRegisterLocks().ToString())
	}

	out <- true
}

// Logs a message
func (wbu *WritebackUnit) Log(msg string) {
	wbu.hw.Log(msg)
}
