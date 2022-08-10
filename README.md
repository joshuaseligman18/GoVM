# GoVM

## Contents
1. [Overview](#overview)
2. [Setup Instructions](#setup-instructions)
3. [Instruction Set](#instruction-set)
## Overview
GoVM is a virtual ARM processor written in Go. The project was inspired by Professor Gormanly's 6502 project from his Computer Organization and Architecture class at Marist College. The instruction set is based off of that described in *Computer Organization and Design ARM Edition: The Hardware Software Interface* by David A. Patterson and John L. Hennessy. By using the ARM architecture, GoVM is able to take advantage of more modern CPU design practices such as scalar processing, which improves performance and maximizes the utilization of the hardware of the CPU.
## Setup Instructions
1. Install Go from the official website [here](https://go.dev/)
2. Create a Go module by running `go mod init <project-name>` in the directory of your project
3. Add GoVM to your project dependencies by running `go get github.com/joshuaseligman/GoVM`
4. Import GoVM at the top of your code: `import "github.com/joshuaseligman/GoVM"`
5. Clean up the dependencies and run `go mod tidy`

## Instruction Set

### Instruction Types
* [Arithmetic Instructions](#arithmetic-instructions)
* [Data Transfer Instructions](#data-transfer-instructions)
* [Branching Instructions](#branching-instructions)
* [Miscellaneous Instructions](#miscellaneous-instructions)

### Arithmetic Instructions

**ADD** - Adds the contents of 2 registers and saves the output in another register. The ALU flags are ***NOT*** set from this instruction.
```
ADD Rd, Rm, Rn
```
*Rd: The destination register (X0 - X30)* <br />
*Rm: The first register for the operation (X0 - X30, XZR)* <br />
*Rn: The second register for the operation (X0 - X30, XZR)*

**ADDS** - Adds the contents of 2 registers and saves the output in another register. The ALU flags ***ARE*** set from this instruction.
```
ADDS Rd, Rm, Rn
```
*Rd: The destination register (X0 - X30)* <br />
*Rm: The first register for the operation (X0 - X30, XZR)* <br />
*Rn: The second register for the operation (X0 - X30, XZR)*

**ADDI** - Adds a constant to the contents of a register and saves the output in another register. The ALU flags are ***NOT*** set from this instruction.
```
ADDI Rd, Rm, Imm
```
*Rd: The destination register (X0 - X30)* <br />
*Rm: The register for the operation (X0 - X30, XZR)* <br />
*Imm: The 12-bit unsigned immediate value to add (0x000 - 0xFFF)*

**ADDIS** - Adds a constant to the contents of a register and saves the output in another register. The ALU flags ***ARE*** set from this instruction.
```
ADDIS Rd, Rm, Imm
```
*Rd: The destination register (X0 - X30)* <br />
*Rm: The register for the operation (X0 - X30, XZR)* <br />
*Imm: The 12-bit unsigned immediate value to add (0x000 - 0xFFF)*

**SUB** - Subtracts the contents of 2 registers and saves the output in another register. The ALU flags are ***NOT*** set from this instruction.
```
SUB Rd, Rm, Rn
```
*Rd: The destination register (X0 - X30)* <br />
*Rm: The first register for the operation (X0 - X30, XZR)* <br />
*Rn: The second register for the operation (X0 - X30, XZR)*

**SUBS** - Subtracts the contents of 2 registers and saves the output in another register. The ALU flags ***ARE*** set from this instruction.
```
SUBS Rd, Rm, Rn
```
*Rd: The destination register (X0 - X30)* <br />
*Rm: The first register for the operation (X0 - X30, XZR)* <br />
*Rn: The second register for the operation (X0 - X30, XZR)*

**SUBI** - Subtracts a constant from the contents of a register and saves the output in another register. The ALU flags are ***NOT*** set from this instruction.
```
SUBI Rd, Rm, Imm
```
*Rd: The destination register (X0 - X30)* <br />
*Rm: The register for the operation (X0 - X30, XZR)* <br />
*Imm: The 12-bit unsigned immediate value to add (0x000 - 0xFFF)*

**SUBIS** - Subtracts a constant from the contents of a register and saves the output in another register. The ALU flags ***ARE*** set from this instruction.
```
SUBIS Rd, Rm, Imm
```
*Rd: The destination register (X0 - X30)* <br />
*Rm: The register for the operation (X0 - X30, XZR)* <br />
*Imm: The 12-bit unsigned immediate value to add (0x000 - 0xFFF)*

### Data Transfer Instructions
**MOVZ** - Loads a constant into a register ***WITHOUT*** retaining the previous contents of the register.
```
MOVZ Rd, Imm, LSL Amt
```
*Rd: The destination register (X0 - X30)* <br />
*Imm: The 16-bit constant to load to the register (0x0000 - 0xFFFF)* <br />
*Amt: The amount to left-shift the immediate by (0, 16, 32, 48)*

**MOVK** - Loads a constant into a register ***AND*** retains the unaffected contents of the register.
```
MOVK Rd, Imm, LSL Amt
```
*Rd: The destination register (X0 - X30)* <br />
*Imm: The 16-bit constant to load to the register (0x0000 - 0xFFFF)* <br />
*Amt: The amount to left-shift the immediate by (0, 16, 32, 48)*

**LDUR** - Loads a doubleword from memory into a register.
```
LDUR Rd, Rm, Addr
```
*Rd: The destination register (X0 - X30)* <br />
*Rm: The register to use in determining the address (X0 - X30, XZR)* <br />
*Addr: The 9-bit signed 2's complement value to add to the value in Rm to determine the memory location (0x000 - 0x1FF)*

**LDURB** - Loads a byte from memory into a register.
```
LDURB Rd, Rm, Addr
```
*Rd: The destination register (X0 - X30)* <br />
*Rm: The register to use in determining the address (X0 - X30, XZR)* <br />
*Addr: The 9-bit signed 2's complement value to add to the value in Rm to determine the memory location (0x000 - 0x1FF)*

**LDURH** - Loads a halfword from memory into a register.
```
LDURH Rd, Rm, Addr
```
*Rd: The destination register (X0 - X30)* <br />
*Rm: The register to use in determining the address (X0 - X30, XZR)* <br />
*Addr: The 9-bit signed 2's complement value to add to the value in Rm to determine the memory location (0x000 - 0x1FF)*

**LDURSW** - Loads a signed word from memory into a register.
```
LDURSW Rd, Rm, Addr
```
*Rd: The destination register (X0 - X30)* <br />
*Rm: The register to use in determining the address (X0 - X30, XZR)* <br />
*Addr: The 9-bit signed 2's complement value to add to the value in Rm to determine the memory location (0x000 - 0x1FF)*

**STUR** - Stores the contents of a register into memory.
```
STUR Rd, Rm, Addr
```
*Rd: The register whose contents should be stored (X0 - X30)* <br />
*Rm: The register to use in determining the address (X0 - X30, XZR)* <br />
*Addr: The 9-bit signed 2's complement value to add to the value in Rm to determine the memory location (0x000 - 0x1FF)*

**STURB** - Stores a byte from a register into memory.
```
STURB Rd, Rm, Addr
```
*Rd: The destination register (X0 - X30)* <br />
*Rm: The register to use in determining the address (X0 - X30, XZR)* <br />
*Addr: The 9-bit signed 2's complement value to add to the value in Rm to determine the memory location (0x000 - 0x1FF)*

**STURH** - Stores a halfword from a register into memory.
```
STURH Rd, Rm, Addr
```
*Rd: The destination register (X0 - X30)* <br />
*Rm: The register to use in determining the address (X0 - X30, XZR)* <br />
*Addr: The 9-bit signed 2's complement value to add to the value in Rm to determine the memory location (0x000 - 0x1FF)*

**STURW** - Stores a word from a register into memory.
```
STURW Rd, Rm, Addr
```
*Rd: The destination register (X0 - X30)* <br />
*Rm: The register to use in determining the address (X0 - X30, XZR)* <br />
*Addr: The 9-bit signed 2's complement value to add to the value in Rm to determine the memory location (0x000 - 0x1FF)*

### Branching Instructions

**B** - Branches to a new location in the program.
```
B Addr
```
*Addr: The 26-bit signed 2's complement relative address to branch to (0x0000000 - 0x3FFFFFF)*

**CBZ** - Branches to a new location in the program if the given register ***IS*** equal to 0.
```
CBZ Rm, Addr
```
*Rm: The register whose value should be tested* <br />
*Addr: The 19-bit signed 2's complement relative address to branch to (0x00000 - 0x7FFFF)*

**CBNZ** - Branches to a new location in the program if the given register is ***NOT*** equal to 0.
```
CBNZ Rm, Addr
```
*Rm: The register whose value should be tested* <br />
*Addr: The 19-bit signed 2's complement relative address to branch to (0x00000 - 0x7FFFF)*

### Miscellaneous Instructions

**DATA** - Places a value in memory.
```
DATA Val
```
*Val: The 32-bit value to place in memory at the location within the program*

**HLT** - Stops the program.
```
HLT
```
