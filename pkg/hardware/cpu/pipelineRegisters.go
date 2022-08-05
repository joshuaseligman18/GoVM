package cpu

// Struct for the data passed from the fetch unit to the decode unit
type IFIDReg struct {
	Instr         uint32 `json:"instr"`         // The fetched instruction
	IncrementedPC uint64 `json:"incrementedPC"` // The incremented program counter
}

// Gets the instruction in the IFID register
func (ifidReg *IFIDReg) GetInstruction() uint32 {
	return ifidReg.Instr
}

// Gets the incremented program counter in the IFID register
func (ifidReg *IFIDReg) GetIncrementedPC() uint64 {
	return ifidReg.IncrementedPC
}

// Struct for the data passed from the decode unit to the execute unit
type IDEXReg struct {
	Instr         uint32 `json:"instr"`         // The fetched instruction
	IncrementedPC uint64 `json:"incrementedPC"` // The incremented program counter
	RegReadData1  uint64 `json:"regReadData1"`  // The data in the first read register
	RegReadData2  uint64 `json:"regReadData2"`  // The data in the second read register
	SignExtendImm uint64 `json:"signExtendImm"` // The sign extended immediate
	addedLock     bool   // If the lock was added
}

// Gets the instruction in the IDEX register
func (idexReg *IDEXReg) GetInstruction() uint32 {
	return idexReg.Instr
}

// Gets the incremented program counter in the IDEX register
func (idexReg *IDEXReg) GetIncrementedPC() uint64 {
	return idexReg.IncrementedPC
}

// Gets the incremented program counter in the IDEX register
func (idexReg *IDEXReg) GetRegReadData1() uint64 {
	return idexReg.RegReadData1
}

// Gets the incremented program counter in the IDEX register
func (idexReg *IDEXReg) GetRegReadData2() uint64 {
	return idexReg.RegReadData2
}

// Gets the incremented program counter in the IDEX register
func (idexReg *IDEXReg) GetSignExtendedImmediate() uint64 {
	return idexReg.SignExtendImm
}

// Gets if the lock was added
func (idexReg *IDEXReg) GetAddedLock() bool {
	return idexReg.addedLock
}

// Struct for the data passed between the execute and memory units
type EXMEMReg struct {
	Instr         uint32 `json:"instr"`         // The instruction
	IncrementedPC uint64 `json:"incrementedPC"` // The incremented program counter
	WriteVal      uint64 `json:"writeVal"`      // The value to write
	WorkingAddr   uint64 `json:"workingAddr"`   // The address to work with
}

// Gets the instruction in the IDEX register
func (exmemReg *EXMEMReg) GetInstruction() uint32 {
	return exmemReg.Instr
}

// Gets the incremented program counter in the IDEX register
func (exmemReg *EXMEMReg) GetIncrementedPC() uint64 {
	return exmemReg.IncrementedPC
}

// Gets the value to be written
func (exmemReg *EXMEMReg) GetWriteVal() uint64 {
	return exmemReg.WriteVal
}

// Gets the memory address to work with
func (exmemReg *EXMEMReg) GetWorkingAddr() uint64 {
	return exmemReg.WorkingAddr
}

// Struct for the data passed between the execute and memory units
type MEMWBReg struct {
	Instr         uint32 `json:"instr"`         // The instruction
	IncrementedPC uint64 `json:"incrementedPC"` // The incremented program counter
	WriteVal      uint64 `json:"writeVal"`      // The value to write
}

// Gets the instruction in the IDEX register
func (memwbReg *MEMWBReg) GetInstruction() uint32 {
	return memwbReg.Instr
}

// Gets the incremented program counter in the IDEX register
func (memwbReg *MEMWBReg) GetIncrementedPC() uint64 {
	return memwbReg.IncrementedPC
}

// Gets the value to be written
func (memwbReg *MEMWBReg) GetWriteVal() uint64 {
	return memwbReg.WriteVal
}
