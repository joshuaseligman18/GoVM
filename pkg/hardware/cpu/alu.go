package cpu

import "github.com/joshuaseligman/GoVM/pkg/hardware"

// Struct for the ALU
type Alu struct {
	hw *hardware.Hardware // Hardware component
	negativeFlag bool // Flag for a negative output
	zeroFlag bool // Flag for a zero output
	overflowFlag bool // Flag for an overflow
	carryFlag bool // Flag for a carry
}

// Struct for managing the adder
type adderOutput struct {
	sum uint64 // The sum bit
	carry uint64 // The carry bit
}

// Function that creates a new ALU
func NewAlu() *Alu {
	alu := Alu { 
		hw: hardware.NewHardware("ALU", 0),
		negativeFlag: false,
		zeroFlag: false,
		overflowFlag: false,
		carryFlag: false, 
	}
	return &alu
}

// Function that adds 2 numbers
func (alu *Alu) Add(num1 uint64, num2 uint64, useCarry bool) uint64 {
	// Initialize the sum and carry
	sum := uint64(0)
	carry := uint64(0)
	if useCarry && alu.carryFlag {
		carry = 1
	}

	for i := 0; i < 64; i++ {
		// Get the next bits to add
		bit1 := num1 & 0b1
		bit2 := num2 & 0b1

		num1 = num1 >> 1
		num2 = num2 >> 1

		// Add the numbers with the consideration of the carry
		result := alu.fullAdder(bit1, bit2, carry)
		
		// Update the sum and carry
		sum = result.sum << i | sum
		carry = result.carry
	}

	// Update the flags
	alu.zeroFlag = sum == 0
	if carry == 0 {
		alu.carryFlag = false
	} else {
		alu.carryFlag = true
	}

	return sum
}

// Function that performs the tasks of a full adder
func (alu *Alu) fullAdder(bit1 uint64, bit2 uint64, carry uint64) adderOutput {
	// Add the two original bits
	first := alu.halfAdder(bit1, bit2);
	
	// Add the first sum with the carry input
	second := alu.halfAdder(first.sum, carry);

	// Return the result
	sumOut := second.sum;
	carryOut := first.carry | second.carry;
	out := adderOutput {
		sum: sumOut,
		carry: carryOut,
	}
	return out
}

// Function that represents a half adder
func (alu *Alu) halfAdder(bit1 uint64, bit2 uint64) adderOutput {
	// Sum is XOR of the bits
	sumOut := bit1 ^ bit2;
	
	// Carry is the AND of the bits
	carryOut := bit1 & bit2;

	// Return the result
	out := adderOutput {
		sum: sumOut,
		carry: carryOut,
	}
	return out
}


// Function for multiplying 2 numbers
func (alu *Alu) Multiply(multiplicand uint64, multiplier uint64) []uint64 {
	multiplicandTop := uint64(0)

	productBottom := uint64(0)
	productTop := uint64(0)

	for i := 0; i < 64; i++ {
		lastMultiplierBit := multiplier & 0b1

		if lastMultiplierBit == 1 {
			// Add to product
			
		}

		go func () {
			multiplier = multiplier >> 1
			topMultiplicandBit := multiplicand >> 31
			multiplicand = multiplicand << 1
			multiplicandTop = multiplicandTop << 1 | topMultiplicandBit
		}()
	}
}