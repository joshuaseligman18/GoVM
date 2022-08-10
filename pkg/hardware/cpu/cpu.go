package cpu

import (
	"fmt"

	"github.com/joshuaseligman/GoVM/pkg/hardware"
	"github.com/joshuaseligman/GoVM/pkg/hardware/clock"
	"github.com/joshuaseligman/GoVM/pkg/hardware/memory"
	"github.com/joshuaseligman/GoVM/pkg/util"
)

// Struct for the CPU
type Cpu struct {
	hw *hardware.Hardware // The hardware struct
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

// Struct for the CPU to be used with the API
type CpuAPI struct {
	Reg [32]uint64 `json:"registers"`// Other registers
	ProgramCounter uint64 `json:"programCounter"` // The address of the current instruction being fetched
	IfidReg *IFIDReg `json:"ifidReg"`// The register between the fetch and decode units
	IdexReg *IDEXReg `json:"idexReg"` // The register between the decode and execute units
	ExmemReg *EXMEMReg `json:"exmemReg"` // The register between the execute and memory data units
	MemwbReg *MEMWBReg `json:"memwbReg"` // The register between the memory data and writeback units
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
func NewCpu(mem *memory.Memory, clk *clock.Clock) *Cpu {
	cpu := Cpu {
		hw: hardware.NewHardware("CPU", 0),
		reg: make([]uint64, 32),
		programCounter: 0,
		fetchUnit: NewFetchUnit(mem),
		memDataUnit: NewMemDataUnit(mem),
		regLocks: util.NewQueue(),
	}
	cpu.decodeUnit = NewDecodeUnit(&cpu)
	cpu.executeUnit = NewExecuteUnit(&cpu, clk)
	cpu.writebackUnit = NewWritebackUnit(&cpu)
	return &cpu
}

// Function that gets called every clock cycle
func (cpu *Cpu) Pulse() any {
	cpu.Log(fmt.Sprintf("%t %t %t %t %t", fetchRunning, decodeRunning, executeRunning, memRunning, writebackRunning))
	cpu.Log(fmt.Sprintf("%d %d %d %d %d", len(ifidChan), len(idexChan), len(exmemChan), len(memwbChan), len(endInstrChan)))

	// Clear the writeback unit
	if len(endInstrChan) == 1 {
		<- endInstrChan
		writebackRunning = false
	}

	// Clear the mem data unit and writeback next instruction if available
	if len(endInstrChan) == 0 && len(memwbChan) == 1 && !writebackRunning {
		cpu.memwbReg = <- memwbChan
		cpu.Log(fmt.Sprintf("Starting writeback: %d", cpu.memwbReg.IncrementedPC - 4))
		go cpu.writebackUnit.HandleWriteback(endInstrChan, cpu.memwbReg)
		writebackRunning = true
		memRunning = false
	}

	// Clear the execute unit and handle memory if available
	if len(memwbChan) == 0 && len(exmemChan) == 1 && !memRunning {
		cpu.exmemReg = <- exmemChan
		executeRunning = false
		cpu.Log(fmt.Sprintf("Starting mem data access: %d", cpu.exmemReg.IncrementedPC - 4))
		go cpu.memDataUnit.HandleMemoryAccess(memwbChan, cpu.exmemReg)
		memRunning = true
	}

	// Clear decode unit and execute if available
	if len(exmemChan) == 0 && len(idexChan) == 1 && !executeRunning && !cpu.executeUnit.flushing {
		cpu.idexReg = <- idexChan
		decodeRunning = false
		cpu.Log(fmt.Sprintf("Starting execute: %d", cpu.idexReg.IncrementedPC - 4))
		go cpu.executeUnit.ExecuteInstruction(exmemChan, cpu.idexReg, &memRunning, &writebackRunning)
		executeRunning = true
	}

	// Clear the fetch unit and decode if available
	if len(idexChan) == 0 && len(ifidChan) == 1 && !decodeRunning {
		cpu.ifidReg = <- ifidChan
		fetchRunning = false
		cpu.Log(fmt.Sprintf("Starting decode: %d", cpu.ifidReg.IncrementedPC - 4))
		go cpu.decodeUnit.DecodeInstruction(idexChan, cpu.ifidReg)
		decodeRunning = true
	}

	// Fetch if available
	if len(ifidChan) == 0 && !fetchRunning {
		cpu.Log(fmt.Sprintf("Starting fetch: %d", cpu.programCounter))
		go cpu.fetchUnit.FetchInstruction(ifidChan, &cpu.programCounter)
		fetchRunning = true
	}

	return cpu.ConvertAPI()
}

// Function to stop the pipeline
func (cpu *Cpu) FlushPipeline(newPC uint64) {
    // Wait until both fetch and decode units are sitting
    for len(ifidChan) == 0 || len(idexChan) == 0 {
        continue
    }

    // Remove the data in the channels
    <- ifidChan
    lastDecode := <- idexChan

	if lastDecode.addedLock {
		cpu.regLocks.RemoveLast()
	}

	// Set the new program counter and reset the fetch and decode units and intermediate registers
	cpu.programCounter = newPC
	cpu.ifidReg = nil
	cpu.idexReg = nil
	fetchRunning = false
	decodeRunning = false

	// Tell the execute unit it is safe to execute its next instruction
	cpu.executeUnit.flushing = false
}

// Logs a message
func (cpu *Cpu) Log(msg string) {
	cpu.hw.Log(msg)
}

// Gets the program counter
func (cpu *Cpu) GetProgramCounter() uint64 {
	return cpu.programCounter
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

// Resets the CPU for future use
func (cpu *Cpu) ResetCpu() {
	// All registers go back to 0
	for i := 0; i < len(cpu.reg); i++ {
		cpu.reg[i] = 0x0
	}
	cpu.programCounter = 0x0

	// Reset the units that have data in them
	cpu.fetchUnit.Reset()
	cpu.executeUnit.Reset()
	cpu.memDataUnit.Reset()

	// More registers get cleared
	cpu.ifidReg = nil
	cpu.idexReg = nil
	cpu.exmemReg = nil
	cpu.memwbReg = nil
	cpu.regLocks.ResetQueue()

	// Clear all channels
	for len(ifidChan) > 0 {
		<- ifidChan
	}

	for len(idexChan) > 0 {
		<- idexChan
	}

	for len(exmemChan) > 0 {
		<- exmemChan
	}

	for len(memwbChan) > 0 {
		<- memwbChan
	}

	for len(endInstrChan) > 0 {
		<- endInstrChan
	}
	
	// Lastly allow the units to run again by setting the flags to false
	fetchRunning = false
	decodeRunning = false
	executeRunning = false
	memRunning = false
	writebackRunning = false
}

// Exports the CPU status into an API-friendly format
func (cpu *Cpu) ConvertAPI() *CpuAPI {
	// Convert the slice to an array for the API to output the actual array
	newReg := [32]uint64{}
	copy(newReg[:], cpu.reg)

	return & CpuAPI {
		Reg: newReg,
		ProgramCounter: cpu.programCounter,
		IfidReg: cpu.ifidReg,
		IdexReg: cpu.idexReg,
		ExmemReg: cpu.exmemReg,
		MemwbReg: cpu.memwbReg,
	}
}