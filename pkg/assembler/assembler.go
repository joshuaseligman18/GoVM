package assembler

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"math"
	"errors"
)

// Assembles a program into instructions for the computer to read
func AssembleProgramFile(filePath string, maxSize int) ([]uint32, error) {
	// Open the file
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	// Create the array that will become the memory
	program := make([]uint32, maxSize)
	instrIndex := 0

	// Read the file line by line
	for scanner.Scan() {
		instr := scanner.Text()

		// Get the end of the opcode
		opcodeSplit := strings.Index(instr, " ")

		// Define the variables for the opcode and operands
		var opcode string
		var operands []string

		if opcodeSplit == -1 && instr != "HLT" {
			errMsg := fmt.Sprint("Invalid instruction ", instr)
			return nil, errors.New(errMsg)
		} else if opcodeSplit != -1 && instr != "HLT" {
			// Get the opcode
			opcode = instr[:opcodeSplit]

			// Get a list of operands
			operands = strings.Split(instr[opcodeSplit + 1:], ", ")
		} else if instr == "HLT" {
			opcode = "HLT"
		}

		instrBin := uint32(0)
		var err error

		switch opcode {
		// IM instructions
		case "MOVZ", "MOVK":
			instrBin, err = instrIM(opcode, operands, filePath, instrIndex + 1)
		// R instructions
		case "ADD", "ADDS", "SUB", "SUBS":
			instrBin, err = instrR(opcode, operands, filePath, instrIndex + 1)
		// I instructions
		case "ADDI", "ADDIS", "SUBI", "SUBIS":
			instrBin, err = instrI(opcode, operands, filePath, instrIndex + 1)
		// D instructions
		case "LDUR", "LDURB", "LDURH", "LDURSW", "STUR", "STURB", "STURH", "STURW":
			instrBin, err = instrD(opcode, operands, filePath, instrIndex + 1)
		// B instructions
		case "B":
			instrBin, err = instrB(opcode, operands, filePath, instrIndex + 1)
		// CB instructions
		case "CBZ", "CBNZ":
			instrBin, err = instrCB(opcode, operands, filePath, instrIndex + 1)
		// Constant data
		case "DATA":
			instrBin, err = instrData(operands, filePath, instrIndex + 1)
		// Halt
		case "HLT":
			instrBin = 0
		default:
			errMsg := fmt.Sprintf("Invalid opcode; File: %s; Line: %d", filePath, instrIndex + 1)
			err = errors.New(errMsg)
		}

		if err != nil {
			return nil, err
		}

		// Add the instruction to the program
		program[instrIndex] = instrBin
		instrIndex++
	}
	
	return program, nil
}

// Assembles a program into instructions for the computer to read
func AssembleProgramAPI(progStr string) ([]uint32, error) {
	// Create the array that will become the memory
	program := make([]uint32, 0)
	filePath := "Web form"

	// Read the file line by line
	for instrIndex, instr := range strings.Split(progStr, "\n") {

		// Get the end of the opcode
		opcodeSplit := strings.Index(instr, " ")

		// Define the variables for the opcode and operands
		var opcode string
		var operands []string

		if opcodeSplit == -1 && instr != "HLT" {
			errMsg := fmt.Sprintf("Invalid instruction: %s; File: %s; Line: %d", instr, filePath, instrIndex + 1)
			return nil, errors.New(errMsg)
		} else if opcodeSplit != -1 && instr != "HLT" {
			// Get the opcode
			opcode = instr[:opcodeSplit]

			// Get a list of operands
			operands = strings.Split(instr[opcodeSplit + 1:], ", ")
		} else if instr == "HLT" {
			opcode = "HLT"
		}

		instrBin := uint32(0)
		var err error

		switch opcode {
		// IM instructions
		case "MOVZ", "MOVK":
			instrBin, err = instrIM(opcode, operands, filePath, instrIndex + 1)
		// R instructions
		case "ADD", "ADDS", "SUB", "SUBS":
			instrBin, err = instrR(opcode, operands, filePath, instrIndex + 1)
		// I instructions
		case "ADDI", "ADDIS", "SUBI", "SUBIS":
			instrBin, err = instrI(opcode, operands, filePath, instrIndex + 1)
		// D instructions
		case "LDUR", "LDURB", "LDURH", "LDURSW", "STUR", "STURB", "STURH", "STURW":
			instrBin, err = instrD(opcode, operands, filePath, instrIndex + 1)
		// B instructions
		case "B":
			instrBin, err = instrB(opcode, operands, filePath, instrIndex + 1)
		// CB instructions
		case "CBZ", "CBNZ":
			instrBin, err = instrCB(opcode, operands, filePath, instrIndex + 1)
		// Constant data
		case "DATA":
			instrBin, err = instrData(operands, filePath, instrIndex + 1)
		// Halt
		case "HLT":
			instrBin = 0
		default:
			errMsg := fmt.Sprintf("Invalid opcode; File: %s; Line: %d", filePath, instrIndex + 1)
			err = errors.New(errMsg)
		}

		if err != nil {
			return nil, err
		}

		// Add the instruction to the program
		program = append(program, instrBin)
	}
	
	return program, nil
}

// Generates the binary for IM instructions
func instrIM(opcode string, operands []string, fileName string, lineNumber int) (uint32, error) {
	// Make sure we have the right number of operands
	if len(operands) != 3 {
		errMsg := fmt.Sprintf("Invalid instruction format: Expected 3 operands but got %d; File: %s; Line: %d", len(operands), fileName, lineNumber)
		return 0, errors.New(errMsg)
	}

	outBin := uint32(0)

	// Generate initial binary
	switch opcode {
	case "MOVZ":
		outBin = 0b110100101
	case "MOVK":
		outBin = 0b111100101
	}

	// Get the shift amount
	shiftStr := strings.Split(operands[2], " ")[1]
	shiftInt, errConv := strconv.ParseInt(shiftStr, 10, 0)
	if errConv == nil {
		// Add the shift to the binary
		outBin  = outBin << 2 | uint32(shiftInt / 16)
	} else {
		// Bad shift value error
		errMsg := fmt.Sprintf("Bad shift value; File: %s; Line: %d", fileName, lineNumber)
		return 0, errors.New(errMsg)
	}

	// Get the value to move into the register
	val, err := getValue(operands[1], 16, "move immediate", fileName, lineNumber)
	if err == nil {
		outBin = outBin << 16 | uint32(val)
	} else {
		return 0, err
	}

	// Get the register to move the value to
	destReg, err := getRegister(operands[0], fileName, lineNumber)
	if err == nil {
		outBin = outBin << 5 | uint32(destReg)
	} else {
		return 0, err
	}

	// Return the instruction binary
	return outBin, nil
}

// Generates the binary for R instructions
func instrR(opcode string, operands []string, fileName string, lineNumber int) (uint32, error) {
	// Make sure we have the right number of operands
	switch opcode {
	case "ADD", "ADDS", "SUB", "SUBS":
		if len(operands) != 3 {
			errMsg := fmt.Sprintf("Invalid instruction format: Expected 3 operands but got %d; File: %s; Line: %d", len(operands), fileName, lineNumber)
			return 0, errors.New(errMsg)
		}
	}

	outBin := uint32(0)

	// Generate initial binary
	switch opcode {
	case "ADD":
		outBin = 0b10001011000
	case "ADDS":
		outBin = 0b10101011000
	case "SUB":
		outBin = 0b11001011000
	case "SUBS":
		outBin = 0b11101011000
	}

	// Generate the remaining binary based on the instruction
	switch opcode {
	case "ADD", "ADDS", "SUB", "SUBS":
		// Get the first register for the operation
		readReg1, err := getRegister(operands[1], fileName, lineNumber)
		if err == nil {
			outBin = outBin << 5 | uint32(readReg1)
		} else {
			return 0, err
		}
		
		// Add an empty shift amount
		outBin = outBin << 6

		// Get the second register for the operation
		readReg2, err := getRegister(operands[2], fileName, lineNumber)
		if err == nil {
			outBin = outBin << 5 | uint32(readReg2)
		} else {
			return 0, err
		}

		// Get the destination register for the operation
		destReg, err := getRegister(operands[0], fileName, lineNumber)
		if err == nil {
			outBin = outBin << 5 | uint32(destReg)
		} else {
			return 0, err
		}
	}

	// Return the instruction binary
	return outBin, nil
}

// Generates the binary for I instructions
func instrI(opcode string, operands []string, fileName string, lineNumber int) (uint32, error) {
	// Make sure we have the right number of operands
	switch opcode {
	case "ADDI", "ADDIS", "SUBI":
		if len(operands) != 3 {
			errMsg := fmt.Sprintf("Invalid instruction format: Expected 3 operands but got %d; File: %s; Line: %d", len(operands), fileName, lineNumber)
			return 0, errors.New(errMsg)
		}
	}

	outBin := uint32(0)

	// Generate initial binary
	switch opcode {
	case "ADDI":
		outBin = 0b1001000100
	case "ADDIS":
		outBin = 0b1011000100
	case "SUBI":
		outBin = 0b1101000100
	case "SUBIS":
		outBin = 0b1111000100
	}

	// Generate the remaining binary based on the instruction
	switch opcode {
	case "ADDI", "ADDIS", "SUBI", "SUBIS":
		// Get the immediate value for adding
		val, err := getValue(operands[2], 12, "ALU immediate", fileName, lineNumber)
		if err == nil {
			outBin = outBin << 12 | uint32(val)
		} else {
			return 0, err
		}

		// Get the register for the operation
		srcReg, err := getRegister(operands[1], fileName, lineNumber)
		if err == nil {
			outBin = outBin << 5 | uint32(srcReg)
		} else {
			return 0, err
		}

		// Get the destination register for the operation
		destReg, err := getRegister(operands[0], fileName, lineNumber)
		if err == nil {
			outBin = outBin << 5 | uint32(destReg)
		} else {
			return 0, err
		}
	}

	// Return the instruction binary
	return outBin, nil
}

// Generates the binary for D instructions
func instrD(opcode string, operands []string, fileName string, lineNumber int) (uint32, error) {
	// Make sure we have the right number of operands
	if len(operands) != 3 {
		errMsg := fmt.Sprintf("Invalid instruction format: Expected 3 operands but got %d; File: %s; Line: %d", len(operands), fileName, lineNumber)
		return 0, errors.New(errMsg)
	}

	outBin := uint32(0)

	// Generate initial binary
	switch opcode {
	case "LDUR":
		outBin = 0b11111000010
	case "LDURB":
		outBin = 0b00111000010
	case "LDURH":
		outBin = 0b01111000010
	case "LDURSW":
		outBin = 0b10111000100
	case "STUR":
		outBin = 0b11111000000
	case "STURB":
		outBin = 0b00111000000
	case "STURH":
		outBin = 0b01111000000
	case "STURW":
		outBin = 0b10111000000
	}

	// Get the immediate value for adding
	val, err := getValue(operands[2], 9, "destination address", fileName, lineNumber)
	if err == nil {
		outBin = outBin << 9 | uint32(val)
	} else {
		return 0, err
	}

	// Op is always 0
	outBin = outBin << 2

	// Get the register for the operation
	reg1, err := getRegister(operands[1], fileName, lineNumber)
	if err == nil {
		outBin = outBin << 5 | uint32(reg1)
	} else {
		return 0, err
	}

	// Get the destination register for the operation
	destReg, err := getRegister(operands[0], fileName, lineNumber)
	if err == nil {
		outBin = outBin << 5 | uint32(destReg)
	} else {
		return 0, err
	}

	// Return the instruction binary
	return outBin, nil
}

// Generates the binary for branch instructions
func instrB(opcode string, operands []string, fileName string, lineNumber int) (uint32, error) {
	// Make sure we only have 1 operand
	if len(operands) != 1 {
		errMsg := fmt.Sprintf("Invalid instruction format: Expected 1 operand but got %d; File: %s; Line: %d", len(operands), fileName, lineNumber)
		return 0, errors.New(errMsg)
	}

	outBin := uint32(0)

	// Generate initial binary
	switch opcode {
	case "B":
		outBin = 0b000101
	}

	// Get the relative branch address
	branchAddr, err := getValue(operands[0], 26, "branch address", fileName, lineNumber)
	if err == nil {
		outBin = outBin << 26 | uint32(branchAddr)
	} else {
		return 0, err
	}

	return outBin, nil
}

// Generates the binary for conditional branch instructions
func instrCB(opcode string, operands []string, fileName string, lineNumber int) (uint32, error) {
	// Make sure we only have 1 operand
	if len(operands) != 2 {
		errMsg := fmt.Sprintf("Invalid instruction format: Expected 2 operands but got %d; File: %s; Line: %d", len(operands), fileName, lineNumber)
		return 0, errors.New(errMsg)
	}

	outBin := uint32(0)

	// Generate initial binary
	switch opcode {
	case "CBZ":
		outBin = 0b10110100
	case "CBNZ":
		outBin = 0b10110101
	}

	// Get the branch address
	branchAddr, err := getValue(operands[1], 19, "branch address", fileName, lineNumber)
	if err == nil {
		outBin = outBin << 19 | uint32(branchAddr)
	} else {
		return 0, err
	}

	// Get the register for the condition
	register, err := getRegister(operands[0], fileName, lineNumber)
	if err == nil {
		outBin = outBin << 5 | uint32(register)
	} else {
		return 0, err
	}

	return outBin, nil
}

// Generates the binary for constant data
func instrData(operands []string, fileName string, lineNumber int) (uint32, error) {
	// Make sure we only have 1 number
	if len(operands) != 1 {
		errMsg := fmt.Sprintf("Invalid instruction format: Expected 3 operands but got %d; File: %s; Line: %d", len(operands), fileName, lineNumber)
		return 0, errors.New(errMsg)
	}

	// Get the value and return it
	val, err := getValue(operands[0], 32, "data", fileName, lineNumber)
	var outBin uint32
	if err == nil {
		outBin = uint32(val)
	} else {
		return 0, err
	}

	return outBin, nil
}

// Parses a string for a register value
func getRegister(regString string, fileName string, lineNumber int) (int64, error) {
	// Get the register for the operation
	reg, errConv := strconv.ParseInt(regString[1:], 10, 0)
	if errConv != nil {
		// Account for XZR register
		if regString[1:] == "ZR" {
			return 0x1F, nil
		}
		// Bad value error
		errMsg := fmt.Sprintf("Bad register value; File: %s; Line: %d", fileName, lineNumber)
		return -1, errors.New(errMsg)
	} else if reg < 0 || reg > 30 {
		// Invalid register error
		errMsg := fmt.Sprintf("Bad register value: Register must be between 0 and 30 (inclusive); File: %s; Line: %d", fileName, lineNumber)
		return -1, errors.New(errMsg)
	} else {
		// Return the register
		return reg, nil
	}
}

// Parses a string for a constant value
func getValue(valStr string, maxSize int, valName string, fileName string, lineNumber int) (uint64, error) {
	// Get the value to move into the register
	base := 0
	cut := 0
	if len(valStr) >= 4 && valStr[:3] == "#0x" {
		// Base 16
		base = 16
		cut = 3
	} else {
		// Base 10
		base = 10
		cut = 1
	}
	// Get the value based on the base that was decided earlier
	val, errConv := strconv.ParseUint(valStr[cut:], base, maxSize)
	if errConv == nil {
		return val, nil
	} else {
		errMsg := ""
		if strings.Contains(errConv.Error(), "value out of range") {
			maxValue := uint(math.Pow(2, float64(maxSize)) - 1)
			// Out of range errors
			if base == 10 {
				errMsg = fmt.Sprintf("Bad %s value: Value must be between 0 and %d (%d bits) but got %s; File: %s; Line: %d", valName, maxValue, maxSize, valStr, fileName, lineNumber)
			} else {
				repeatedZeros := strings.Repeat("0", int(math.Ceil(float64(maxSize) / 4)))
				errMsg = fmt.Sprintf("Bad %s value: Value must be between 0x%s and 0x%X (%d bits) but got %s; File: %s; Line: %d", valName, repeatedZeros, maxValue, maxSize, valStr, fileName, lineNumber)
			}
		} else {
			// Bad value error
			errMsg = fmt.Sprintf("Bad %s value; File: %s; Line: %d", valName, fileName, lineNumber)
		}
		return 0, errors.New(errMsg)
	}
}