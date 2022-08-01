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
3. Add GoVM to your project dependencies by running `go get github.com/joshuaseligman/GoVM@v0.6.0`
4. Import GoVM at the top of your code: `import "github.com/joshuaseligman/GoVM"`

## Instruction Set

### Instruction Types
* [Arithmetic Instructions](#arithmetic-instructions)
* [Data Transfer Instructions](#data-transfer-instructions)
* [Branching Instructions](#branching-instructions)
* [Miscellaneous Instructions](#miscellaneous-instructions)

### Arithmetic Instructions

**ADD** - Adds the contents of 2 registers and saves the output into another register. The ALU flags are ***NOT*** set from this instruction.
```
ADD Rd, Rm, Rn
```
*Rd: The destination register (X0 - X30)* <br />
*Rm: The first register for the operation (X0 - X30, XZR)* <br />
*Rn: The second register for the operation (X0 - X30, XZR)*

**ADDS** - Adds the contents of 2 registers and saves the output into another register. The ALU flags ***ARE*** set from this instruction.
```
ADDS Rd, Rm, Rn
```
*Rd: The destination register (X0 - X30)* <br />
*Rm: The first register for the operation (X0 - X30, XZR)* <br />
*Rn: The second register for the operation (X0 - X30, XZR)*

### Data Transfer Instructions
### Branching Instructions
### Miscellaneous Instructions