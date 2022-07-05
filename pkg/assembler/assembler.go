package assembler

import (
	"os"
	"log"
	"bufio"
	"fmt"
)

// Assembles a program into instructions for the computer to read
func AssembleProgram(filePath string, maxSize int) []uint32 {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	program := make([]uint32, maxSize)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	
	return program
}