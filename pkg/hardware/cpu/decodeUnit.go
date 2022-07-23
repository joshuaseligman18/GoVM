package cpu

import (
	"fmt"
	"log"
	"sync"

	"github.com/joshuaseligman/GoVM/pkg/hardware"
	"github.com/joshuaseligman/GoVM/pkg/util"
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
func (idu *DecodeUnit) DecodeInstruction(out chan *IDEXReg, ifidReg *IFIDReg, flushChan chan bool, flushWg *sync.WaitGroup) {
	opcode := ifidReg.instr >> 21
	idu.Log(fmt.Sprintf("%X", opcode))
	switch opcode {
	// IM instructions
	case 0x694, 0x695, 0x696, 0x697: // MOVZ
		// Register to write to
		regWrite := ifidReg.instr & 0x1F
		if regWrite == 0x1F {
			log.Fatalf("Bad regsiter write; cannot write to register XZR; PC: %d", ifidReg.incrementedPC - 4)
		}

		// Immediate to write
		immediate := ifidReg.instr & 0x1FFFFF >> 5

		// Flush if there is a signal
		if len(flushChan) > 0 {
			idu.Log("Flushing")
			flushWg.Done()
		} else {
			idu.cpu.GetRegisterLocks().Enqueue(regWrite)
			out <- &IDEXReg {
				instr: ifidReg.instr,
				incrementedPC: ifidReg.incrementedPC,
				regReadData1: 0,
				regReadData2: 0,
				signExtendImm: util.SignExtend(immediate, 32), // Should always be positive
			}
		}

	case 0x794, 0x795, 0x796, 0x797: // MOVK
		// Register to write to
		regWrite := ifidReg.instr & 0x1F
		if regWrite == 0x1F {
			log.Fatalf("Bad regsiter write; cannot write to register XZR; PC: %d", ifidReg.incrementedPC - 4)
		}

		// Wait until the updated value is written
		for idu.cpu.GetRegisterLocks().Contains(regWrite) {
			continue
		}
		// Register to read from
		regReadData1 := idu.cpu.GetRegisters()[regWrite]
		// Immediate to write
		immediate := ifidReg.instr & 0x1FFFFF >> 5
		
		// Flush if there is a signal
		if len(flushChan) > 0 {
			idu.Log("Flushing")
			flushWg.Done()
		} else {
			idu.cpu.GetRegisterLocks().Enqueue(regWrite)
			out <- &IDEXReg {
				instr: ifidReg.instr,
				incrementedPC: ifidReg.incrementedPC,
				regReadData1: regReadData1,
				regReadData2: 0,
				signExtendImm: util.SignExtend(immediate, 32), // Should always be positive
			}
		}

	case 0x458, 0x558, 0x658, 0x758: // ADD, ADDS, SUB
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
		if regWrite == 0x1F {
			log.Fatalf("Bad regsiter write; cannot write to register XZR; PC: %d", ifidReg.incrementedPC - 4)
		}

		// Flush if there is a signal
		if len(flushChan) > 0 {
			idu.Log("Flushing")
			flushWg.Done()
		} else {
			idu.cpu.GetRegisterLocks().Enqueue(regWrite)
			out <- &IDEXReg {
				instr: ifidReg.instr,
				incrementedPC: ifidReg.incrementedPC,
				regReadData1: regData1,
				regReadData2: regData2,
				signExtendImm: 0,
			}
		}

	case 0x488, 0x489, // ADDI
		 0x588, 0x589, // ADDIS
		 0x688, 0x689, // SUBI
		 0x788, 0x789: // SUBIS
		// Get the immediate value
		immediate := ifidReg.instr & 0x3FFFFF >> 10
		signExtendImm := util.SignExtend(immediate, 32)

		// Get the most updated read value
		regRead := ifidReg.instr & 0x3FF >> 5
		for idu.cpu.GetRegisterLocks().Contains(regRead) {
			continue
		}
		regData1 := idu.cpu.GetRegisters()[regRead]

		// Add the destination register to the queue
		regWrite := ifidReg.instr & 0x1F
		if regWrite == 0x1F {
			log.Fatalf("Bad regsiter write; cannot write to register XZR; PC: %d", ifidReg.incrementedPC - 4)
		}

		// Flush if there is a signal
		if len(flushChan) > 0 {
			idu.Log("Flushing")
			flushWg.Done()
		} else {
			idu.cpu.GetRegisterLocks().Enqueue(regWrite)
			out <- &IDEXReg {
				instr: ifidReg.instr,
				incrementedPC: ifidReg.incrementedPC,
				regReadData1: regData1,
				regReadData2: 0,
				signExtendImm: signExtendImm,
			}
		}

	case 0x7C2, 0x1C2, // LDUR, LDURB
		 0x3C2, 0x5C4: // LDURH, LDURSW
		// Get the immediate value
		immediate := ifidReg.instr & 0x1FFFFF >> 12
		signExtendImm := util.SignExtend(immediate, 8)

		// Get the most updated value to work with
		regRead := ifidReg.instr & 0x3FF >> 5
		for idu.cpu.GetRegisterLocks().Contains(regRead) {
			continue
		}
		regData1 := idu.cpu.GetRegisters()[regRead]

		// Add the destination register to the queue
		regWrite := ifidReg.instr & 0x1F
		if regWrite == 0x1F {
			log.Fatalf("Bad regsiter write; cannot write to register XZR; PC: %d", ifidReg.incrementedPC - 4)
		}

		// Flush if there is a signal
		if len(flushChan) > 0 {
			idu.Log("Flushing")
			flushWg.Done()
		} else {
			idu.cpu.GetRegisterLocks().Enqueue(regWrite)
			out <- &IDEXReg {
				instr: ifidReg.instr,
				incrementedPC: ifidReg.incrementedPC,
				regReadData1: regData1,
				regReadData2: 0,
				signExtendImm: signExtendImm,
			}
		}
	
	case 0x7C0, 0x1C0, // STUR, STURB
		 0x3C0, 0x5C0: // STURH, STURW
		// Get the immediate value
		immediate := ifidReg.instr & 0x1FFFFF >> 12
		signExtendImm := util.SignExtend(immediate, 8)

		// Get the most updated values to work with
		regRead1 := ifidReg.instr & 0x3FF >> 5
		regRead2 := ifidReg.instr & 0x1F
		for idu.cpu.GetRegisterLocks().Contains(regRead1) || idu.cpu.GetRegisterLocks().Contains(regRead2){
			continue
		}
		regData1 := idu.cpu.GetRegisters()[regRead1]
		regData2 := idu.cpu.GetRegisters()[regRead2]

		// Flush if there is a signal
		if len(flushChan) > 0 {
			idu.Log("Flushing")
			flushWg.Done()
		} else {
			out <- &IDEXReg {
				instr: ifidReg.instr,
				incrementedPC: ifidReg.incrementedPC,
				regReadData1: regData1,
				regReadData2: regData2,
				signExtendImm: signExtendImm,
			}
		}
	}
	// Branch instructions
	if opcode >= 0x0A0 && opcode <= 0x0BF { // B
		branchAddr := ifidReg.instr & 0x3FFFFFF
		signExtendBranchAddr := util.SignExtend(branchAddr, 26)

		// Flush if there is a signal
		if len(flushChan) > 0 {
			idu.Log("Flushing")
			flushWg.Done()
		} else {
			out <- &IDEXReg {
				instr: ifidReg.instr,
				incrementedPC: ifidReg.incrementedPC,
				signExtendImm: signExtendBranchAddr,
			}
		}
	}
}

// Logs a message
func (idu *DecodeUnit) Log(msg string) {
	idu.hw.Log(msg)
}