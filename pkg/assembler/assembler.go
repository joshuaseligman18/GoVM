package assembler

import (
	"os"
	"log"
	"bufio"
	"strings"
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

	// Read the file line by line
	for scanner.Scan() {
		instr := scanner.Text()

		// Get the end of the opcode
		opcodeSplit := strings.Index(instr, " ")
		if opcodeSplit == -1 {
			log.Fatal("Invalid instruction ", instr)
		}

		opcode := instr[:opcodeSplit]
		fmt.Println(opcode)

		operands := strings.Split(instr[opcodeSplit + 1:], ", ")
		fmt.Println(operands)
	}
	
	return program
}