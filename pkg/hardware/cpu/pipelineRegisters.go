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

// Gets the instruction in the IFID register
func (ifidReg *IFIDReg) GetInstruction() uint32 {
	return ifidReg.instr
}

// Gets the incremented program counter in the IFID register
func (ifidReg *IFIDReg) GetIncrementedPC() uint {
	return ifidReg.incrementedPC
}

// Struct for the data passed from the decode unit to the execute unit
type IDEXReg struct {
	instr uint32 // The fetched instruction
	incrementedPC uint // The incremented program counter
	regReadData1 uint64 // The data in the first read register
	regReadData2 uint64 // The data in the second read register
	signExtendImm uint64 // The sign extended immediate
}

// Function to create the IDEX register
func NewIDEXReg(instruction uint32, pc uint, reg1 uint64, reg2, uint64, imm uint64) *IDEXReg {
	reg := IDEXReg {
		instr: instruction,
		incrementedPC: pc,
		regReadData1: reg1,
		regReadData2: reg2,
		signExtendImm: imm,
	}
	return &reg
}

// Gets the instruction in the IDEX register
func (idexReg *IDEXReg) GetInstruction() uint32 {
	return idexReg.instr
}

// Gets the incremented program counter in the IDEX register
func (idexReg *IDEXReg) GetIncrementedPC() uint {
	return idexReg.incrementedPC
}

// Gets the incremented program counter in the IDEX register
func (idexReg *IDEXReg) GetRegReadData1() uint64 {
	return idexReg.regReadData1
}

// Gets the incremented program counter in the IDEX register
func (idexReg *IDEXReg) GetRegReadData2() uint64 {
	return idexReg.regReadData2
}

// Gets the incremented program counter in the IDEX register
func (idexReg *IDEXReg) GetSignExtendedImmediate() uint64 {
	return idexReg.signExtendImm
}

// Struct for the data passed between the execute and memory units
type EXMEMReg struct {
	instr uint32 // The instruction
	writeVal uint64 // The value to write
}

// Gets the instruction in the IDEX register
func (exmemReg *EXMEMReg) GetInstruction() uint32 {
	return exmemReg.instr
}

// Gets the value to be written
func (exmemReg *EXMEMReg) GetWriteVal() uint64 {
	return exmemReg.writeVal
}

// Struct for the data passed between the execute and memory units
type MEMWBReg struct {
	instr uint32 // The instruction
	writeVal uint64 // The value to write
}

// Gets the instruction in the IDEX register
func (memwbReg *MEMWBReg) GetInstruction() uint32 {
	return memwbReg.instr
}

// Gets the value to be written
func (memwbReg *MEMWBReg) GetWriteVal() uint64 {
	return memwbReg.writeVal
}