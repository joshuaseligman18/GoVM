package assembler

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"math"
)

// Assembles a program into instructions for the computer to read
func AssembleProgram(filePath string, maxSize int) []uint32 {
	// Open the file
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
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
		if opcodeSplit == -1 {
			log.Fatal("Invalid instruction ", instr)
		}

		// Get the opcode
		opcode := instr[:opcodeSplit]
		fmt.Println(opcode)

		// Get a list of operands
		operands := strings.Split(instr[opcodeSplit + 1:], ", ")
		fmt.Println(operands)

		instrBin := uint32(0)

		switch opcode {
		// IM instructions
		case "MOVZ", "MOVK":
			instrBin = instrIM(opcode, operands, filePath, instrIndex + 1)
		// R instructions
		case "ADD", "ADDS", "SUB", "SUBS":
			instrBin = instrR(opcode, operands, filePath, instrIndex + 1)
		// I instructions
		case "ADDI", "ADDIS", "SUBI":
			instrBin = instrI(opcode, operands, filePath, instrIndex + 1)
		}

		// Add the instruction to the program
		program[instrIndex] = instrBin
		instrIndex++
	}
	
	return program
}

// Generates the binary for IM instructions
func instrIM(opcode string, operands []string, fileName string, lineNumber int) uint32 {
	// Make sure we have the right number of operands
	if len(operands) != 3 {
		errMsg := fmt.Sprintf("Invalid instruction format: Expected 3 operands but got %d; File: %s; Line: %d", len(operands), fileName, lineNumber)
		log.Fatal(errMsg)
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
		log.Fatal(errMsg)
	}

	// Get the value to move into the register
	val := getValue(operands[1], 16, "move immediate", fileName, lineNumber)
	outBin = outBin << 16 | uint32(val)

	// Get the register to move the value to
	destReg := getRegister(operands[0], fileName, lineNumber)
	outBin = outBin << 5 | uint32(destReg)

	// Return the instruction binary
	return outBin
}

// Generates the binary for R instructions
func instrR(opcode string, operands []string, fileName string, lineNumber int) uint32 {
	// Make sure we have the right number of operands
	switch opcode {
	case "ADD", "ADDS", "SUB", "SUBS":
		if len(operands) != 3 {
			errMsg := fmt.Sprintf("Invalid instruction format: Expected 3 operands but got %d; File: %s; Line: %d", len(operands), fileName, lineNumber)
			log.Fatal(errMsg)
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
		readReg1 := getRegister(operands[1], fileName, lineNumber)
		outBin = outBin << 5 | uint32(readReg1)
		
		// Add an empty shift amount
		outBin = outBin << 6

		// Get the second register for the operation
		readReg2 := getRegister(operands[2], fileName, lineNumber)
		outBin = outBin << 5 | uint32(readReg2)

		// Get the destination register for the operation
		destReg := getRegister(operands[0], fileName, lineNumber)
		outBin = outBin << 5 | uint32(destReg)
	}

	// Return the instruction binary
	return outBin
}

// Generates the binary for I instructions
func instrI(opcode string, operands []string, fileName string, lineNumber int) uint32 {
	// Make sure we have the right number of operands
	switch opcode {
	case "ADDI", "ADDIS", "SUBI":
		if len(operands) != 3 {
			errMsg := fmt.Sprintf("Invalid instruction format: Expected 3 operands but got %d; File: %s; Line: %d", len(operands), fileName, lineNumber)
			log.Fatal(errMsg)
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
	}

	// Generate the remaining binary based on the instruction
	switch opcode {
	case "ADDI", "ADDIS", "SUBI":
		// Get the immediate value for adding
		val := getValue(operands[2], 12, "ALU immediate", fileName, lineNumber)
		outBin = outBin << 12 | uint32(val)

		// Get the register for the operation
		srcReg := getRegister(operands[1], fileName, lineNumber)
		outBin = outBin << 5 | uint32(srcReg)

		// Get the destination register for the operation
		destReg := getRegister(operands[0], fileName, lineNumber)
		outBin = outBin << 5 | uint32(destReg)
	}

	// Return the instruction binary
	return outBin
}

// Parses a string for a register value
func getRegister(regString string, fileName string, lineNumber int) int64 {
	// Get the register for the operation
	reg, errConv := strconv.ParseInt(regString[1:], 10, 0)
	if errConv != nil {
		// Bad value error
		errMsg := fmt.Sprintf("Bad register value; File: %s; Line: %d", fileName, lineNumber)
		log.Fatal(errMsg)
		return -1
	} else if reg < 0 || reg > 30 {
		// Invalid register error
		errMsg := fmt.Sprintf("Bad register value: Register must be between 0 and 30 (inclusive); File: %s; Line: %d", fileName, lineNumber)
		log.Fatal(errMsg)
		return -1
	} else {
		// Return the register
		return reg
	}
}

// Parses a string for a constant value
func getValue(valStr string, maxSize int, valName string, fileName string, lineNumber int) uint64 {
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
		return val
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
		log.Fatal(errMsg)
		return 0
	}
}