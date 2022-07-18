package cpu

import (
	"fmt"

	"github.com/joshuaseligman/GoVM/pkg/hardware"
	"github.com/joshuaseligman/GoVM/pkg/hardware/memory"
	"github.com/joshuaseligman/GoVM/pkg/util"
)

// Struct for the CPU
type Cpu struct {
	hw *hardware.Hardware // The hardware struct
	acc uint64 // The accumulator
	reg []uint64 // Other registers
	programCounter uint64 // The address of the current instruction being fetched
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

var (
	ifidChan chan *IFIDReg = make(chan *IFIDReg, 1) // The channel to manage async communication between fetch and decode units
	idexChan chan *IDEXReg = make(chan *IDEXReg, 1) // The channel to manage async communication between decode and execute units
	exmemChan chan *EXMEMReg = make(chan *EXMEMReg, 1) // The channel to manage async communication between execute and memory data units
	memwbChan chan *MEMWBReg = make(chan *MEMWBReg, 1) // The channel to manage async communication between memory data and writeback units
	endInstrChan chan bool = make(chan bool, 1) // The channel to manage async communication for the end of the instruction
	fetchRunning bool = false // Determines if the fetch unit is currently running
	decodeRunning bool = false // Determines if the decode unit is currently running
	executeRunning bool = false // Determines if the execute unit is currently running
	memRunning bool = false // Determines if the mem data unit is currently running
	writebackRunning bool = false // Determines if the writeback unit is currently running
)

// Creates the CPU
func NewCpu(mem *memory.Memory) *Cpu {
	cpu := Cpu {
		hw: hardware.NewHardware("CPU", 0),
		acc: 0x0,
		reg: make([]uint64, 32),
		programCounter: 0,
		fetchUnit: NewFetchUnit(mem),
		executeUnit: NewExecuteUnit(),
		memDataUnit: NewMemDataUnit(mem),
		regLocks: util.NewQueue(),
	}
	cpu.decodeUnit = NewDecodeUnit(&cpu)
	cpu.writebackUnit = NewWritebackUnit(&cpu)
	return &cpu
}

// Function that gets called every clock cycle
func (cpu *Cpu) Pulse() {
	// Clear the writeback unit
	if len(endInstrChan) == 1 {
		<- endInstrChan
		writebackRunning = false
	}

	// Clear the mem data unit and writeback next instruction if available
	if len(endInstrChan) == 0 && len(memwbChan) == 1 && !writebackRunning {
		cpu.memwbReg = <- memwbChan
		memRunning = false
		cpu.Log(fmt.Sprintf("Starting writeback: %d", cpu.memwbReg.incrementedPC - 4))
		go cpu.writebackUnit.HandleWriteback(endInstrChan, cpu.memwbReg)
		writebackRunning = true
	}

	// Clear the execute unit and handle memory if available
	if len(memwbChan) == 0 && len(exmemChan) == 1 && !memRunning {
		cpu.exmemReg = <- exmemChan
		executeRunning = false
		cpu.Log(fmt.Sprintf("Starting mem data access: %d", cpu.exmemReg.incrementedPC - 4))
		go cpu.memDataUnit.HandleMemoryAccess(memwbChan, cpu.exmemReg)
		memRunning = true
	}

	// Clear decode unit and execute if available
	if len(exmemChan) == 0 && len(idexChan) == 1 && !executeRunning {
		cpu.idexReg = <- idexChan
		decodeRunning = false
		cpu.Log(fmt.Sprintf("Starting execute: %d", cpu.idexReg.incrementedPC - 4))
		go cpu.executeUnit.ExecuteInstruction(exmemChan, cpu.idexReg)
		executeRunning = true
	}

	// Clear the fetch unit and decode if available
	if len(idexChan) == 0 && len(ifidChan) == 1 && !decodeRunning {
		cpu.ifidReg = <- ifidChan
		fetchRunning = false
		cpu.Log(fmt.Sprintf("Starting decode: %d", cpu.ifidReg.incrementedPC - 4))
		go cpu.decodeUnit.DecodeInstruction(idexChan, cpu.ifidReg)
		decodeRunning = true
	}

	// Fetch if available
	if len(ifidChan) == 0 && !fetchRunning {
		cpu.Log(fmt.Sprintf("Starting fetch: %d", cpu.programCounter))
		go cpu.fetchUnit.FetchInstruction(ifidChan, &cpu.programCounter)
		fetchRunning = true
	}
}

// Logs a message
func (cpu *Cpu) Log(msg string) {
	cpu.hw.Log(msg)
}

// Gets the program counter
func (cpu *Cpu) GetProgramCounter() uint64 {
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