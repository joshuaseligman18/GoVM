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
	ifidChan chan *IFIDReg // The channel to manage async communication between fetch and decode units
	idexChan chan *IDEXReg // The channel to manage async communication between decode and execute units
	exmemChan chan *EXMEMReg // The channel to manage async communication between execute and memory data units
	memwbChan chan *MEMWBReg // The channel to manage async communication between memory data and writeback units
	endInstrChan chan bool // The channel to manage async communication for the end of the instruction
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
		ifidChan: make(chan *IFIDReg, 1),
		idexChan: make(chan *IDEXReg, 1),
		exmemChan: make(chan *EXMEMReg, 1),
		memwbChan: make(chan *MEMWBReg, 1),
		endInstrChan: make(chan bool, 1),
		regLocks: util.NewQueue(),
	}
	cpu.decodeUnit = NewDecodeUnit(&cpu)
	cpu.writebackUnit = NewWritebackUnit(&cpu)
	return &cpu
}

// Function that gets called every clock cycle
func (cpu *Cpu) Pulse() {
	// Clear the writeback unit
	if len(cpu.endInstrChan) == 1 {
		<- cpu.endInstrChan
	}

	// Clear the mem data unit and writeback next instruction if available
	if len(cpu.endInstrChan) == 0 && len(cpu.memwbChan) == 1 {
		cpu.Log("Starting writeback")
		cpu.memwbReg = <- cpu.memwbChan
		go cpu.writebackUnit.HandleWriteback(cpu.endInstrChan, cpu.memwbReg)
	}

	// Clear the execute unit and handle memory if available
	if len(cpu.memwbChan) == 0 && len(cpu.exmemChan) == 1 {
		cpu.Log("Starting mem data access")
		cpu.exmemReg = <- cpu.exmemChan
		go cpu.memDataUnit.HandleMemoryAccess(cpu.memwbChan, cpu.exmemReg)
	}

	// Clear decode unit and execute if available
	if len(cpu.exmemChan) == 0 && len(cpu.idexChan) == 1 {
		cpu.Log("Starting execute")
		cpu.idexReg = <- cpu.idexChan
		go cpu.executeUnit.ExecuteInstruction(cpu.exmemChan, cpu.idexReg)
	}

	// Clear the fetch unit and decode if available
	if len(cpu.idexChan) == 0 && len(cpu.ifidChan) == 1 {
		cpu.Log("Starting decode")
		cpu.ifidReg = <- cpu.ifidChan
		go cpu.decodeUnit.DecodeInstruction(cpu.idexChan, cpu.ifidReg)
	}

	// Fetch if available
	if len(cpu.ifidChan) == 0 {
		cpu.Log("Starting fetch")
		go cpu.fetchUnit.FetchInstruction(cpu.ifidChan, &cpu.programCounter)
	}
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