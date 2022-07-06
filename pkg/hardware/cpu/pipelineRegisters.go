package cpu

// Struct for the data passed from the fetch unit to the decode unit
type IFIDReg struct {
	instr uint32 // The fetched instruction
	incrementedPC uint // The incremented program counter
}

// Function to create the IFID register
func NewIFIDReg(instruction uint32, pc uint) *IFIDReg {
	reg := IFIDReg {
		instr: instruction,
		incrementedPC: pc,
	}
	return &reg
}