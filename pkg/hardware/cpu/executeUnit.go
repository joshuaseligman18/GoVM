package cpu

import (
	"fmt"

	"github.com/joshuaseligman/GoVM/pkg/hardware"
	"github.com/joshuaseligman/GoVM/pkg/hardware/clock"
	"github.com/joshuaseligman/GoVM/pkg/util"
)

// Struct for the execute unit
type ExecuteUnit struct {
	hw       *hardware.Hardware // The hardware component
	cpu      *Cpu               // The CPU
	alu      *Alu               // The ALU
	flushing bool               // Variable to determine if the next instruction can be taken in
	clk      *clock.Clock       // The clock being used
}

// Function that creates the execute unit
func NewExecuteUnit(cpuPtr *Cpu, clk *clock.Clock) *ExecuteUnit {
	exu := ExecuteUnit{
		hw:       hardware.NewHardware("EXU", 0),
		cpu:      cpuPtr,
		alu:      NewAlu(),
		flushing: false,
		clk:      clk,
	}
	return &exu
}

// Function to execute the given instruction
func (exu *ExecuteUnit) ExecuteInstruction(out chan *EXMEMReg, idexReg *IDEXReg, memRunning *bool, wbRunning *bool) {
	opcode := idexReg.Instr >> 21

	switch opcode {
	case 0x694, 0x695, 0x696, 0x697: // MOVZ
		// Get the shift amount
		shiftAmt := opcode & 0b11
		actualShiftAmt := exu.alu.Multiply(uint64(shiftAmt), 0x10)[1]
		exu.alu.ClearFlags()

		// Compute the new register value
		newReg := idexReg.SignExtendImm << actualShiftAmt

		out <- &EXMEMReg{
			Instr:         idexReg.Instr,
			IncrementedPC: idexReg.IncrementedPC,
			WriteVal:      newReg,
		}

	case 0x794, 0x795, 0x796, 0x797: // MOVK
		// Get the shift amount
		shiftAmt := opcode & 0b11
		actualShiftAmt := exu.alu.Multiply(uint64(shiftAmt), 0x10)[1]
		exu.alu.ClearFlags()

		exu.Log(fmt.Sprintf("Old register: %s", util.ConvertToHexUint64(idexReg.RegReadData1)))

		// Compute the new register value
		newReg := idexReg.RegReadData1 >> (actualShiftAmt + 16)
		newReg = newReg<<16 | idexReg.SignExtendImm
		newReg = (newReg << actualShiftAmt) | (idexReg.RegReadData1 & (movkBitAndAmountUtil(actualShiftAmt)))
		exu.Log(fmt.Sprintf("New register: %s", util.ConvertToHexUint64(newReg)))

		out <- &EXMEMReg{
			Instr:         idexReg.Instr,
			IncrementedPC: idexReg.IncrementedPC,
			WriteVal:      newReg,
		}

	case 0x458, 0x558: // ADD, ADDS
		output := exu.alu.Add(idexReg.RegReadData1, idexReg.RegReadData2)

		// Clear flags if ADD
		if opcode == 0x458 {
			exu.alu.ClearFlags()
		}

		exu.Log(fmt.Sprintf("Sum: %s", util.ConvertToHexUint64(output)))

		out <- &EXMEMReg{
			Instr:         idexReg.Instr,
			IncrementedPC: idexReg.IncrementedPC,
			WriteVal:      output,
		}

	case 0x488, 0x489, // ADDI
		0x588, 0x589: // ADDIS
		output := exu.alu.Add(idexReg.RegReadData1, idexReg.SignExtendImm)

		// Clear flags if ADDI
		if opcode == 0x488 || opcode == 0x489 {
			exu.alu.ClearFlags()
		}

		exu.Log(fmt.Sprintf("Sum: %s", util.ConvertToHexUint64(output)))

		out <- &EXMEMReg{
			Instr:         idexReg.Instr,
			IncrementedPC: idexReg.IncrementedPC,
			WriteVal:      output,
		}

	case 0x658, 0x758: // SUB, SUBS
		output := exu.alu.Add(idexReg.RegReadData1, exu.alu.Negate(idexReg.RegReadData2))

		// Clear the flags if it is SUB
		if opcode == 0x658 {
			exu.alu.ClearFlags()
		}

		fmt.Println(exu.alu)

		exu.Log(fmt.Sprintf("Difference: %s", util.ConvertToHexUint64(output)))

		out <- &EXMEMReg{
			Instr:         idexReg.Instr,
			IncrementedPC: idexReg.IncrementedPC,
			WriteVal:      output,
		}

	case 0x688, 0x689, // SUBI
		0x788, 0x789: // SUBIS
		output := exu.alu.Add(idexReg.RegReadData1, exu.alu.Negate(idexReg.SignExtendImm))

		// Clear flags if SUBI
		if opcode == 688 || opcode == 689 {
			exu.alu.ClearFlags()
		}

		exu.Log(fmt.Sprintf("Difference: %s", util.ConvertToHexUint64(output)))

		out <- &EXMEMReg{
			Instr:         idexReg.Instr,
			IncrementedPC: idexReg.IncrementedPC,
			WriteVal:      output,
		}

	case 0x7C2, 0x1C2, // LDUR, LDURB
		0x3C2, 0x5C4: // LDURH, LDURSW
		// Get the address to load from
		output := exu.alu.Add(idexReg.RegReadData1, idexReg.SignExtendImm)
		exu.alu.ClearFlags()

		out <- &EXMEMReg{
			Instr:         idexReg.Instr,
			IncrementedPC: idexReg.IncrementedPC,
			WorkingAddr:   output,
		}

	case 0x7C0, 0x1C0, // STUR, STURB
		0x3C0, 0x5C0: // STURH, STURW
		// Get the address to load from
		output := exu.alu.Add(idexReg.RegReadData1, idexReg.SignExtendImm)
		exu.alu.ClearFlags()

		out <- &EXMEMReg{
			Instr:         idexReg.Instr,
			IncrementedPC: idexReg.IncrementedPC,
			WorkingAddr:   output,
			WriteVal:      idexReg.RegReadData2,
		}

	case 0x000: // HLT
		for *memRunning || *wbRunning {
			continue
		}
		exu.Log("Halting program")
		exu.clk.StopClock()
	}

	if opcode >= 0x0A0 && opcode <= 0x0BF { // B
		// Get the new program counter
		offset := idexReg.SignExtendImm << 2
		newPC := exu.alu.Add(idexReg.IncrementedPC, offset)
		exu.alu.ClearFlags()
		exu.Log(util.ConvertToHexUint64(newPC))
		// Flush the pipeline
		exu.flushing = true
		go exu.cpu.FlushPipeline(newPC)

		// Continue execution of the branch instruction
		out <- &EXMEMReg{
			Instr:         idexReg.Instr,
			IncrementedPC: idexReg.IncrementedPC,
		}
	}

	if opcode >= 0x5A0 && opcode <= 0x5A7 || // CBZ
		opcode >= 0x5A8 && opcode <= 0x5AF { // CBNZ
		exu.alu.ClearFlags()
		exu.alu.Add(idexReg.RegReadData1, 0)
		if opcode >= 0x5A0 && opcode <= 0x5A7 && exu.alu.zeroFlag || // CBZ condition
			opcode >= 0x5A8 && opcode <= 0x5AF && !exu.alu.zeroFlag { // CBNZ condition
			offSet := idexReg.SignExtendImm << 2
			exu.alu.ClearFlags()
			newPC := exu.alu.Add(idexReg.IncrementedPC, offSet)
			exu.flushing = true
			go exu.cpu.FlushPipeline(newPC)
		}
		exu.alu.ClearFlags()
		// Continue execution of the branch instruction
		out <- &EXMEMReg{
			Instr:         idexReg.Instr,
			IncrementedPC: idexReg.IncrementedPC,
		}
	}
}

// Logs a message
func (exu *ExecuteUnit) Log(msg string) {
	exu.hw.Log(msg)
}

// Util for determining how many bits to AND from the original value
func movkBitAndAmountUtil(actShiftAmt uint64) uint64 {
	if actShiftAmt == 0 {
		return 0
	}
	sum := uint64(0)
	for i := 0; i < int(actShiftAmt)/4; i++ {
		sum = sum<<4 | 0xF
	}
	return sum
}

// Resets the execute unit
func (exu *ExecuteUnit) Reset() {
	exu.alu.Reset()
	exu.flushing = false
}
