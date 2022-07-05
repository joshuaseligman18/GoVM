package assembler

import (
	"os"
	"log"
	"bufio"
	"strings"
	"strconv"
	"fmt"
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
		case "MOVZ":
			instrBin = instrIM(opcode, operands, filePath, instrIndex + 1)
		}

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
		outBin  = outBin << 2 | uint32(shiftInt / 16)
		fmt.Printf("%x\n", outBin)
	} else {
		errMsg := fmt.Sprintf("Bad shift value: File: %s; Line: %d", fileName, lineNumber)
		log.Fatal(errMsg)
	}

	return outBin
}