package cpu

import (
	"github.com/joshuaseligman/GoVM/pkg/hardware"
	"github.com/joshuaseligman/GoVM/pkg/hardware/memory"
	"github.com/joshuaseligman/GoVM/pkg/util"
)

// Struct for the CPU
type Cpu struct {
	hw *hardware.Hardware // The hardware struct
	acc uint64 // The accumulator
	reg []uint64 // Other registers
	programCounter uint // The address of the current instruction being fetched
	fetchUnit *FetchUnit // The fetch unit
	decodeUnit *DecodeUnit // The decode unit
	executeUnit *ExecuteUnit // The execute unit
	memDataUnit *MemDataUnit // The memory data unit
	writebackUnit *WritebackUnit // The writeback unit
	ifidReg *IFIDReg // The register between the fetch and decode units
	idexReg *IDEXReg // The register between the decode and execute units
	exmemReg *EXMEMReg // The register between the execute and memory data units
	memwbReg *MEMWBReg // The register between the memory data and writeback units

	regLocks *util.Queue // Manages locks for reading and writing to registers
}

// Creates the CPU
func NewCpu(mem *memory.Memory) *Cpu {
	cpu := Cpu {
		hw: hardware.NewHardware("CPU", 0),
		acc: 0x0,
		reg: make([]uint64, 32),
		programCounter: 0,
		fetchUnit: NewFetchUnit(mem),
		executeUnit: NewExecuteUnit(),
		memDataUnit: NewMemDataUnit(),
		ifidReg: NewIFIDReg(0, 0),
		regLocks: util.NewQueue(),
	}
	cpu.decodeUnit = NewDecodeUnit(&cpu)
	cpu.writebackUnit = NewWritebackUnit(&cpu)
	return &cpu
}

// Function that gets called every clock cycle
func (cpu *Cpu) Pulse() {
	fetchOut := cpu.fetchUnit.FetchInstruction(&cpu.programCounter)
	cpu.ifidReg = fetchOut
	decodeOut := cpu.decodeUnit.DecodeInstruction(cpu.ifidReg)
	cpu.idexReg = decodeOut
	executeOut := cpu.executeUnit.ExecuteInstruction(cpu.idexReg)
	cpu.exmemReg = executeOut
	memOut := cpu.memDataUnit.HandleMemoryAccess(cpu.exmemReg)
	cpu.memwbReg = memOut
	cpu.writebackUnit.HandleWriteback(cpu.memwbReg)
}

// Logs a message
func (cpu *Cpu) Log(msg string) {
	cpu.hw.Log(msg)
}

// Gets the program counter
func (cpu *Cpu) GetProgramCounter() uint{
	return cpu.programCounter
}

// Gets the accumulator
func (cpu *Cpu) GetAcc() uint64 {
	return cpu.acc
}

// Gets the registers
func (cpu *Cpu) GetRegisters() []uint64 {
	return cpu.reg
}

// Gets the IFID register
func (cpu *Cpu) GetIFIDReg() *IFIDReg {
	return cpu.ifidReg
}

// Gets the IDEX register
func (cpu *Cpu) GetIDEXReg() *IDEXReg {
	return cpu.idexReg
}

// Gets the EXMEM register
func (cpu *Cpu) GetEXMEMReg() *EXMEMReg {
	return cpu.exmemReg
}

// Gets the MEMWB register
func (cpu *Cpu) GetMEMWBReg() *MEMWBReg {
	return cpu.memwbReg
}

// Gets the register locks queue
func (cpu *Cpu) GetRegisterLocks() *util.Queue {
	return cpu.regLocks
}